FROM alpine:latest

WORKDIR /app

COPY ./bin/dreamhost-dyndns-linux-amd64 /app/dreamhost-dyndns

ENTRYPOINT [ "/app/dreamhost-dyndns" ]