FROM alpine:latest
LABEL MAINTAINER="onyxhat"
LABEL REPO="https://github.com/onyxhat/dreamhost-dyndns"
COPY ./bin/dreamhost-dyndns-linux-386 /app/dreamhost-dyndns
CMD [ "/app/dreamhost-dyndns" ]