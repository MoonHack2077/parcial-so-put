#!/bin/bash

echo "📦 Construyendo imagen del microservicio..."
docker-compose build

echo ""
echo "🚀 Levantando contenedores..."
docker-compose up -d

echo ""
echo "✅ Microservicio y Mongo están corriendo."
echo "🌐 Accede a: http://localhost:8080"
