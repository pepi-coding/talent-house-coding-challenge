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
   - Ruta de operación QR en Go: [http://localhost:3000/api/v1/qr](http://localhost:3000/api/v1/qr)

   ### Input (Body)

   ```bash
   {
       "matrix": [
           [1, 2],
           [3, 4],
           [5, 6]
       ]
   }
   ```

   ### Output (Response)

   ```bash
   {
       "Q": [
           [
               -0.16903085094570325,
               0.8970852271450604,
               0.4082482904638626
           ],
           [
               -0.50709255283711,
               0.2760262237369414,
               -0.816496580927726
           ],
           [
               -0.8451542547285166,
               -0.34503277967117696,
               0.4082482904638632
           ]
       ],
       "R": [
           [
               -5.916079783099615,
               -7.437357441610946
           ],
           [
               0,
               0.8280786712108259
           ],
           [
               0,
               0
           ]
       ],
       "stats": {
           "max": 0.8970852271450604,
           "min": -7.437357441610946,
           "average": -0.8812371693866827,
           "sum": -13.21855754080024,
           "isQDiagonal": false,
           "isRDiagonal": false
       }
   }
   ```

   - Ruta de operación de Operaciones en Node: [http://localhost:4000/api/v1/matrix/process](http://localhost:4000/api/v1/matrix/process)

---

## 🛑 Detener los Servicios

Para detener los servicios:

```bash
docker-compose down
```

Esto detendrá y eliminará los contenedores, pero no eliminará las imágenes.

---
