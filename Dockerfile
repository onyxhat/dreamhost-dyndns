FROM alpine:latest

WORKDIR /app

COPY ./bin/dreamhost-dyndns-linux-amd64 /dreamhost-dyndns

ENTRYPOINT [ "/app/dreamhost-dyndns" ]