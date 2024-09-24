FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./users ./cmd/api/main.go

FROM alpine:3.20.3

WORKDIR /build
COPY --from=builder /app/users ./users
COPY --from=builder /app/swagger ./swagger
EXPOSE 8080
EXPOSE 8081
EXPOSE 50051
CMD ["/build/users"]