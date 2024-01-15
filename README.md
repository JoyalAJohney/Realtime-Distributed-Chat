
<div align="center">
  <a href="https://github.com/JoyalAJohney/Distributed-Chat-Backend/">
    <img src="https://raw.githubusercontent.com/othneildrew/Best-README-Template/master/images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Realtime Distributed Chat</h3>

  <p align="center">
    Real-time distributed chat that processes message with very high throughput and low latency
    <br />
    <a href="https://github.com/JoyalAJohney/Distributed-Chat-Backend/"><strong>Explore the docs »</strong></a>
    <br />
  </p>
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Golang Badge">
  *
  <img src="https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB" alt="React Badge">
  *
  <img src="https://img.shields.io/badge/redis-%23DD0031.svg?&style=for-the-badge&logo=redis&logoColor=white" alt="Redis Badge">
  *
  <img src="https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white" alt="Postgres Badge">
</div>
  

## Setting Up

### Env changes
* Create a .env file from the env.sample file.
* Fill in the values based on your required configuration.
* Make sure that the .env file is in the same level as docker-compose.yml file
  
```bash
# Redis Config
REDIS_PORT=6379
REDIS_HOST=redis

# Database Config
POSTGRES_DATABASE=chat_db
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_HOST=postgres
POSTGRES_PORT=5432

# Kafka config
KAFKA_HOST=kafka
KAFKA_PORT=9092
KAFKA_TOPIC=chat_messages
KAFKA_GROUP_ID=chat_group
ZOOKEEPER_PORT=2181

# Authentication
JWT_SECRET=secret

# Reverse proxy config
NGINX_ENV=local
NGINX_PORT=8080
NGINX_HOST=localhost

# Backend Servers
SERVER_PORT=8080
```

### TLS/SSL setup

If you do not have a TLS/SSL certificate setup, disable or comment these code changes

docker-compose.yml
```bash
"${NGINX_SSL_PORT}:443"

/etc/letsencrypt:/etc/letsencrypt:ro
```

nginx.conf

```bash
# Remove this code
server {
    listen 80;

    # Redirect all HTTP requests to HTTPS
    return 301 https://$host$request_uri;
}

# Remove ssl_certificate and ssl_certificate_key
```


## Running the app

Execute the below command to build the application containers
```bash
$ docker-compose up --build
```
If the application starts perfectly fine, you should be able to head over to http://NGINX_HOST:NGINX_PORT/


## Testing
* Make sure you are running on node version 16 or above
* Change ```POSTGRES_HOST=localhost``` in .env file when running tests

```bash
# unit tests
$ npm run test

# e2e tests
$ npm run test:e2e

# test coverage
$ npm run test:cov
```