# ---- Etapa 1: Builder (Compilación) ----
FROM golang:latest AS builder

# Establecer el directorio de trabajo
WORKDIR /src

# Copiar go.mod y go.sum primero para aprovechar la caché de capas
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar la aplicación desde la carpeta cmd/api
# Nota: Asegúrate de que el paquete en main.go sea "package main"
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/barberia-cola-virtual ./cmd/api

# ---- Etapa 2: Runner (Imagen final ligera) ----
FROM alpine:3.20

# Instalar certificados (importante si tu API hace peticiones HTTPS)
RUN apk add --no-cache ca-certificates

# Crear un usuario no root por seguridad
RUN adduser -D -u 10001 appuser
USER appuser

# Copiar el binario desde la etapa de compilación
COPY --from=builder /bin/barberia-cola-virtual /app/barberia-cola-virtual

# Exponer el puerto (cambia 8080 si tu API usa otro)
EXPOSE 8080

# Comando para ejecutar el binario
ENTRYPOINT ["/app/barberia-cola-virtual"]