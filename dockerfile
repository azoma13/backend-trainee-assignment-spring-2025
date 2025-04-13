FROM golang:1.22.2-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/main cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/bin/main .
COPY --from=builder /app/.env .
CMD ["./main"]