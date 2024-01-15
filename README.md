
<div align="center">
<!--   <a href="https://github.com/JoyalAJohney/Distributed-Chat-Backend/">
    <img src="https://raw.githubusercontent.com/othneildrew/Best-README-Template/master/images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

  <h1 align="center">Realtime Distributed Messaging Platform 🚀</h1>
  <img src="https://raw.githubusercontent.com/JoyalAJohney/Realtime-Distributed-Chat/main/assets/babylon.png" alt="landing page">

  <div align="center">
    <br/>
    High throuput, low latency messaging platform build using Go-fiber and Websockets. Configured Nginx as a reverse proxy for efficient layer 7 loadbalancing 🌌. Redis pub/sub for real-time communication ⚡, Kafka for low-latency processing and Postgres for data storage. Setup infrastructure on AWS using Terraform 🌀 and build CI/CD pipelines using github actions.
  </div>

  <br />

  <img src="https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB" alt="React Badge">
  *
  <img src="https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white" alt="Nginx Badge">
  *
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Golang Badge">
  *
  <img src="https://img.shields.io/badge/redis-%23DD0031.svg?&style=for-the-badge&logo=redis&logoColor=white" alt="Redis Badge">
  *
  <img src="https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white" alt="Postgres Badge">
  *
  <img src="https://img.shields.io/badge/Apache%20Kafka-000?style=for-the-badge&logo=apachekafka" alt="Kafka Badge">
  *
  <img src="https://img.shields.io/badge/terraform-%235835CC.svg?style=for-the-badge&logo=terraform&logoColor=white" alt="Terraform Badge">
  *
  <img src="https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white" alt="Docker Badge">
  *
  <img src="https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white" alt="AWS Badge">

</div>


  

## Setting Up

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

## Running the app

Execute the below command to build the application containers
```bash
$ docker-compose up --build
```
If the application starts perfectly fine, you should be able to head over to http://NGINX_HOST:NGINX_PORT/
