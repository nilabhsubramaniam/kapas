# 🚀 Quick Start Guide - Tantuka Backend

## Prerequisites Checklist

Before starting, ensure you have:

- [ ] **Go 1.21+** installed ([Download](https://go.dev/dl/))
- [ ] **PostgreSQL 18** installed (or Docker)
- [ ] **Git** installed
- [ ] **pgAdmin** (already installed on your system)
- [ ] **Docker Desktop** (optional, for containerized setup)

---

## 🎯 Setup Options

### Option A: Quick Docker Setup (Recommended for Beginners)

```powershell
# 1. Navigate to project
cd c:\Users\Nilabh\Projects\kapas

# 2. Start PostgreSQL, Redis, and pgAdmin
docker-compose up -d

# 3. Copy environment file
Copy-Item .env.example .env

# 4. Install Go dependencies (requires Go installed)
go mod download

# 5. Run the server
go run cmd/server/main.go
```

**Access Points**:
- API: http://localhost:8080
- pgAdmin: http://localhost:5050 (email: `admin@tantuka.com`, password: `admin`)
- Swagger Docs: http://localhost:8080/swagger/index.html

---

### Option B: Manual PostgreSQL Setup

```powershell
# 1. Create database using pgAdmin or psql
# Open pgAdmin → Create New Database: "tantuka_db"

# 2. Copy environment file
Copy-Item .env.example .env

# 3. Edit .env with your PostgreSQL credentials
notepad .env

# Update these lines:
# DATABASE_HOST=localhost
# DATABASE_PORT=5432
# DATABASE_USER=postgres
# DATABASE_PASSWORD=your_postgres_password
# DATABASE_NAME=tantuka_db

# 4. Install dependencies
go mod download

# 5. Run the server
go run cmd/server/main.go
```

---

## 📋 Installation Steps (Detailed)

### Step 1: Install Go

**Windows**:
1. Download installer from https://go.dev/dl/
2. Run installer (go1.21.x.windows-amd64.msi)
3. Restart PowerShell/Terminal
4. Verify: `go version`

**Expected Output**:
```
go version go1.21.x windows/amd64
```

### Step 2: Setup Database

**Using pgAdmin** (already installed):
1. Open pgAdmin
2. Connect to PostgreSQL server
3. Right-click "Databases" → Create → Database
4. Name: `tantuka_db`
5. Owner: `postgres` (or create new user `tantuka_user`)

**Using Docker** (alternative):
```powershell
# Start only PostgreSQL
docker-compose up -d postgres

# View logs
docker-compose logs -f postgres
```

### Step 3: Configure Environment

```powershell
# Copy template
Copy-Item .env.example .env

# Edit configuration
notepad .env
```

**Minimum Required Settings**:
```env
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASSWORD=YOUR_PASSWORD_HERE
DATABASE_NAME=tantuka_db
JWT_SECRET=change-this-to-random-string
PORT=8080
```

### Step 4: Install Dependencies

```powershell
# Navigate to project
cd c:\Users\Nilabh\Projects\kapas

# Download all Go packages
go mod download

# Verify dependencies
go mod tidy
```

### Step 5: Run the Application

```powershell
# Development mode (with hot reload - install Air first)
# go install github.com/cosmtrek/air@latest
# air

# Standard run
go run cmd/server/main.go

# Or build executable
go build -o tantuka-backend.exe cmd/server/main.go
.\tantuka-backend.exe
```

**Expected Output**:
```
Database connected successfully
Database migration completed
Server starting on port 8080
[GIN-debug] Listening and serving HTTP on :8080
```

---

## ✅ Verify Installation

### 1. Check API Health
```powershell
# Using curl (install from: https://curl.se/windows/)
curl http://localhost:8080/api/health

# Or open in browser
start http://localhost:8080/api/health
```

Expected Response:
```json
{
  "status": "ok",
  "timestamp": "2025-12-03T12:00:00Z"
}
```

### 2. Check Swagger Documentation
```powershell
start http://localhost:8080/swagger/index.html
```

### 3. Test Database Connection

Open pgAdmin → Query Tool:
```sql
-- Check if tables are created
SELECT table_name 
FROM information_schema.tables 
WHERE table_schema = 'public';
```

---

## 🛠️ Development Workflow

### Daily Development

```powershell
# 1. Start Docker services (if using Docker)
docker-compose up -d

# 2. Run the server
go run cmd/server/main.go

# 3. Make code changes (server will need restart)

# 4. Stop services when done
docker-compose down
```

### With Hot Reload (Air)

```powershell
# Install Air once
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

Creates `.air.toml` config automatically. Now server restarts on file changes!

---

## 📝 Common Commands

### Go Commands
```powershell
# Run application
go run cmd/server/main.go

# Build executable
go build -o tantuka-backend.exe cmd/server/main.go

# Run tests
go test ./...

# Format code
go fmt ./...

# Install new package
go get github.com/package/name

# Update dependencies
go mod tidy
```

### Docker Commands
```powershell
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop all services
docker-compose down

# Restart a service
docker-compose restart postgres

# View running containers
docker-compose ps

# Remove all data (WARNING: Deletes database!)
docker-compose down -v
```

### Database Commands (psql)
```powershell
# Connect to database
docker exec -it tantuka-postgres psql -U tantuka_user -d tantuka_db

# Or using local psql
psql -U postgres -d tantuka_db

# Inside psql:
\dt              # List tables
\d users         # Describe users table
\q               # Quit
```

---

## 🐛 Troubleshooting

### Issue: "Go is not recognized"
**Solution**: 
1. Ensure Go is installed
2. Restart terminal/PowerShell
3. Check PATH: `$env:PATH -split ';'` should include Go path

### Issue: "Cannot connect to database"
**Solution**:
1. Check if PostgreSQL is running: `docker-compose ps`
2. Verify credentials in `.env`
3. Check firewall settings (port 5432)
4. Try connecting via pgAdmin first

### Issue: "Port 8080 already in use"
**Solution**:
```powershell
# Find process using port 8080
netstat -ano | findstr :8080

# Kill process (replace PID)
taskkill /PID <PID> /F

# Or change PORT in .env
```

### Issue: "Module not found"
**Solution**:
```powershell
go mod download
go mod tidy
```

### Issue: Docker services won't start
**Solution**:
```powershell
# Check Docker Desktop is running
docker --version

# Reset and restart
docker-compose down
docker-compose up -d
```

---

## 📚 Next Steps

After successful setup:

1. ✅ Test API endpoints using Swagger UI
2. ✅ Review database schema in pgAdmin
3. ✅ Read `README.md` for full documentation
4. ✅ Check `PROGRESS.md` for development roadmap
5. ✅ Start building features!

---

## 🔗 Useful Resources

- **Go Documentation**: https://go.dev/doc/
- **Gin Framework**: https://gin-gonic.com/docs/
- **GORM**: https://gorm.io/docs/
- **PostgreSQL**: https://www.postgresql.org/docs/
- **Docker**: https://docs.docker.com/

---

## 📞 Need Help?

- Check `README.md` for detailed information
- Review `PROGRESS.md` for current status
- Open an issue on GitHub
- Check error logs in terminal

---

**Happy Coding! 🚀**

*Last Updated: December 3, 2025*
