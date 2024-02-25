from diagrams import Diagram, Cluster
from diagrams.onprem.client import User
from diagrams.onprem.compute import Server
from diagrams.onprem.network import Nginx
from diagrams.onprem.queue import Kafka
from diagrams.onprem.database import PostgreSQL
from diagrams.onprem.inmemory import Redis
from diagrams.programming.framework import React



with Diagram(name="Distributed Chat Architecture", show=False):
    frontend = React("Frontend")
    
    with Cluster("Backend Servers"):
        nginx = Nginx("Nginx Server")
        go_servers = [Server(f"Go Server {i}") for i in range(1, 4)]
        nginx >> go_servers

    redis_pubsub = Redis("Redis Pub/Sub")
    kafka = Kafka("Kafka Queue")
    go_servers >> redis_pubsub >> kafka

    batch_server = Server("Batch Go Server")
    db = PostgreSQL("PostgreSQL")
    
    kafka >> batch_server >> db

    frontend >> nginx
