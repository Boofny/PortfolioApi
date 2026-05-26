FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8000

CMD ["./server"]
