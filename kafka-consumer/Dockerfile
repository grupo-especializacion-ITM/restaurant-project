# Etapa 1: build
FROM golang:1.23-bullseye AS builder

# Instalar librerías necesarias
RUN apt-get update && apt-get install -y \
    librdkafka-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src ./src

ENV CGO_ENABLED=1
ENV GOARCH=arm64
ENV GOOS=linux

# Compilar
RUN go build -o app ./src

# Etapa 2: final
FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y \
    librdkafka1 \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/app .

ENTRYPOINT ["./app"]
