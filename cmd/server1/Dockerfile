FROM golang:1.25.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server1 ./cmd/server1

FROM alpine:latest

COPY --from=builder /app/server1 /usr/local/bin/server1

CMD ["server1"]


