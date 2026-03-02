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

# 5 task

## Docker Compose Setup & High Availability Implementation

All requirements from the task have been successfully implemented:

- Configured Docker networks so that both application replicas can only access the HAProxy service, but have no direct visibility to PostgreSQL or ZooKeeper services. This was achieved by isolating networks.
- Based on the `postgres0.yml` config, created `postgres1.yml` for the second node. Launched the cluster and verified its operation by checking replication status, leader election, and data consistency across nodes.
- Transformed the two application instances into proper replicas by ensuring consistency in environment variables, configurations, and codebases. Implemented traffic balancing through HAProxy similar to the PostgreSQL example in `haproxy.cfg` (using round-robin load balancing). Restricted external access to replica ports; endpoints are only accessible via HAProxy.
- **Advanced**: Devised and implemented a mechanism to ensure the `pg-master` container is always the actual master of the cluster. 

### Technical Highlights

#### pg-slave launched without pg-master
![pg-slave](https://github.com/GrigoriyShitov/LearnRepo/blob/master/images/pg-slave-error.png)

#### HAproxy report
![HAproxy](https://github.com/GrigoriyShitov/LearnRepo/blob/master/images/HAproxy.png)
