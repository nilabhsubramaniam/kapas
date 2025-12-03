package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/nilabhsubramaniam/kapas/internal/models"
)

var DB *gorm.DB

// InitDatabase initializes the database connection and runs migrations
func InitDatabase() {
	var err error

	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		getEnv("DATABASE_HOST", "localhost"),
		getEnv("DATABASE_USER", "postgres"),
		getEnv("DATABASE_PASSWORD", ""),
		getEnv("DATABASE_NAME", "tantuka_db"),
		getEnv("DATABASE_PORT", "5432"),
		getEnv("DATABASE_SSL_MODE", "disable"),
	)

	// Configure GORM logger
	gormLogger := logger.Default
	if getEnv("APP_ENV", "development") == "production" {
		gormLogger = logger.Default.LogMode(logger.Silent)
	} else {
		gormLogger = logger.Default.LogMode(logger.Info)
	}

	// Open database connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Get underlying SQL database
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get database instance: %v", err)
	}

	// Set connection pool settings
	maxIdleConns, _ := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", "10"))
	maxOpenConns, _ := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", "100"))
	connMaxLifetime, _ := strconv.Atoi(getEnv("DB_CONN_MAX_LIFETIME", "3600"))

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	// Test connection
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("❌ Failed to ping database: %v", err)
	}

	log.Println("✅ Database connected successfully")

	// Run auto-migration
	if err := runMigrations(); err != nil {
		log.Fatalf("❌ Failed to run migrations: %v", err)
	}

	log.Println("✅ Database migrations completed")
}

// runMigrations runs GORM auto-migrations for all models
func runMigrations() error {
	return DB.AutoMigrate(
		// Location Hierarchy (NEW - must be first for foreign keys)
		&models.Country{},
		&models.State{},
		&models.District{},
		&models.Region{},

		// User & Auth
		&models.User{},
		&models.Address{},

		// Vendors (NEW)
		&models.Vendor{},

		// Products & Catalog
		&models.Product{},
		&models.ProductImage{},
		&models.Category{},
		&models.Review{},

		// Cart & Wishlist
		&models.CartItem{},
		&models.WishlistItem{},

		// Orders
		&models.Order{},
		&models.OrderItem{},
		&models.OrderStatusHistory{},

		// Payments
		&models.Payment{},
		&models.Coupon{},
		&models.CouponUsage{},

		// Inventory & Warehouses
		&models.Warehouse{},
		&models.Inventory{},

		// Shipping & Logistics
		&models.Shipment{},
		&models.TrackingEvent{},
		&models.LogisticsProvider{},

		// Returns
		&models.Return{},
		&models.ReturnItem{},

		// Admin & System
		&models.Notification{},
		&models.ActivityLog{},
	)
}

// CheckDatabaseHealth checks if database is healthy
func CheckDatabaseHealth() string {
	sqlDB, err := DB.DB()
	if err != nil {
		return "unhealthy"
	}

	if err := sqlDB.Ping(); err != nil {
		return "unhealthy"
	}

	return "healthy"
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}

// getEnv gets environment variable with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
