# Stage 1: Build
FROM golang:1.21-alpine AS builder
WORKDIR /app

# Copia tutto il contenuto della cartella
COPY . .

# Se il file go.mod manca, lo creiamo al volo dentro il container
RUN if [ ! -f go.mod ]; then go mod init logistics/analytics; fi

RUN go build -o main .

# Stage 2: Final
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]