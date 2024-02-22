FROM ollama/ollama

RUN apt-get update && apt-get install -y curl

COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
