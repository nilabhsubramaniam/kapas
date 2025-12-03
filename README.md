# 🛍️ Tantuka E-Commerce Backend (Kapas)

**Premium Indian Saree E-Commerce Platform - Go Backend**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18-336791?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

> A production-grade, scalable e-commerce backend built with Go (Golang), designed to handle the complete order lifecycle from browsing to delivery for Tantuka - a premium saree marketplace.

---

## 📋 Table of Contents

- [Project Overview](#-project-overview)
- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Project Structure](#-project-structure)
- [Getting Started](#-getting-started)
- [Environment Setup](#-environment-setup)
- [Database Schema](#-database-schema)
- [API Documentation](#-api-documentation)
- [Development Progress](#-development-progress)
- [Deployment](#-deployment)

---

## 🎯 Project Overview

**Kapas** (meaning "cotton" in Hindi) is the backend powerhouse for **Tantuka**, an e-commerce platform specializing in premium Indian sarees from all 28+ states. This backend handles:

- ✅ Complete order lifecycle (cart → checkout → payment → shipping → delivery)
- ✅ Multi-warehouse inventory management
- ✅ Real-time logistics tracking
- ✅ Payment gateway integration (Razorpay)
- ✅ Admin panel operations
- ✅ Analytics and reporting
- ✅ Returns and refunds management

### Frontend Repository
- **Location**: [Tantuka Frontend](https://github.com/nilabhsubramaniam/Tantuka)
- **Tech**: Next.js 14, React 18, Tailwind CSS
- **Live**: [tantuka.com](https://nilabhsubramaniam.github.io/Tantuka/)

---

## ✨ Features

### Core E-Commerce
- 🔐 **Authentication & Authorization** - JWT-based with role management (Customer, Admin, Vendor)
- 🛒 **Shopping Cart** - Session-based cart with Redis caching
- 💳 **Payment Processing** - Razorpay integration with webhook support
- 📦 **Order Management** - Complete order lifecycle with status tracking
- 🏪 **Product Catalog** - Advanced filtering, search, and pagination
- ⭐ **Reviews & Ratings** - Customer feedback system
- 💝 **Wishlist** - Save favorite products

### Logistics & Fulfillment
- 📍 **Multi-warehouse Support** - Inventory across multiple locations
- 🚚 **Shipping Integration** - Delhivery, Blue Dart, Shiprocket APIs
- 📱 **Real-time Tracking** - Live shipment status updates
- 🔄 **Returns Management** - Automated return and refund processing
- 📊 **Inventory Alerts** - Low stock notifications

### Admin Operations
- 📈 **Analytics Dashboard** - Sales, revenue, conversion metrics
- 👥 **User Management** - Customer and admin user controls
- 🎟️ **Coupon System** - Discount codes and promotions
- 📧 **Notification System** - Email/SMS for order updates
- 📝 **Activity Logging** - Audit trail for all operations
- 📄 **Report Generation** - Sales reports, inventory reports

### Advanced Features
- 🔍 **Full-text Search** - Fast product search
- 🌍 **Multi-state Support** - State-wise product categorization
- 💰 **Dynamic Pricing** - Discount calculations, tax management
- 🔒 **Security** - Rate limiting, CORS, input validation
- 📱 **Mobile Optimized** - REST API for mobile apps

---

## 🛠️ Tech Stack

### Backend Framework
- **Language**: Go 1.21+
- **Web Framework**: [Gin](https://gin-gonic.com/) - Fast HTTP router
- **ORM**: [GORM](https://gorm.io/) - Object-relational mapping

### Database & Caching
- **Primary Database**: PostgreSQL 18
- **Cache**: Redis 7 (sessions, cart, rate limiting)
- **Search**: PostgreSQL Full-text Search (Elasticsearch in future)

### External Services
- **Payment**: Razorpay SDK
- **Email**: SendGrid / AWS SES
- **SMS**: Twilio / AWS SNS
- **Storage**: AWS S3 / Cloudinary (product images)
- **Logistics**: Shiprocket, Delhivery, Blue Dart APIs

### DevOps & Tools
- **Containerization**: Docker & Docker Compose
- **API Docs**: Swagger/OpenAPI 3.0
- **Logging**: Zap (structured logging)
- **Monitoring**: Prometheus + Grafana (future)
- **CI/CD**: GitHub Actions

---

## 📁 Project Structure

```
kapas/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
│
├── internal/                       # Private application code
│   ├── config/
│   │   ├── database.go            # Database connection
│   │   ├── redis.go               # Redis connection
│   │   └── environment.go         # Environment variables
│   │
│   ├── models/                     # GORM database models
│   │   ├── user.go                # User, Admin models
│   │   ├── product.go             # Product, ProductImage
│   │   ├── order.go               # Order, OrderItem, OrderStatusHistory
│   │   ├── payment.go             # Payment transactions
│   │   ├── shipping.go            # Shipment, TrackingEvents
│   │   ├── inventory.go           # Warehouse, Inventory
│   │   ├── address.go             # Customer addresses
│   │   ├── coupon.go              # Coupons, CouponUsage
│   │   ├── notification.go        # Notifications
│   │   └── return.go              # Returns, Refunds
│   │
│   ├── handlers/                   # HTTP request handlers
│   │   ├── auth.go                # Authentication endpoints
│   │   ├── product.go             # Product CRUD
│   │   ├── cart.go                # Cart operations
│   │   ├── order.go               # Order management
│   │   ├── payment.go             # Payment processing
│   │   ├── shipping.go            # Shipping & tracking
│   │   ├── admin.go               # Admin operations
│   │   └── analytics.go           # Analytics endpoints
│   │
│   ├── services/                   # Business logic layer
│   │   ├── auth_service.go
│   │   ├── order_service.go
│   │   ├── payment_service.go
│   │   ├── shipping_service.go
│   │   ├── inventory_service.go
│   │   └── notification_service.go
│   │
│   ├── repository/                 # Data access layer
│   │   ├── user_repo.go
│   │   ├── product_repo.go
│   │   ├── order_repo.go
│   │   └── inventory_repo.go
│   │
│   ├── middleware/                 # HTTP middleware
│   │   ├── auth.go                # JWT authentication
│   │   ├── cors.go                # CORS handling
│   │   ├── logger.go              # Request logging
│   │   ├── rate_limit.go          # Rate limiting
│   │   └── role.go                # Role-based access
│   │
│   └── utils/                      # Helper functions
│       ├── jwt.go                 # JWT utilities
│       ├── validator.go           # Input validation
│       ├── pagination.go          # Pagination helper
│       └── helpers.go             # General utilities
│
├── pkg/                            # Public packages (reusable)
│   ├── razorpay/                  # Payment gateway SDK
│   ├── logistics/                 # Shipping APIs
│   └── cache/                     # Redis caching
│
├── migrations/                     # Database migrations
│   ├── 001_create_users.sql
│   ├── 002_create_products.sql
│   └── ...
│
├── docs/                           # API documentation (Swagger)
│   └── swagger.json
│
├── tests/                          # Test files
│   ├── integration/
│   └── unit/
│
├── .env.example                    # Environment variables template
├── .gitignore                      # Git ignore rules
├── docker-compose.yml              # Docker services setup
├── Dockerfile                      # Production Docker image
├── go.mod                          # Go dependencies
├── go.sum                          # Dependency checksums
└── README.md                       # This file
```

---

## 🚀 Getting Started

### Prerequisites

1. **Install Go** (1.21 or higher)
   ```powershell
   # Download from: https://go.dev/dl/
   # Verify installation
   go version
   ```

2. **Install PostgreSQL 18**
   - Already installed (pgAdmin available)
   - Or use Docker (see Docker Setup below)

3. **Install Git**
   ```powershell
   git --version
   ```

### Installation

1. **Clone the repository**
   ```powershell
   cd c:\Users\Nilabh\Projects
   git clone https://github.com/nilabhsubramaniam/kapas.git
   cd kapas
   ```

2. **Install Go dependencies**
   ```powershell
   go mod download
   go mod tidy
   ```

3. **Setup environment variables**
   ```powershell
   # Copy example env file
   Copy-Item .env.example .env
   
   # Edit .env file with your settings
   notepad .env
   ```

4. **Start PostgreSQL & Redis (using Docker)**
   ```powershell
   docker-compose up -d postgres redis pgadmin
   ```

   **Access pgAdmin**: http://localhost:5050
   - Email: `admin@tantuka.com`
   - Password: `admin`

5. **Run database migrations**
   ```powershell
   # Auto-migration will run on first server start
   # Or manually run migrations
   go run cmd/server/main.go
   ```

6. **Start the server**
   ```powershell
   # Development mode
   go run cmd/server/main.go
   
   # Or build and run
   go build -o tantuka-backend.exe cmd/server/main.go
   .\tantuka-backend.exe
   ```

7. **Access the API**
   - API Base URL: http://localhost:8080
   - Swagger Docs: http://localhost:8080/swagger/index.html
   - Health Check: http://localhost:8080/api/health

---

## ⚙️ Environment Setup

### Local Development (.env)

```env
# Database
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=tantuka_user
DATABASE_PASSWORD=your_password
DATABASE_NAME=tantuka_db

# Application
PORT=8080
APP_ENV=development
GIN_MODE=debug

# JWT
JWT_SECRET=your-secret-key-change-in-production
JWT_EXPIRATION_HOURS=24

# CORS
CORS_ORIGINS=http://localhost:3000
```

### Docker Setup (Quick Start)

```powershell
# Start all services (PostgreSQL, Redis, pgAdmin)
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Reset database (WARNING: Deletes all data)
docker-compose down -v
docker-compose up -d
```

---

## 🗄️ Database Schema

### Core Tables (25+ tables)

#### Authentication & Users
- `users` - Customer accounts
- `admin_users` - Admin hierarchy
- `addresses` - Shipping/billing addresses

#### Products & Catalog
- `products` - Product catalog
- `product_images` - Product photos
- `categories` - Product categories (state-wise, fabric, occasion)
- `reviews` - Customer reviews
- `wishlists` - Saved products

#### Orders & Payments
- `cart` - Shopping cart items
- `orders` - Order header
- `order_items` - Order line items
- `order_status_history` - Status tracking
- `payments` - Payment transactions
- `coupons` - Discount codes
- `coupon_usage` - Coupon tracking
- `taxes` - Tax rules

#### Logistics & Fulfillment
- `warehouses` - Warehouse locations
- `inventory` - Stock per warehouse
- `shipments` - Shipping records
- `tracking_events` - Shipment tracking
- `logistics_providers` - Delivery partners
- `returns` - Return requests
- `return_items` - Returned products

#### System & Admin
- `notifications` - User notifications
- `activity_logs` - Audit trail
- `settings` - System configuration
- `analytics_events` - User behavior
- `sales_reports` - Aggregated data

### ER Diagram
*(Will be added after schema finalization)*

---

## 📚 API Documentation

### Authentication Endpoints

```http
POST   /api/auth/register      # Register new user
POST   /api/auth/login         # Login (returns JWT)
GET    /api/auth/me            # Get current user (protected)
POST   /api/auth/logout        # Logout
POST   /api/auth/refresh       # Refresh JWT token
```

### Product Endpoints

```http
GET    /api/products           # List products (with filters)
GET    /api/products/:slug     # Get single product
GET    /api/products/state/:state  # Get by state
POST   /api/products           # Create product (admin)
PUT    /api/products/:id       # Update product (admin)
DELETE /api/products/:id       # Delete product (admin)
```

### Cart Endpoints

```http
GET    /api/cart               # Get user's cart
POST   /api/cart/items         # Add to cart
PUT    /api/cart/items/:id     # Update quantity
DELETE /api/cart/items/:id     # Remove from cart
```

### Order Endpoints

```http
POST   /api/orders             # Create order (checkout)
GET    /api/orders             # List user's orders
GET    /api/orders/:id         # Get order details
GET    /api/orders/:id/track   # Track shipment
PUT    /api/orders/:id/cancel  # Cancel order
```

### Admin Endpoints

```http
GET    /api/admin/dashboard    # Analytics dashboard
GET    /api/admin/orders       # All orders
PUT    /api/admin/orders/:id/status  # Update status
GET    /api/admin/inventory    # Inventory management
GET    /api/admin/analytics/sales    # Sales reports
```

**Full API Documentation**: http://localhost:8080/swagger/index.html (when server is running)

---

## 📊 Development Progress

### ✅ Phase 1: Foundation (Week 1-2)
- [x] Project initialization
- [x] Go module setup
- [x] Directory structure
- [x] Environment configuration
- [x] Docker Compose setup
- [ ] Database connection
- [ ] GORM models (25+ tables)
- [ ] Auto-migration

### 🔄 Phase 2: Core APIs (Week 2-3)
- [ ] Authentication (JWT)
- [ ] User registration/login
- [ ] Product CRUD
- [ ] Category management
- [ ] Cart operations
- [ ] Filtering & pagination

### 📋 Phase 3: E-Commerce Features (Week 3-4)
- [ ] Order creation
- [ ] Payment integration (Razorpay)
- [ ] Address management
- [ ] Coupon system
- [ ] Review system
- [ ] Wishlist

### 🚚 Phase 4: Logistics (Week 4-5)
- [ ] Inventory management
- [ ] Warehouse system
- [ ] Shipping integration
- [ ] Order tracking
- [ ] Returns & refunds

### 👨‍💼 Phase 5: Admin Panel (Week 5-6)
- [ ] Admin dashboard
- [ ] Order management
- [ ] Inventory alerts
- [ ] Analytics endpoints
- [ ] User management
- [ ] Activity logs

### 🚀 Phase 6: Deployment (Week 6)
- [ ] Swagger documentation
- [ ] Unit tests
- [ ] Integration tests
- [ ] Docker production image
- [ ] CI/CD pipeline
- [ ] Production deployment

---

## 🌐 Deployment

### Production Deployment Options

#### Option 1: Railway (Recommended)
```powershell
# Install Railway CLI
npm install -g @railway/cli

# Login and deploy
railway login
railway init
railway up
```

#### Option 2: Docker Deployment
```powershell
# Build production image
docker build -t tantuka-backend .

# Run container
docker run -p 8080:8080 --env-file .env tantuka-backend
```

#### Option 3: AWS EC2
- Setup EC2 instance (t2.medium or higher)
- Install Go, PostgreSQL
- Configure nginx reverse proxy
- Setup systemd service
- Enable SSL (Let's Encrypt)

### Environment Variables (Production)

```env
APP_ENV=production
GIN_MODE=release
DATABASE_SSL_MODE=require
CORS_ORIGINS=https://tantuka.com
# Add production API keys
```

---

## 🧪 Testing

```powershell
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/services/...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

---

## 📝 Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open Pull Request

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 👥 Team

- **Developer**: Nilabh Subramaniam
- **GitHub**: [@nilabhsubramaniam](https://github.com/nilabhsubramaniam)
- **Project**: Tantuka E-Commerce Platform

---

## 🔗 Related Repositories

- **Frontend**: [Tantuka](https://github.com/nilabhsubramaniam/Tantuka) - Next.js 14 frontend
- **Admin Panel**: *(Coming soon)*
- **Mobile App**: *(Future)*

---

## 📞 Support

For questions or issues:
- Open an issue on GitHub
- Email: support@tantuka.com
- Documentation: [Wiki](https://github.com/nilabhsubramaniam/kapas/wiki)

---

## 🙏 Acknowledgments

- Go community
- Gin framework team
- GORM team
- All open-source contributors

---

**Built with ❤️ for promoting Indian heritage craftsmanship**

*Last Updated: December 3, 2025*
