from fastapi import FastAPI, Request, Response
import pika
import redis
import json
import os


rmq_queue = os.getenv('RABBITMQ_QUEUE', 'my_rmq_queue')
redis_queue = os.getenv('REDIS_QUEUE', 'my_redis_queue')
r = redis.Redis(host=os.getenv('REDIS_HOST', 'redis'),
                port=6379,
                decode_responses=True)
app = FastAPI()

def rmq_get_connection():
    return pika.BlockingConnection(pika.ConnectionParameters(
        host=os.getenv('RABBITMQ_HOST', 'rabbitmq'),
        credentials=pika.PlainCredentials(
            os.getenv('RABBITMQ_USER', 'admin'),
            os.getenv('RABBITMQ_PASSWORD', 'admin')
        )
    ))

@app.get("/")
def read_root():
    return {"message": "Hello, world"}

@app.get("/healthy")
def healthy(response: Response):
    try:
        r.ping()
        redis_status = "ok"
    except:
        redis_status = "nok"

    try:
        connection = rmq_get_connection()
        connection.close()
        rabbit_status = "ok"
    except:
        rabbit_status = "nok"

    if redis_status != "ok" or rabbit_status != "ok":
        response.status_code = 503
    return {
        "redis": redis_status,
        "rabbitmq": rabbit_status
    }

@app.post("/rmq-send")
async def rmq_send_message(request: Request):
    body = await request.body()
    connection = rmq_get_connection()
    channel = connection.channel()
    channel.queue_declare(rmq_queue)
    channel.basic_publish(exchange='', routing_key=rmq_queue, body=body)
    connection.close()
    return {"yoo, check what I've just sent": body.decode()}

@app.get("/rmq-get")
def rmq_receive_message():
    connection = rmq_get_connection()
    channel = connection.channel()
    channel.queue_declare(rmq_queue)
    method, properties, body = channel.basic_get(queue=rmq_queue, auto_ack=True)
    connection.close()
    return {"damn, dats what I've just received": body.decode() if body else 'hah actually no message found in the queue'}

@app.post("/redis-send")
async def redis_send(request: Request):
    message = await request.body()
    r.rpush(redis_queue, message.decode('utf-8'))
    return {"status": "sent AF boi", "message": message.decode('utf-8')}

@app.get("/redis-get")
async def redis_get():
    message = r.lpop(redis_queue)
    return {"look what I've got": message if message else 'oopsie no messages present'}

@app.get("/calculate")
def heavy_calc():
    cached = r.get("calc_result")
    if cached:
        return {"result": json.loads(cached), "from_cache": True}

    result = sum(i ** 2 for i in range(1000000))
    r.setex("calc_result", os.getenv('REDIS_CACHE_TIME', 60), json.dumps(result))
    return {"result": result, "from_cache": False}