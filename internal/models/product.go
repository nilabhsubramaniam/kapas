package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type ProductType string
type SareeType string
type FabricType string

const (
	ProductTypeSaree           ProductType = "SAREE"
	ProductTypeChikankariKurti ProductType = "CHIKANKARI_KURTI"
	ProductTypeChikankariDress ProductType = "CHIKANKARI_DRESS"
)

// JSONB type for flexible metadata
type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, j)
}

// Product represents a product in the catalog
type Product struct {
	ID                 uint           `gorm:"primaryKey" json:"id"`
	Name               string         `gorm:"not null;index" json:"name" binding:"required"`
	Slug               string         `gorm:"uniqueIndex;not null" json:"slug"`
	Description        string         `gorm:"type:text" json:"description"`
	ProductType        ProductType    `gorm:"type:varchar(50);not null;index" json:"product_type"`
	
	// Location & Vendor (NEW - using IDs instead of strings)
	RegionID           *uint          `gorm:"index" json:"region_id"`              // Links to regions table for origin display
	VendorID           *uint          `gorm:"index" json:"vendor_id"`              // Which vendor supplies this product
	
	// Deprecated fields (keep for backward compatibility, will migrate data)
	StateOrigin        string         `gorm:"type:varchar(10);index" json:"state_origin,omitempty"` // UP, KL, TN, etc. (deprecated)
	
	SareeType          string         `gorm:"index" json:"saree_type"`                    // Chikankari, Kasavu, Kanchipuram, etc.
	BasePrice          float64        `gorm:"not null" json:"base_price" binding:"required,gt=0"`
	DiscountPercentage float64        `gorm:"default:0" json:"discount_percentage"`
	FinalPrice         float64        `gorm:"not null" json:"final_price"`
	Fabric             string         `json:"fabric"`
	WeaveType          string         `json:"weave_type"`
	Occasion           string         `json:"occasion"`
	StockQuantity      int            `gorm:"default:0" json:"stock_quantity"`
	IsActive           bool           `gorm:"default:true;index" json:"is_active"`
	Metadata           JSONB          `gorm:"type:jsonb" json:"metadata,omitempty"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Region   *Region        `gorm:"foreignKey:RegionID" json:"region,omitempty"`
	Vendor   *Vendor        `gorm:"foreignKey:VendorID" json:"vendor,omitempty"`
	Images   []ProductImage `gorm:"foreignKey:ProductID" json:"images,omitempty"`
	Reviews  []Review       `gorm:"foreignKey:ProductID" json:"reviews,omitempty"`
	Category []Category     `gorm:"many2many:product_categories;" json:"categories,omitempty"`
}

// ProductImage represents product images
type ProductImage struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ProductID    uint      `gorm:"not null;index" json:"product_id"`
	ImageURL     string    `gorm:"not null" json:"image_url"`
	AltText      string    `json:"alt_text,omitempty"`
	DisplayOrder int       `gorm:"default:0" json:"display_order"`
	IsPrimary    bool      `gorm:"default:false" json:"is_primary"`
	CreatedAt    time.Time `json:"created_at"`

	// Relationship
	Product Product `gorm:"foreignKey:ProductID" json:"-"`
}

// Category represents product categorization
type Category struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Name         string         `gorm:"not null;index" json:"name"`
	Slug         string         `gorm:"uniqueIndex;not null" json:"slug"`
	ParentID     *uint          `gorm:"index" json:"parent_id,omitempty"`
	CategoryType string         `gorm:"type:varchar(50)" json:"category_type"` // STATE, FABRIC, OCCASION, etc.
	StateCode    string         `gorm:"type:varchar(10)" json:"state_code,omitempty"`
	Description  string         `gorm:"type:text" json:"description,omitempty"`
	DisplayOrder int            `gorm:"default:0" json:"display_order"`
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	Metadata     JSONB          `gorm:"type:jsonb" json:"metadata,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Parent   *Category  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children []Category `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}

// Review represents a product review
type Review struct {
	ID                 uint           `gorm:"primaryKey" json:"id"`
	ProductID          uint           `gorm:"not null;index" json:"product_id"`
	UserID             uint           `gorm:"not null;index" json:"user_id"`
	Rating             int            `gorm:"not null" json:"rating" binding:"required,min=1,max=5"`
	Comment            string         `gorm:"type:text" json:"comment,omitempty"`
	IsVerifiedPurchase bool           `gorm:"default:false" json:"is_verified_purchase"`
	IsApproved         bool           `gorm:"default:true" json:"is_approved"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Product Product `gorm:"foreignKey:ProductID" json:"-"`
	User    User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (Product) TableName() string {
	return "products"
}

func (ProductImage) TableName() string {
	return "product_images"
}

func (Category) TableName() string {
	return "categories"
}

func (Review) TableName() string {
	return "reviews"
}
