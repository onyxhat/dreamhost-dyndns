#Compress the app
FROM gruebel/upx:latest as upx

COPY ./bin/dreamhost-dyndns-linux-amd64 /dreamhost-dyndns.org
RUN chmod +x /dreamhost-dyndns.org && \
    upx --best --lzma -o /dreamhost-dyndns /dreamhost-dyndns.org

# Store the app
FROM alpine:latest

WORKDIR /app

COPY --from=upx /dreamhost-dyndns ./

ENTRYPOINT [ "/app/dreamhost-dyndns" ]