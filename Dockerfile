FROM golang:1.22.1-alpine3.18 AS builder
RUN apk add --no-cache git upx

WORKDIR /app

COPY ["src/*", "go.mod", "go.sum", "./"]
RUN go mod download -x && \
    go build -ldflags="-s -w" -o scds -v . && \
    upx scds

FROM alpine:3.19
LABEL Name=scds

RUN apk update && apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/scds .

ENTRYPOINT ["./scds"]
