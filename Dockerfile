# Build the app
FROM golang:alpine AS build

WORKDIR /go/src/github.com/onyxhat/dreamhost-dyndns

COPY . .

RUN go build -ldflags="-s -w" -o "./bin/dreamhost-dyndns"

#Compress the app
FROM gruebel/upx:latest as upx

COPY --from=build /go/src/github.com/onyxhat/dreamhost-dyndns/bin/* /dreamhost-dyndns.org
RUN upx --best --lzma -o /dreamhost-dyndns /dreamhost-dyndns.org

# Store the app
FROM scratch

WORKDIR /app

COPY --from=upx /dreamhost-dyndns ./

ENTRYPOINT [ "dreamhost-dyndns" ]