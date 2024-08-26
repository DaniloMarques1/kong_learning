FROM golang:1.22.5

WORKDIR /app

COPY go.sum .
COPY go.mod .

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o todo-api

EXPOSE 5000

ENV QUEUE_URL="amqp://fitz:fitz@rabbitmq:5672"

ENTRYPOINT ["./todo-api"]
