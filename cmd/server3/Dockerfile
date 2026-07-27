FROM golang:1.25.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o server2 ./cmd/server3

FROM alpine:latest

COPY --from=builder /app/server2 /usr/local/bin/server3

CMD ["server3"]




