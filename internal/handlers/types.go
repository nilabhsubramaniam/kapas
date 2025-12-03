package handlers

import "time"

// UserResponse represents user data in API responses
type UserResponse struct {
	ID            uint      `json:"id" example:"1"`
	Email         string    `json:"email" example:"user@example.com"`
	Name          string    `json:"name" example:"John Doe"`
	Phone         string    `json:"phone" example:"+919876543210"`
	Role          string    `json:"role" example:"customer"`
	EmailVerified bool      `json:"email_verified" example:"true"`
	IsActive      bool      `json:"is_active" example:"true"`
	CreatedAt     time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
}

// ProductResponse represents product data in API responses
type ProductResponse struct {
	ID                 uint                   `json:"id" example:"1"`
	Name               string                 `json:"name" example:"Lucknow White Chikankari Cotton Saree"`
	Slug               string                 `json:"slug" example:"lucknow-white-chikankari-cotton-saree"`
	Description        string                 `json:"description" example:"Beautiful handcrafted Chikankari saree"`
	ProductType        string                 `json:"product_type" example:"SAREE"`
	StateOrigin        string                 `json:"state_origin" example:"UP"`
	SareeType          string                 `json:"saree_type" example:"Chikankari"`
	BasePrice          float64                `json:"base_price" example:"4999.00"`
	DiscountPercentage float64                `json:"discount_percentage" example:"20.00"`
	FinalPrice         float64                `json:"final_price" example:"3999.00"`
	Fabric             string                 `json:"fabric" example:"Cotton"`
	WeaveType          string                 `json:"weave_type" example:"Hand-embroidered"`
	Occasion           string                 `json:"occasion" example:"Casual,Festival"`
	StockQuantity      int                    `json:"stock_quantity" example:"50"`
	IsActive           bool                   `json:"is_active" example:"true"`
	Metadata           map[string]interface{} `json:"metadata"`
	CreatedAt          time.Time              `json:"created_at" example:"2024-01-01T00:00:00Z"`
}

// PaginatedProductsResponse represents paginated products response
type PaginatedProductsResponse struct {
	Data       []ProductResponse `json:"data"`
	Pagination PaginationMeta    `json:"pagination"`
}

// PaginationMeta represents pagination metadata
type PaginationMeta struct {
	Page       int   `json:"page" example:"1"`
	PerPage    int   `json:"per_page" example:"20"`
	Total      int64 `json:"total" example:"100"`
	TotalPages int   `json:"total_pages" example:"5"`
}

// LoginResponse represents login response
type LoginResponse struct {
	Message string       `json:"message" example:"Login successful"`
	User    UserResponse `json:"user"`
	Token   string       `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// RegisterResponse represents registration response
type RegisterResponse struct {
	Message string       `json:"message" example:"User registered successfully"`
	User    UserResponse `json:"user"`
	Token   string       `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// ErrorResponse represents error response
type ErrorResponse struct {
	Error string `json:"error" example:"Something went wrong"`
}

// MessageResponse represents simple message response
type MessageResponse struct {
	Message string `json:"message" example:"Operation successful"`
}
