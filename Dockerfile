# Build the app
FROM golang:alpine AS build

WORKDIR /go/src/github.com/onyxhat/dreamhost-dyndns

RUN apk update && apk add --no-cache git

COPY . .

RUN go get . && \
    go build -ldflags="-s -w" -o "./bin/dreamhost-dyndns"

#Compress the app
#FROM gruebel/upx:latest as upx

#COPY --from=build /go/src/github.com/onyxhat/dreamhost-dyndns/bin/* /dreamhost-dyndns.org
#RUN upx --best --lzma -o /dreamhost-dyndns /dreamhost-dyndns.org

# Store the app
FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/github.com/onyxhat/dreamhost-dyndns/bin/dreamhost-dyndns ./

ENTRYPOINT [ "/app/dreamhost-dyndns" ]