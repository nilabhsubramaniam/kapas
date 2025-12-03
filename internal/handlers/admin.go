package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nilabhsubramaniam/kapas/internal/config"
	"github.com/nilabhsubramaniam/kapas/internal/models"
	"github.com/nilabhsubramaniam/kapas/internal/utils"
)

// ============================================
// DASHBOARD & ANALYTICS
// ============================================

// GetDashboard godoc
// @Summary Get admin dashboard statistics
// @Description Get overview statistics for admin dashboard
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Dashboard statistics"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Router /admin/dashboard [get]
func GetDashboard(c *gin.Context) {
	var stats struct {
		TotalUsers       int64   `json:"total_users"`
		TotalOrders      int64   `json:"total_orders"`
		TotalProducts    int64   `json:"total_products"`
		TotalRevenue     float64 `json:"total_revenue"`
		PendingOrders    int64   `json:"pending_orders"`
		LowStockProducts int64   `json:"low_stock_products"`
	}

	// Get counts
	config.DB.Model(&models.User{}).Count(&stats.TotalUsers)
	config.DB.Model(&models.Order{}).Count(&stats.TotalOrders)
	config.DB.Model(&models.Product{}).Count(&stats.TotalProducts)
	config.DB.Model(&models.Order{}).Where("status = ?", "PENDING").Count(&stats.PendingOrders)
	config.DB.Model(&models.Product{}).Where("stock_quantity < ?", 10).Count(&stats.LowStockProducts)

	// Calculate total revenue
	var revenue struct {
		Total float64
	}
	config.DB.Model(&models.Order{}).
		Where("status IN ?", []string{"COMPLETED", "DELIVERED"}).
		Select("COALESCE(SUM(total_amount), 0) as total").
		Scan(&revenue)
	stats.TotalRevenue = revenue.Total

	c.JSON(http.StatusOK, stats)
}

// GetSalesAnalytics godoc
// @Summary Get sales analytics
// @Description Get sales data for charts and reports
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param period query string false "Period (day, week, month, year)" default(month)
// @Success 200 {object} map[string]interface{} "Sales analytics"
// @Router /admin/analytics/sales [get]
func GetSalesAnalytics(c *gin.Context) {
	period := c.DefaultQuery("period", "month")
	
	var groupBy string
	var startDate time.Time
	
	switch period {
	case "day":
		groupBy = "DATE(created_at)"
		startDate = time.Now().AddDate(0, 0, -30) // Last 30 days
	case "week":
		groupBy = "DATE_TRUNC('week', created_at)"
		startDate = time.Now().AddDate(0, 0, -90) // Last 90 days
	case "year":
		groupBy = "DATE_TRUNC('month', created_at)"
		startDate = time.Now().AddDate(-2, 0, 0) // Last 2 years
	default: // month
		groupBy = "DATE(created_at)"
		startDate = time.Now().AddDate(0, -12, 0) // Last 12 months
	}

	var salesData []struct {
		Date   string  `json:"date"`
		Orders int64   `json:"orders"`
		Revenue float64 `json:"revenue"`
	}

	config.DB.Model(&models.Order{}).
		Select(groupBy + " as date, COUNT(*) as orders, COALESCE(SUM(total_amount), 0) as revenue").
		Where("created_at >= ?", startDate).
		Group("date").
		Order("date ASC").
		Scan(&salesData)

	c.JSON(http.StatusOK, gin.H{
		"period": period,
		"data":   salesData,
	})
}

