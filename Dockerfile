# =====================================================
# DOCKERFILE MULTI-STAGE
# Proyecto: Barbería Cola Virtual
#
# Objetivo:
# 1. Compilar la API en una imagen con Go.
# 2. Ejecutar solo el binario final en una imagen liviana.
# =====================================================

# ============================
# ETAPA 1: BUILDER
# Esta etapa tiene Go instalado y sirve solo para compilar.
# ============================
FROM golang:1.26-alpine AS builder

# Carpeta de trabajo dentro del contenedor
WORKDIR /app

# Copiamos primero go.mod y go.sum para aprovechar la caché de Docker.
# Si las dependencias no cambian, Docker no las descarga otra vez.
COPY go.mod go.sum ./

# Descarga las dependencias del proyecto
RUN go mod download

# Copia todo el código fuente al contenedor
COPY . .

# Compila la API.
# CGO_ENABLED=0 genera un binario estático, más fácil de ejecutar en Alpine.
RUN CGO_ENABLED=0 go build -o /bin/api ./cmd/api

# ============================
# ETAPA 2: RUNNER
# Esta etapa es mínima. No tiene Go instalado.
# Solo contiene el ejecutable final.
# ============================
FROM alpine:3.20

# Carpeta donde correrá la aplicación
WORKDIR /app

# Copiamos el binario generado en la etapa builder
COPY --from=builder /bin/api /app/api

# Copiamos la carpeta web para servir HTML/CSS/JS si el proyecto la usa
# COPY web /app/web

# Puerto usado por la API
EXPOSE 8080

# Comando que se ejecuta al iniciar el contenedor
ENTRYPOINT ["/app/api"]