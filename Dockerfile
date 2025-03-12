# Imagen base
FROM golang:1.24-alpine AS builder

# Configuración del entorno
WORKDIR /app

# Copiar archivos del proyecto
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilar la aplicación
RUN go build -o main cmd/main.go

# Imagen final
FROM alpine:latest
WORKDIR /root/

# Copiar el binario desde la imagen de compilación
COPY --from=builder /app/main .

# Exponer el puerto
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./main"]
