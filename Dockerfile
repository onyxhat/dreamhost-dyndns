FROM alpine:latest

COPY ./bin/dreamhost-dyndns-linux-amd64 /app/dreamhost-dyndns
RUN chmod 770 /app/dreamhost-dyndns

CMD [ "/app/dreamhost-dyndns" ]