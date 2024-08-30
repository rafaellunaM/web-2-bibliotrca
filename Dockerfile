FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o main ./cmd/main.go

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/main .
COPY init.sql .

EXPOSE 8080

CMD ["./main"]
