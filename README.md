# Proyecto de coding challenge

Este proyecto contiene **dos APIs** que se ejecutan en contenedores aislados a través de Docker Compose:

- ⚙️ **API en Go** – construida con Golang 1.18
- 🌐 **API en Node.js** – construida con Node.js 16

---

## 📁 Estructura del Proyecto

```
.
├── docker-compose.yml
├── go-api/
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── (Archivos fuente de Go)
└── node-api/
    ├── Dockerfile
    ├── package.json
    ├── package-lock.json
    └── (Archivos fuente de Node)
```

---

## 🚀 Empezando

### ✅ Requisitos previos

Asegúrate de tener instalados los siguientes elementos:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/) (generalmente incluido con Docker Desktop)

---

## 🔧 Ejecutando el Proyecto

1. **Clona el repositorio**:

   ```bash
   git clone https://github.com/pepi-coding/talent-house-coding-challenge.git
   cd talent-house-coding-challenge
   ```

2. **Inicia los servicios** usando Docker Compose:

   ```bash
   docker-compose up --build
   ```

   Esto hará lo siguiente:

   - Construir la API en Go desde la carpeta `go-api`
   - Construir la API en Node desde la carpeta `node-api`
   - Ejecutar ambos contenedores en la misma red virtual

3. **Accede a las APIs**:

   - API en Go: [http://localhost:3000](http://localhost:3000)
   - API en Node: [http://localhost:4000](http://localhost:4000)
   - Ruta de operación QR en Go: [http://localhost:3000](http://localhost:3000/api/v1/qr)
   - Ruta de operación de Operaciones en Node: [http://localhost:3000](http://localhost:4000/api/v1/matrix/process)

---

## 🛑 Detener los Servicios

Para detener los servicios:

```bash
docker-compose down
```

Esto detendrá y eliminará los contenedores, pero no eliminará las imágenes.

---
