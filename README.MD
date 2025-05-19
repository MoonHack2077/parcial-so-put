# 🧪 Microservicio de Tareas (Parcial 2 - Ingeniería de Software)

Microservicio desarrollado en Golang con MongoDB, contenerizado con Docker, probado con Thunder Client y cubierto con pruebas unitarias usando Testify. Forma parte del parcial 2.

---

## 🚀 Funcionalidad

CRUD completo sobre la entidad `Tarea`, que incluye:

- [x] Crear tarea (POST `/tareas`)
- [x] Obtener tareas (GET `/tareas`)
- [x] Actualizar tarea por ID (PUT `/tareas/{id}`)
- [x] Eliminar tarea por ID (DELETE `/tareas/{id}`)

---

## 📦 Tecnologías utilizadas

- Golang
- MongoDB
- Docker y Docker Compose
- Thunder Client (VSCode)
- Testify para pruebas unitarias
- Godotenv para manejar entornos

---

## 🛠️ Instrucciones para ejecutar

### 1. Clonar el repositorio y cargar variables

```bash
git clone https://github.com/MoonHack2077/parcial-so-put.git
cd parcial-so-put
```

Crear un archivo `.env` con:

```env
ENV=prod
MONGO_URI=mongodb://mongo:27017
MONGO_URI_TEST=mongodb://localhost:27017
MONGO_DB=tareasdb
PORT=8080
```

### 2. Ejecutar el microservicio con Docker

```bash
# O usar el script automatizado:
./automatizar.ps1     # Windows PowerShell

# o si estas en Linux o WSL
# Asignar permisos de ejecución (solo una vez):
chmod +x automatizar.sh

# Ejecutar el script
./automatizar.sh
```

La app estará disponible en: [http://localhost:8080](http://localhost:8080)

---

## 🐳 Imagen en Docker Hub

La imagen del microservicio fue construida y publicada en Docker Hub. Puede ser ejecutada en cualquier entorno compatible con Docker:

🔗 **https://hub.docker.com/r/juanma1301/parcial-so-put**

### Para ejecutarla directamente:

```bash
docker pull juanma1301/parcial-so-put
docker run -p 8080:8080 --env-file .env juanma1301/parcial-so-put
```

> Asegúrate de tener un archivo `.env` válido al momento de correr el contenedor.

---

## 🧪 Pruebas

### Pruebas unitarias

```bash
go test -cover ./...
```

- Cubre servicios: crear, obtener, actualizar y eliminar tareas

📊 Reporte de cobertura generado:
Archivo: coverage.out  
```bash
go test -coverprofile=coverage.out ./...
```

---

### Thunder Client (colección de pruebas)

- Incluye todos los endpoints
- Validaciones de estado, respuestas y estructura
- Archivo de colección exportado: `thunder-collection_parcial2.json`

---

## 🗄️ Backup de base de datos

Se genera backup dentro del contenedor Mongo con:

```powershell
docker exec -t mongo mongodump --archive="/data/db/backup-$(Get-Date -Format yyyy-MM-dd).gz" --gzip --db=tareasdb
docker cp mongo:/data/db/backup-2025-04-04.gz ./backup-2025-04-04.gz
```

---

## 📐 Diagrama de Infraestructura

Disponible en el archivo `infraestructura_microservicio.pdf`, representa:

- Cliente (Thunder Client)
- Contenedor app (Golang)
- Contenedor MongoDB
- Red Docker y volumen persistente

---

## 🎯 Entregables

- Código fuente con estructura modular
- Dockerfile y docker-compose.yml funcionales
- Archivo `.env` y scripts
- Colección Thunder Client exportada
- Pruebas unitarias y cobertura
- Backup generado desde MongoDB
- Diagrama de infraestructura
- BONUS: Script de automatización (`automatizar.ps1` / `automatizar.sh`)
- 🐳 Imagen publicada en Docker Hub: [juanma1301/parcial-so-put](https://hub.docker.com/r/juanma1301/parcial-so-put)

---

Desarrollado por **Juan Manuel** | Ingeniería de Software | 2025 ✅

