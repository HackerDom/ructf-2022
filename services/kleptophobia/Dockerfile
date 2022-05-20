FROM golang:1.18.2-alpine3.15

COPY . /app
WORKDIR /app/server
RUN go build

ENTRYPOINT ["./server", "-config", "prod_config.json"]
