FROM alpine:latest
ARG DOCKER_BIN

LABEL MAINTAINER="onyxhat"
LABEL REPO="https://github.com/onyxhat/dreamhost-dyndns"

COPY ./bin/${DOCKER_BIN} /app/dreamhost-dyndns

RUN chmod -R +x /app

CMD [ "/app/dreamhost-dyndns" ]