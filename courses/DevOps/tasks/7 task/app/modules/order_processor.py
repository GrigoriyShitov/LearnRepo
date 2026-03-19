from fastapi import FastAPI, HTTPException
import aio_pika
import asyncpg
import asyncio
import os
import json
import logging
import socket
from pythonjsonlogger import jsonlogger


logger = logging.getLogger()
logger.setLevel(logging.INFO)
log_formatter = jsonlogger.JsonFormatter('%(timestamp)s %(levelname)s %(module)s %(message)s %(hostname)s', timestamp=True)

console_handler = logging.StreamHandler()
console_handler.setFormatter(log_formatter)
logger.addHandler(console_handler)

json_handler = logging.FileHandler('log/modules/processor.log')
json_handler.setFormatter(log_formatter)
logger.addHandler(json_handler)

hostname = socket.gethostname()

app = FastAPI()

async def get_db_pool():
    return await asyncpg.create_pool(
        host=os.getenv("ORDERS_DB_HOST"),
        port=os.getenv("ORDERS_DB_PORT"),
        user=os.getenv("ORDERS_DB_USER"),
        password=os.getenv("ORDERS_DB_PASSWORD"),
        database=os.getenv("ORDERS_DB_NAME")
    )

@app.get("/order/{order_id}")
async def get_order(order_id: int):
    logger.info(f"Requesting order {order_id} status...", extra={'hostname': hostname})
    pool = await get_db_pool()

    async with pool.acquire() as conn:
        row = await conn.fetchrow(
            "SELECT id, status, description FROM orders WHERE id = $1",
            order_id
        )

    await pool.close()

    if not row:
        logger.error(f"Order {order_id} not found in orders database", extra={'hostname': hostname})
        raise HTTPException(status_code=404, detail="Order not found")

    logger.info(f"Order {order_id} status requested - current status: {row['status']}", extra={'hostname': hostname})

    return {
        "id": row['id'],
        "status": row['status'],
        "description": row['description']
    }

async def process_message(message: aio_pika.IncomingMessage):
    async with message.process():
        logger.info("Message received, processing...", extra={'hostname': hostname})
        data = json.loads(message.body.decode('utf-8'))
        order_id = data['id']
        description = data['description']

        pool = await get_db_pool()

        async with pool.acquire() as conn:
            await conn.execute(
                "INSERT INTO orders (id, status, description) VALUES ($1, 'created', $2)",
                order_id, description
            )

        await pool.close()
        logger.info(f"Order {order_id} created: {description}", extra={'hostname': hostname})

@app.on_event("startup")
async def startup():
    asyncio.create_task(consume_queue())

async def consume_queue():
    connection = await aio_pika.connect_robust(
        host=os.getenv("RMQ_HOST", "rabbitmq"),
        login=os.getenv("RMQ_USER", "admin"),
        password=os.getenv("RMQ_PASSWORD", "admin")
    )

    channel = await connection.channel()
    queue = await channel.declare_queue(
        os.getenv("RMQ_QUEUE", "orders_queue"), durable=True
    )

    await queue.consume(process_message)
    logger.info("Processor started, waiting for messages...", extra={'hostname': hostname})
