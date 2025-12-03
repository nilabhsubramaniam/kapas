# Tantuka Backend Setup Script for PowerShell
# Run this script after installing Go and creating PostgreSQL database

Write-Host "🚀 Tantuka Backend Setup" -ForegroundColor Cyan
Write-Host "=========================" -ForegroundColor Cyan
Write-Host ""

# Check if Go is installed
Write-Host "Checking Go installation..." -ForegroundColor Yellow
$goVersion = & go version 2>$null
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Go is not installed or not in PATH" -ForegroundColor Red
    Write-Host "Please install Go from: https://go.dev/dl/" -ForegroundColor Yellow
    Write-Host "After installation, restart PowerShell and run this script again" -ForegroundColor Yellow
    exit 1
}
Write-Host "✅ $goVersion" -ForegroundColor Green
Write-Host ""

# Navigate to project directory
$projectPath = "c:\Users\Nilabh\Projects\kapas"
Set-Location $projectPath

# Download Go dependencies
Write-Host "📦 Downloading Go dependencies..." -ForegroundColor Yellow
go mod download
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Failed to download dependencies" -ForegroundColor Red
    exit 1
}
Write-Host "✅ Dependencies downloaded" -ForegroundColor Green
Write-Host ""

# Tidy up go.mod
Write-Host "📝 Tidying go.mod..." -ForegroundColor Yellow
go mod tidy
Write-Host "✅ go.mod updated" -ForegroundColor Green
Write-Host ""

# Check if .env exists
if (!(Test-Path ".env")) {
    Write-Host "⚙️  Creating .env file from template..." -ForegroundColor Yellow
    Copy-Item ".env.example" ".env"
    Write-Host "✅ .env file created" -ForegroundColor Green
    Write-Host "⚠️  Please edit .env file with your database credentials" -ForegroundColor Yellow
    Write-Host ""
}

# Ask about database configuration
Write-Host "Database Configuration:" -ForegroundColor Cyan
$dbHost = Read-Host "Enter DATABASE_HOST (default: localhost)"
if ([string]::IsNullOrWhiteSpace($dbHost)) { $dbHost = "localhost" }

$dbPort = Read-Host "Enter DATABASE_PORT (default: 5432)"
if ([string]::IsNullOrWhiteSpace($dbPort)) { $dbPort = "5432" }

$dbUser = Read-Host "Enter DATABASE_USER (default: postgres)"
if ([string]::IsNullOrWhiteSpace($dbUser)) { $dbUser = "postgres" }

$dbPassword = Read-Host "Enter DATABASE_PASSWORD" -AsSecureString
$dbPasswordPlain = [Runtime.InteropServices.Marshal]::PtrToStringAuto(
    [Runtime.InteropServices.Marshal]::SecureStringToBSTR($dbPassword)
)

$dbName = Read-Host "Enter DATABASE_NAME (default: tantuka_db)"
if ([string]::IsNullOrWhiteSpace($dbName)) { $dbName = "tantuka_db" }

# Update .env file
Write-Host ""
Write-Host "📝 Updating .env file with database credentials..." -ForegroundColor Yellow

$envContent = Get-Content ".env"
$envContent = $envContent -replace "DATABASE_HOST=.*", "DATABASE_HOST=$dbHost"
$envContent = $envContent -replace "DATABASE_PORT=.*", "DATABASE_PORT=$dbPort"
$envContent = $envContent -replace "DATABASE_USER=.*", "DATABASE_USER=$dbUser"
$envContent = $envContent -replace "DATABASE_PASSWORD=.*", "DATABASE_PASSWORD=$dbPasswordPlain"
$envContent = $envContent -replace "DATABASE_NAME=.*", "DATABASE_NAME=$dbName"
$envContent | Set-Content ".env"

Write-Host "✅ .env file updated" -ForegroundColor Green
Write-Host ""

# Build the application
Write-Host "🔨 Building application..." -ForegroundColor Yellow
go build -o tantuka-backend.exe cmd/server/main.go
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Build failed" -ForegroundColor Red
    exit 1
}
Write-Host "✅ Build successful" -ForegroundColor Green
Write-Host ""

# Summary
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "✅ Setup Complete!" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "Next steps:" -ForegroundColor Yellow
Write-Host "1. Ensure PostgreSQL database '$dbName' exists" -ForegroundColor White
Write-Host "2. Run the backend server:" -ForegroundColor White
Write-Host "   go run cmd/server/main.go" -ForegroundColor Cyan
Write-Host "   OR" -ForegroundColor White
Write-Host "   .\tantuka-backend.exe" -ForegroundColor Cyan
Write-Host ""
Write-Host "3. Access the API at: http://localhost:8080" -ForegroundColor White
Write-Host "4. API Documentation: http://localhost:8080/swagger/index.html" -ForegroundColor White
Write-Host ""
Write-Host "Optional - Start Docker services:" -ForegroundColor Yellow
Write-Host "   docker-compose up -d" -ForegroundColor Cyan
Write-Host ""

# Ask if user wants to start the server
$startServer = Read-Host "Do you want to start the server now? (y/n)"
if ($startServer -eq "y" -or $startServer -eq "Y") {
    Write-Host ""
    Write-Host "🚀 Starting Tantuka Backend..." -ForegroundColor Green
    Write-Host "Press Ctrl+C to stop the server" -ForegroundColor Yellow
    Write-Host ""
    go run cmd/server/main.go
}