// GetRevenueAnalytics godoc
// @Summary Get revenue analytics
// @Description Get revenue breakdown by product type, state, etc.
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Revenue analytics"
// @Router /admin/analytics/revenue [get]
func GetRevenueAnalytics(c *gin.Context) {
	// Revenue by product type
	var byProductType []struct {
		ProductType string  `json:"product_type"`
		Revenue     float64 `json:"revenue"`
		Orders      int64   `json:"orders"`
	}

	config.DB.Table("order_items").
		Select("products.product_type, COUNT(DISTINCT order_items.order_id) as orders, COALESCE(SUM(order_items.price * order_items.quantity), 0) as revenue").
		Joins("JOIN products ON products.id = order_items.product_id").
		Group("products.product_type").
		Scan(&byProductType)

	// Revenue by state origin
	var byState []struct {
		State   string  `json:"state"`
		Revenue float64 `json:"revenue"`
		Orders  int64   `json:"orders"`
	}

	config.DB.Table("order_items").
		Select("products.state_origin as state, COUNT(DISTINCT order_items.order_id) as orders, COALESCE(SUM(order_items.price * order_items.quantity), 0) as revenue").
		Joins("JOIN products ON products.id = order_items.product_id").
		Where("products.state_origin IS NOT NULL AND products.state_origin != ''").
		Group("products.state_origin").
		Scan(&byState)

	c.JSON(http.StatusOK, gin.H{
		"by_product_type": byProductType,
		"by_state":        byState,
	})
}

// ============================================
// USER MANAGEMENT
// ============================================

