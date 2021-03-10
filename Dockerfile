FROM gcr.io/distroless/base
COPY ./bin/dreamhost-dyndns-linux-amd64 /app
ENTRYPOINT [ "/app" ]