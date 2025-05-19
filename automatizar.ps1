Write-Host "📦 Construyendo imagen del microservicio..."
docker-compose build

Write-Host "`n🚀 Levantando contenedores..."
docker-compose up -d

Write-Host "`n✅ Microservicio y Mongo están corriendo."
Write-Host "🌐 Accede a: http://localhost:8080"
