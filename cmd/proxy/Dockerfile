FROM golang:1.25.5 AS builder


WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o proxy ./cmd/proxy

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/proxy /usr/local/bin/proxy
COPY --from=builder /app/web /app/web

CMD ["proxy"]