// ListAllUsers godoc
// @Summary List all users
// @Description Get paginated list of all users (Admin only)
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(20)
// @Param role query string false "Filter by role (customer, admin, vendor)"
// @Param search query string false "Search by name or email"
// @Success 200 {object} map[string]interface{} "Paginated users list"
// @Router /admin/users [get]
func ListAllUsers(c *gin.Context) {
	pagination := utils.GetPaginationParams(c)
	
	var users []models.User
	var total int64

	query := config.DB.Model(&models.User{})

	// Filter by role
	if role := c.Query("role"); role != "" {
		query = query.Where("role = ?", role)
	}

	// Search by name or email
	if search := c.Query("search"); search != "" {
		query = query.Where("name ILIKE ? OR email ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	query.Count(&total)

	query.Order("created_at DESC").
		Limit(pagination.PerPage).
		Offset((pagination.Page - 1) * pagination.PerPage).
		Find(&users)

	c.JSON(http.StatusOK, utils.PaginatedResponse(users, total, pagination.Page, pagination.PerPage))
}

// GetUserDetails godoc
// @Summary Get user details
// @Description Get detailed information about a specific user (Admin only)
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} models.User "User details"
// @Failure 404 {object} ErrorResponse "User not found"
// @Router /admin/users/{id} [get]
func GetUserDetails(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := config.DB.Preload("Addresses").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserOrders godoc
// @Summary Get user's orders
// @Description Get all orders for a specific user (Admin only)
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(20)
// @Success 200 {object} map[string]interface{} "User's orders"
// @Router /admin/users/{id}/orders [get]
func GetUserOrders(c *gin.Context) {
	userID := c.Param("id")
	pagination := utils.GetPaginationParams(c)

	var orders []models.Order
	var total int64

	config.DB.Model(&models.Order{}).Where("user_id = ?", userID).Count(&total)

	config.DB.Where("user_id = ?", userID).
		Preload("OrderItems").
		Preload("Payment").
		Order("created_at DESC").
		Limit(pagination.PerPage).
		Offset((pagination.Page - 1) * pagination.PerPage).
		Find(&orders)

	c.JSON(http.StatusOK, utils.PaginatedResponse(orders, total, pagination.Page, pagination.PerPage))
}

// UpdateUserStatus godoc
// @Summary Update user status
// @Description Activate, deactivate, or verify a user (Admin only)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param request body map[string]interface{} true "Status update"
// @Success 200 {object} MessageResponse "User updated successfully"
// @Router /admin/users/{id}/status [put]
func UpdateUserStatus(c *gin.Context) {
	userID := c.Param("id")
	
	var req struct {
		IsActive      *bool `json:"is_active"`
		EmailVerified *bool `json:"email_verified"`
		Role          string `json:"role"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if req.IsActive != nil {
		updates["is_active"] = *req.IsActive
	}
	if req.EmailVerified != nil {
		updates["email_verified"] = *req.EmailVerified
	}
	if req.Role != "" {
		updates["role"] = req.Role
	}

	if err := config.DB.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// ============================================
// ORDER MANAGEMENT
// ============================================

// ListAllOrders godoc
// @Summary List all orders
// @Description Get paginated list of all orders from all users (Admin only)
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(20)
// @Param status query string false "Filter by status (PENDING, CONFIRMED, SHIPPED, etc.)"
// @Param user_id query int false "Filter by user ID"
// @Success 200 {object} map[string]interface{} "Paginated orders list"
// @Router /admin/orders [get]
func ListAllOrders(c *gin.Context) {
	pagination := utils.GetPaginationParams(c)
	
	var orders []models.Order
	var total int64

	query := config.DB.Model(&models.Order{})

	// Filter by status
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// Filter by user
	if userID := c.Query("user_id"); userID != "" {
		query = query.Where("user_id = ?", userID)
	}

	query.Count(&total)

	query.Preload("User").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Preload("Payment").
		Order("created_at DESC").
		Limit(pagination.PerPage).
		Offset((pagination.Page - 1) * pagination.PerPage).
		Find(&orders)

	c.JSON(http.StatusOK, utils.PaginatedResponse(orders, total, pagination.Page, pagination.PerPage))
}

// UpdateOrderStatus godoc
// @Summary Update order status
// @Description Update order status and tracking (Admin only)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Param request body map[string]interface{} true "Status update"
// @Success 200 {object} MessageResponse "Order updated successfully"
// @Router /admin/orders/{id}/status [put]
func UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")
	
	var req struct {
		Status         string `json:"status"`
		TrackingNumber string `json:"tracking_number"`
		Notes          string `json:"notes"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if req.Status != "" {
		updates["status"] = req.Status
	}
	if req.TrackingNumber != "" {
		updates["tracking_number"] = req.TrackingNumber
	}

	if err := config.DB.Model(&models.Order{}).Where("id = ?", orderID).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	// Create activity log
	if req.Notes != "" {
		orderIDInt, _ := strconv.ParseUint(orderID, 10, 32)
		log := models.ActivityLog{
			UserID:      0, // System action
			Action:      "ORDER_STATUS_UPDATED",
			Description: req.Notes,
			EntityType:  "order",
			EntityID:    uint(orderIDInt),
		}
		config.DB.Create(&log)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

// ============================================
// INVENTORY MANAGEMENT
// ============================================

// GetInventory godoc
// @Summary Get inventory list
// @Description Get all products with stock information (Admin only)
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(50)
// @Param low_stock query bool false "Show only low stock items (< 10)"
// @Success 200 {object} map[string]interface{} "Inventory list"
// @Router /admin/inventory [get]
func GetInventory(c *gin.Context) {
	pagination := utils.GetPaginationParams(c)
	
	var products []models.Product
	var total int64

	query := config.DB.Model(&models.Product{})

	// Filter low stock
	if c.Query("low_stock") == "true" {
		query = query.Where("stock_quantity < ?", 10)
	}

	query.Count(&total)

	query.Select("id, name, slug, product_type, state_origin, stock_quantity, base_price, final_price, is_active").
		Order("stock_quantity ASC").
		Limit(pagination.PerPage).
		Offset((pagination.Page - 1) * pagination.PerPage).
		Find(&products)

	c.JSON(http.StatusOK, utils.PaginatedResponse(products, total, pagination.Page, pagination.PerPage))
}

// UpdateInventory godoc
// @Summary Update product inventory
// @Description Update stock quantity for a product (Admin only)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param request body map[string]interface{} true "Inventory update"
// @Success 200 {object} MessageResponse "Inventory updated successfully"
// @Router /admin/inventory/{id} [put]
func UpdateInventory(c *gin.Context) {
	productID := c.Param("id")
	
	var req struct {
		StockQuantity int  `json:"stock_quantity"`
		IsActive      *bool `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{
		"stock_quantity": req.StockQuantity,
	}
	if req.IsActive != nil {
		updates["is_active"] = *req.IsActive
	}

	if err := config.DB.Model(&models.Product{}).Where("id = ?", productID).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory updated successfully"})
}
