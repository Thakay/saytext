FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y flite

ADD app /app
RUN chmod +x /app
ENTRYPOINT ["/app"]