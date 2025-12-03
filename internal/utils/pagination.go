package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// PaginationParams represents pagination parameters
type PaginationParams struct {
	Page    int
	PerPage int
	Offset  int
}

// GetPaginationParams extracts pagination from query parameters
func GetPaginationParams(c *gin.Context) PaginationParams {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))

	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 20
	}

	offset := (page - 1) * perPage

	return PaginationParams{
		Page:    page,
		PerPage: perPage,
		Offset:  offset,
	}
}

// PaginatedResponse creates a paginated JSON response
func PaginatedResponse(data interface{}, total int64, page, perPage int) gin.H {
	totalPages := int((total + int64(perPage) - 1) / int64(perPage))

	return gin.H{
		"data": data,
		"pagination": gin.H{
			"page":        page,
			"per_page":    perPage,
			"total":       total,
			"total_pages": totalPages,
		},
	}
}
