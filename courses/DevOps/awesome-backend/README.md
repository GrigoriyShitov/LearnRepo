# 4 Task

## Done

- **Redis** has been added to the same Docker network as **RabbitMQ** (`awesome-backend-network`).
- **Redis** data is now fully persistent thanks to a named volume `redis_data`.
- Cache lifetime in Redis is fully parameterized through the environment variable `REDIS_CACHE_TIME` (in seconds). Can be easily changed in `.env`.
- Added a proper `healthcheck` for the Redis service using the official `redis-cli ping` command.
- Updated `depends_on` for the main application(s) to match the RabbitMQ pattern — the app now waits for `service_healthy` state of both RabbitMQ and Redis.
- **Advanced**: Implemented a separate watchdog container (`app-checker`) that checks the `/healthcheck` endpoint of the application(s) every 60 seconds. The container **fails** (exits with code 1) if it receives HTTP status **503** from any monitored app version.
- **Advanced**: Implemented communication redundancy (failover) between RabbitMQ and Redis:
    - If RabbitMQ becomes unavailable, the application automatically switches message/queue operations to Redis.
    - If Redis is down, the application falls back to RabbitMQ.
    - Both services act as mutual backup for critical communication paths.

## Technical Highlights

### Successful healthcheck 
![Successful healthcheck](https://github.com/GrigoriyShitov/LearnRepo/blob/master/images/Healthcheck%20200.png)

### Healthcheck without rabbitMQ
![Successful healthcheck](https://github.com/GrigoriyShitov/LearnRepo/blob/master/images/Healthcheck%20with%20unavailabe%20RabbitMQ.png)

### Healthcheck unavailable service
![Unsuccessful healthcheck](https://github.com/GrigoriyShitov/LearnRepo/blob/master/images/Healthcheck%20without%20cache%20and%20RabbitMQ.png)