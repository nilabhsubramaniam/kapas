package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	RoleCustomer UserRole = "customer"
	RoleAdmin    UserRole = "admin"
	RoleVendor   UserRole = "vendor"
)

// User represents a user account (customer, admin, vendor)
type User struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Email         string         `gorm:"uniqueIndex;not null" json:"email" binding:"required,email"`
	PasswordHash  string         `gorm:"not null" json:"-"`
	Name          string         `gorm:"not null" json:"name" binding:"required"`
	Phone         string         `gorm:"size:15" json:"phone"`
	Role          UserRole       `gorm:"type:varchar(20);default:'customer'" json:"role"`
	EmailVerified bool           `gorm:"default:false" json:"email_verified"`
	IsActive      bool           `gorm:"default:true" json:"is_active"`
	LastLogin     *time.Time     `json:"last_login,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Addresses []Address      `gorm:"foreignKey:UserID" json:"addresses,omitempty"`
	CartItems []CartItem     `gorm:"foreignKey:UserID" json:"cart_items,omitempty"`
	Orders    []Order        `gorm:"foreignKey:UserID" json:"orders,omitempty"`
	Reviews   []Review       `gorm:"foreignKey:UserID" json:"reviews,omitempty"`
	Wishlists []WishlistItem `gorm:"foreignKey:UserID" json:"wishlists,omitempty"`
}

// Address represents a shipping or billing address
type Address struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	UserID       uint           `gorm:"not null;index" json:"user_id"`
	FullName     string         `gorm:"not null" json:"full_name"`
	Phone        string         `gorm:"not null" json:"phone"`
	AddressLine1 string         `gorm:"not null" json:"address_line1"`
	AddressLine2 string         `json:"address_line2,omitempty"`
	Landmark     string         `json:"landmark,omitempty"`
	
	// Location Hierarchy (using IDs instead of strings)
	DistrictID   uint           `gorm:"not null;index" json:"district_id"`
	StateID      uint           `gorm:"not null;index" json:"state_id"`
	CountryID    uint           `gorm:"not null;index" json:"country_id"`
	PinCode      string         `gorm:"not null" json:"pin_code"`
	
	AddressType  string         `gorm:"type:varchar(20);default:'shipping'" json:"address_type"` // shipping, billing, home, work
	IsDefault    bool           `gorm:"default:false" json:"is_default"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	User     User     `gorm:"foreignKey:UserID" json:"-"`
	Country  Country  `gorm:"foreignKey:CountryID" json:"country,omitempty"`
	State    State    `gorm:"foreignKey:StateID" json:"state,omitempty"`
	District District `gorm:"foreignKey:DistrictID" json:"district,omitempty"`
}

// TableName overrides the table name
func (User) TableName() string {
	return "users"
}

func (Address) TableName() string {
	return "addresses"
}
