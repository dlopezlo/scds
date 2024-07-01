FROM golang:1.22.3-alpine3.18 AS builder
RUN apk add --no-cache git upx

WORKDIR /app

COPY ["src/", "./"]
RUN ls -l utils/ && go mod download -x && \
    go test ./... && \ 
    go build -ldflags="-s -w" -o scds -v . && \
    upx scds

FROM alpine:3
LABEL Name=scds

RUN apk update && apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder ["/app/scds", "/app/*json", "/app/template.tpl", "./"]

ENTRYPOINT ["./scds"]
