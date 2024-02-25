from diagrams import Diagram, Cluster
from diagrams.onprem.client import User
from diagrams.onprem.compute import Server
from diagrams.onprem.network import Nginx
from diagrams.onprem.queue import Kafka
from diagrams.onprem.database import PostgreSQL
from diagrams.onprem.inmemory import Redis
from diagrams.programming.framework import React
from diagrams.aws.ml import Sagemaker as LLMModel

with Diagram(name="", show=False):
    frontend = React("Frontend")

    with Cluster("Backend Servers"):
        nginx = Nginx("Nginx Server")
        go_servers = [Server(f"Go Server {i}") for i in range(1, 4)]
        nginx >> go_servers


    kafka = Kafka("Kafka Queue")
    llmodel = LLMModel("LLM Model")
    redis_pubsub = Redis("Redis Pub/Sub")
    
    go_servers >> redis_pubsub

    redis_pubsub >> kafka
    redis_pubsub >> llmodel
    llmodel >> redis_pubsub


    batch_server = Server("Batch Go Server")
    db = PostgreSQL("PostgreSQL")
    
    kafka >> batch_server >> db
    frontend >> nginx
