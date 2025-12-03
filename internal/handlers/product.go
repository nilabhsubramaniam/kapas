package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nilabhsubramaniam/kapas/internal/config"
	"github.com/nilabhsubramaniam/kapas/internal/models"
	"github.com/nilabhsubramaniam/kapas/internal/utils"
)

// CreateProductRequest represents product creation request
type CreateProductRequest struct {
	Name               string              `json:"name" binding:"required"`
	Description        string              `json:"description"`
	ProductType        string              `json:"product_type" binding:"required"`
	StateOrigin        string              `json:"state_origin"`
	SareeType          string              `json:"saree_type"`
	BasePrice          float64             `json:"base_price" binding:"required,gt=0"`
	DiscountPercentage float64             `json:"discount_percentage"`
	Fabric             string              `json:"fabric"`
	WeaveType          string              `json:"weave_type"`
	Occasion           string              `json:"occasion"`
	StockQuantity      int                 `json:"stock_quantity"`
	Images             []ProductImageInput `json:"images"`
	Metadata           map[string]interface{} `json:"metadata"`
}

type ProductImageInput struct {
	ImageURL     string `json:"image_url" binding:"required"`
	AltText      string `json:"alt_text"`
	DisplayOrder int    `json:"display_order"`
	IsPrimary    bool   `json:"is_primary"`
}

// ListProducts godoc
// @Summary List all products
// @Description Get paginated list of products with optional filters
// @Tags Products
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(20)
// @Param state query string false "State origin (UP, KL, TN, KA, WB)"
// @Param saree_type query string false "Saree type (Chikankari, Kasavu, Kanchipuram, etc.)"
// @Param fabric query string false "Fabric type (Cotton, Silk, Georgette, etc.)"
// @Param product_type query string false "Product type (SAREE, CHIKANKARI_KURTI, etc.)"
// @Param occasion query string false "Occasion (Wedding, Festival, Casual, Party)"
// @Param min_price query number false "Minimum price"
// @Param max_price query number false "Maximum price"
// @Param sort query string false "Sort field (created_at, price, name)" default(created_at)
// @Param order query string false "Sort order (asc, desc)" default(desc)
// @Success 200 {object} PaginatedProductsResponse "Paginated products list"
// @Router /products [get]
func ListProducts(c *gin.Context) {
	pagination := utils.GetPaginationParams(c)
	
	var products []models.Product
	var total int64

	query := config.DB.Model(&models.Product{}).Where("is_active = ?", true)

	// Filters
	if state := c.Query("state"); state != "" {
		query = query.Where("state_origin = ?", state)
	}
	if sareeType := c.Query("saree_type"); sareeType != "" {
		query = query.Where("saree_type = ?", sareeType)
	}
	if fabric := c.Query("fabric"); fabric != "" {
		query = query.Where("fabric = ?", fabric)
	}
	if productType := c.Query("product_type"); productType != "" {
		query = query.Where("product_type = ?", productType)
	}
	if occasion := c.Query("occasion"); occasion != "" {
		query = query.Where("occasion = ?", occasion)
	}

	// Price range filter
	if minPrice := c.Query("min_price"); minPrice != "" {
		if price, err := strconv.ParseFloat(minPrice, 64); err == nil {
			query = query.Where("final_price >= ?", price)
		}
	}
	if maxPrice := c.Query("max_price"); maxPrice != "" {
		if price, err := strconv.ParseFloat(maxPrice, 64); err == nil {
			query = query.Where("final_price <= ?", price)
		}
	}

	// Sorting
	sortBy := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "desc")
	if order != "asc" && order != "desc" {
		order = "desc"
	}
	query = query.Order(sortBy + " " + order)

	// Count total
	query.Count(&total)

	// Fetch products with pagination
	query.Preload("Images").
		Limit(pagination.PerPage).
		Offset(pagination.Offset).
		Find(&products)

	c.JSON(http.StatusOK, utils.PaginatedResponse(products, total, pagination.Page, pagination.PerPage))
}

