package models

import (
	"time"

	"gorm.io/gorm"
)

// Warehouse represents a warehouse location
type Warehouse struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Code      string         `gorm:"uniqueIndex;not null" json:"code"`
	Address   string         `gorm:"type:text;not null" json:"address"`
	City      string         `gorm:"not null" json:"city"`
	State     string         `gorm:"not null" json:"state"`
	PinCode   string         `gorm:"not null" json:"pin_code"`
	Phone     string         `json:"phone,omitempty"`
	Email     string         `json:"email,omitempty"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Inventory []Inventory `gorm:"foreignKey:WarehouseID" json:"inventory,omitempty"`
}

// Inventory represents stock levels per warehouse
type Inventory struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	ProductID         uint           `gorm:"not null;index" json:"product_id"`
	WarehouseID       uint           `gorm:"not null;index" json:"warehouse_id"`
	Quantity          int            `gorm:"not null;default:0" json:"quantity"`
	ReservedQuantity  int            `gorm:"default:0" json:"reserved_quantity"` // Items in pending orders
	LowStockThreshold int            `gorm:"default:10" json:"low_stock_threshold"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Product   Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Warehouse Warehouse `gorm:"foreignKey:WarehouseID" json:"warehouse,omitempty"`
}

func (Warehouse) TableName() string {
	return "warehouses"
}

func (Inventory) TableName() string {
	return "inventory"
}
