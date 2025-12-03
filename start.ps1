# Quick Start Script - Run the Backend Server
# This script starts the Tantuka backend server

Write-Host "🚀 Starting Tantuka E-Commerce Backend..." -ForegroundColor Cyan
Write-Host ""

# Check if .env exists
if (!(Test-Path ".env")) {
    Write-Host "❌ .env file not found!" -ForegroundColor Red
    Write-Host "Please run setup.ps1 first:" -ForegroundColor Yellow
    Write-Host "   .\setup.ps1" -ForegroundColor Cyan
    exit 1
}

# Check if Go is installed
$goVersion = & go version 2>$null
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Go is not installed or not in PATH" -ForegroundColor Red
    Write-Host "Please restart PowerShell after installing Go" -ForegroundColor Yellow
    exit 1
}

Write-Host "✅ Environment: OK" -ForegroundColor Green
Write-Host "📍 API will be available at: http://localhost:8080" -ForegroundColor Yellow
Write-Host "📚 Swagger Docs: http://localhost:8080/swagger/index.html" -ForegroundColor Yellow
Write-Host ""
Write-Host "Press Ctrl+C to stop the server" -ForegroundColor Gray
Write-Host ""

# Run the server
go run cmd/server/main.go
