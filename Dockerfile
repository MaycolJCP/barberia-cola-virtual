# ---- Etapa 1: Builder (Compilación) ----
FROM golang:1.26.2-alpine AS builder

# Instalar git (necesario en alpine para descargar ciertas dependencias de Go)
RUN apk add --no-cache git

# Establecer el directorio de trabajo
WORKDIR /src

# Copiar go.mod y go.sum primero para aprovechar la caché de capas
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar la aplicación desde la carpeta cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/barberia-cola-virtual ./cmd/api

# ---- Etapa 2: Runner (Imagen final ligera) ----
FROM alpine:3.20

# Instalar certificados y base de datos de zonas horarias (tzdata)
RUN apk add --no-cache ca-certificates tzdata

# Crear un usuario no root por seguridad
RUN adduser -D -u 10001 appuser
USER appuser

# Copiar el binario desde la etapa de compilación
COPY --from=builder /bin/barberia-cola-virtual /app/barberia-cola-virtual

# Exponer el puerto
EXPOSE 8080

# Comando para ejecutar el binario
ENTRYPOINT ["/app/barberia-cola-virtual"]