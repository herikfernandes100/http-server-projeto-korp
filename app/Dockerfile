# Etapa de build
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o http-server-projeto-korp

# Etapa final
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/http-server-projeto-korp .

EXPOSE 8080

CMD ["./http-server-projeto-korp"]