// GetProduct returns a single product by slug
func GetProduct(c *gin.Context) {
	slug := c.Param("slug")

	var product models.Product
	if err := config.DB.Where("slug = ? AND is_active = ?", slug, true).
		Preload("Images").
		Preload("Reviews").
		First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetProductsByState godoc
// @Summary Get products by state
// @Description Get paginated list of products filtered by state origin
// @Tags Products
// @Produce json
// @Param state path string true "State code (UP, KL, TN, KA, WB)"
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(20)
// @Success 200 {object} map[string]interface{} "Paginated products list"
// @Router /products/state/{state} [get]
func GetProductsByState(c *gin.Context) {
	state := c.Param("state")
	pagination := utils.GetPaginationParams(c)

	var products []models.Product
	var total int64

	config.DB.Model(&models.Product{}).
		Where("state_origin = ? AND is_active = ?", state, true).
		Count(&total)

	config.DB.Where("state_origin = ? AND is_active = ?", state, true).
		Preload("Images").
		Limit(pagination.PerPage).
		Offset(pagination.Offset).
		Find(&products)

	c.JSON(http.StatusOK, utils.PaginatedResponse(products, total, pagination.Page, pagination.PerPage))
}

// CreateProduct creates a new product (admin only)
func CreateProduct(c *gin.Context) {
	var req CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate slug from name
	slug := generateSlug(req.Name)

	// Check if slug already exists
	var existingProduct models.Product
	if err := config.DB.Where("slug = ?", slug).First(&existingProduct).Error; err == nil {
		// Slug exists, append random number
		slug = slug + "-" + strconv.FormatInt(int64(existingProduct.ID), 10)
	}

	// Calculate final price
	finalPrice := req.BasePrice
	if req.DiscountPercentage > 0 {
		finalPrice = req.BasePrice - (req.BasePrice * req.DiscountPercentage / 100)
	}

	// Create product
	product := models.Product{
		Name:               req.Name,
		Slug:               slug,
		Description:        req.Description,
		ProductType:        models.ProductType(req.ProductType),
		StateOrigin:        req.StateOrigin,
		SareeType:          req.SareeType,
		BasePrice:          req.BasePrice,
		DiscountPercentage: req.DiscountPercentage,
		FinalPrice:         finalPrice,
		Fabric:             req.Fabric,
		WeaveType:          req.WeaveType,
		Occasion:           req.Occasion,
		StockQuantity:      req.StockQuantity,
		IsActive:           true,
		Metadata:           models.JSONB(req.Metadata),
	}

	// Start transaction
	tx := config.DB.Begin()

	// Create product
	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	// Create product images
	if len(req.Images) > 0 {
		for _, img := range req.Images {
			productImage := models.ProductImage{
				ProductID:    product.ID,
				ImageURL:     img.ImageURL,
				AltText:      img.AltText,
				DisplayOrder: img.DisplayOrder,
				IsPrimary:    img.IsPrimary,
			}
			if err := tx.Create(&productImage).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product images"})
				return
			}
		}
	}

	// Commit transaction
	tx.Commit()

	// Reload product with images
	config.DB.Preload("Images").First(&product, product.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}

// UpdateProduct updates a product (admin only)
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var req CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Calculate final price
	finalPrice := req.BasePrice
	if req.DiscountPercentage > 0 {
		finalPrice = req.BasePrice - (req.BasePrice * req.DiscountPercentage / 100)
	}

	// Update product fields
	product.Name = req.Name
	product.Description = req.Description
	product.ProductType = models.ProductType(req.ProductType)
	product.StateOrigin = req.StateOrigin
	product.SareeType = req.SareeType
	product.BasePrice = req.BasePrice
	product.DiscountPercentage = req.DiscountPercentage
	product.FinalPrice = finalPrice
	product.Fabric = req.Fabric
	product.WeaveType = req.WeaveType
	product.Occasion = req.Occasion
	product.StockQuantity = req.StockQuantity
	product.Metadata = models.JSONB(req.Metadata)

	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	// Reload with images
	config.DB.Preload("Images").First(&product, product.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"product": product,
	})
}

// DeleteProduct deletes a product (admin only)
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Soft delete
	if err := config.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}

// generateSlug creates URL-friendly slug from name
func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")
	// Remove special characters
	allowedChars := "abcdefghijklmnopqrstuvwxyz0123456789-"
	result := ""
	for _, char := range slug {
		if strings.ContainsRune(allowedChars, char) {
			result += string(char)
		}
	}
	return result
}