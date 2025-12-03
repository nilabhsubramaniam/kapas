package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/nilabhsubramaniam/kapas/docs" // Import generated docs
	"github.com/nilabhsubramaniam/kapas/internal/config"
	"github.com/nilabhsubramaniam/kapas/internal/handlers"
	"github.com/nilabhsubramaniam/kapas/internal/middleware"
)

// @title Tantuka E-Commerce API
// @version 1.0
// @description Backend API for Tantuka - Premium Indian Saree E-Commerce Platform
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@tantuka.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	// Initialize database connection
	config.InitDatabase()

	// Get environment
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	// Set Gin mode
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" || appEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	router := gin.Default()

	// Apply middleware
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.LoggerMiddleware())

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Tantuka E-Commerce API",
			"version": "1.0.0",
			"status":  "running",
		})
	})

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":   "healthy",
			"database": config.CheckDatabaseHealth(),
		})
	})

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/docs", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})

	// API routes
	api := router.Group("/api")
	{
		// Health check
		api.GET("/health", handlers.HealthCheck)

		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
			auth.GET("/me", middleware.AuthMiddleware(), handlers.GetCurrentUser)
			auth.POST("/logout", middleware.AuthMiddleware(), handlers.Logout)
		}

		// Product routes
		products := api.Group("/products")
		{
			products.GET("", handlers.ListProducts)
			products.GET("/:slug", handlers.GetProduct)
			products.GET("/state/:state", handlers.GetProductsByState)

			// Protected routes (admin only)
			protected := products.Group("")
			protected.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
			{
				protected.POST("", handlers.CreateProduct)
				protected.PUT("/:id", handlers.UpdateProduct)
				protected.DELETE("/:id", handlers.DeleteProduct)
			}
		}

		// Cart routes (protected)
		cart := api.Group("/cart")
		cart.Use(middleware.AuthMiddleware())
		{
			cart.GET("", handlers.GetCart)
			cart.POST("/items", handlers.AddToCart)
			cart.PUT("/items/:id", handlers.UpdateCartItem)
			cart.DELETE("/items/:id", handlers.RemoveFromCart)
			cart.DELETE("", handlers.ClearCart)
		}

		// Order routes (protected)
		orders := api.Group("/orders")
		orders.Use(middleware.AuthMiddleware())
		{
			orders.POST("", handlers.CreateOrder)
			orders.GET("", handlers.ListUserOrders)
			orders.GET("/:id", handlers.GetOrder)
			orders.PUT("/:id/cancel", handlers.CancelOrder)
			orders.GET("/:id/track", handlers.TrackOrder)
		}

		// Admin routes
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
		{
			// Dashboard & Analytics
			admin.GET("/dashboard", handlers.GetDashboard)
			admin.GET("/analytics/sales", handlers.GetSalesAnalytics)
			admin.GET("/analytics/revenue", handlers.GetRevenueAnalytics)

			// User Management
			admin.GET("/users", handlers.ListAllUsers)
			admin.GET("/users/:id", handlers.GetUserDetails)
			admin.GET("/users/:id/orders", handlers.GetUserOrders)
			admin.PUT("/users/:id/status", handlers.UpdateUserStatus)

			// Order Management
			admin.GET("/orders", handlers.ListAllOrders)
			admin.PUT("/orders/:id/status", handlers.UpdateOrderStatus)

			// Inventory Management
			admin.GET("/inventory", handlers.GetInventory)
			admin.PUT("/inventory/:id", handlers.UpdateInventory)
		}
	}

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("🚀 Tantuka Backend starting on port %s (Environment: %s)", port, appEnv)
	log.Printf("📚 API Documentation: http://localhost:%s/swagger/index.html", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
