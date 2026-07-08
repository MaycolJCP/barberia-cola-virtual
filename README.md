# 💈 Barbería Cola Virtual

## Descripción

Barbería Cola Virtual es una API REST desarrollada en Go que permite administrar una barbería mediante una cola virtual. El sistema facilita la gestión de usuarios, autenticación, catálogo de servicios y turnos, utilizando una arquitectura por capas y persistencia en PostgreSQL.

---

## Integrantes

| Integrante     | Módulo |
|------------    |--------|
| Danny Zambrano | Catálogo de Servicios |
| (Nombre compañero) | Autenticación y Usuarios |
| (Nombre compañero) | Turnos y Seguimiento |

---

## Tecnologías utilizadas

- Go
- Chi Router
- GORM
- PostgreSQL
- Docker
- Docker Compose
- JWT
- GitHub Actions
- Postman

---

## Arquitectura del proyecto

El proyecto utiliza una arquitectura en capas:

```
Cliente (Postman)
        │
        ▼
     Handler
        │
        ▼
     Service
        │
        ▼
    Repository
        │
        ▼
       GORM
        │
        ▼
   PostgreSQL
```

---

## Estructura del proyecto

```
cmd/
internal/
    handlers/
    services/
    repository/
    models/
    middleware/
    db/
    utils/
.github/workflows/
Dockerfile
docker-compose.yml
```

---

## Funcionalidades

### Autenticación

- Registro de usuarios
- Inicio de sesión
- JWT
- Middleware de autenticación
- Roles ADMIN y CLIENTE

### Catálogo de Servicios

- Crear servicios
- Consultar servicios
- Actualizar servicios
- Eliminar servicios
- Categorías
- Promociones

### Turnos

- Crear turnos
- Consultar turnos
- Actualizar turnos
- Eliminar turnos

---

## Ejecución con Docker

### Levantar el proyecto

```bash
docker compose up -d --build
```

### Detener contenedores

```bash
docker compose down
```

### Eliminar completamente la base de datos

```bash
docker compose down -v
```

---

## Endpoints principales

### Autenticación

| Método | Endpoint |
|---------|----------|
| POST | /api/v1/auth/register |
| POST | /api/v1/auth/login |

### Servicios

| Método | Endpoint |
|---------|----------|
| GET | /api/v1/servicios |
| POST | /api/v1/servicios |
| PUT | /api/v1/servicios/{id} |
| DELETE | /api/v1/servicios/{id} |

### Categorías

| Método | Endpoint |
|---------|----------|
| GET | /api/v1/categorias-servicio |
| POST | /api/v1/categorias-servicio |
| PUT | /api/v1/categorias-servicio/{id} |
| DELETE | /api/v1/categorias-servicio/{id} |

### Promociones

| Método | Endpoint |
|---------|----------|
| GET | /api/v1/promociones |
| POST | /api/v1/promociones |
| PUT | /api/v1/promociones/{id} |
| DELETE | /api/v1/promociones/{id} |

### Turnos

| Método | Endpoint |
|---------|----------|
| GET | /api/v1/turnos |
| POST | /api/v1/turnos |
| PUT | /api/v1/turnos/{id} |
| DELETE | /api/v1/turnos/{id} |

---

## Pruebas

El proyecto cuenta con:

- Pruebas unitarias.
- GitHub Actions.
- Integración continua.
- Docker Compose.
- PostgreSQL.

---

## Evidencias

- Docker funcionando.
- PostgreSQL conectado mediante GORM.
- JWT funcionando.
- Roles ADMIN y CLIENTE.
- Pipeline de GitHub Actions en verde.