package models

import (
	"time"

	"gorm.io/gorm"
)

// VendorStatus represents the vendor verification status
type VendorStatus string

const (
	VendorStatusPending  VendorStatus = "PENDING"
	VendorStatusVerified VendorStatus = "VERIFIED"
	VendorStatusRejected VendorStatus = "REJECTED"
	VendorStatusSuspended VendorStatus = "SUSPENDED"
)

// Vendor represents a vendor/seller in the system
type Vendor struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	UserID           *uint          `json:"user_id" gorm:"uniqueIndex"` // Links to users table if vendor has login
	BusinessName     string         `json:"business_name" gorm:"size:200;not null"`
	OwnerName        string         `json:"owner_name" gorm:"size:100;not null"`
	Email            string         `json:"email" gorm:"size:100;not null;uniqueIndex"`
	Phone            string         `json:"phone" gorm:"size:20;not null"`
	AlternatePhone   string         `json:"alternate_phone" gorm:"size:20"`
	
	// Business Details
	GSTNumber        string         `json:"gst_number" gorm:"size:15;uniqueIndex"`
	PANNumber        string         `json:"pan_number" gorm:"size:10"`
	BusinessType     string         `json:"business_type" gorm:"size:50"` // Manufacturer, Wholesaler, Artisan, etc.
	YearEstablished  int            `json:"year_established"`
	
	// Address (for audit/compliance)
	AddressLine1     string         `json:"address_line1" gorm:"size:255;not null"`
	AddressLine2     string         `json:"address_line2" gorm:"size:255"`
	Locality         string         `json:"locality" gorm:"size:100"`     // Parsa Chowk, etc.
	DistrictID       uint           `json:"district_id" gorm:"not null;index"`
	StateID          uint           `json:"state_id" gorm:"not null;index"`
	CountryID        uint           `json:"country_id" gorm:"not null;index"`
	Pincode          string         `json:"pincode" gorm:"size:10;not null"`
	
	// Bank Details (for payments)
	BankName         string         `json:"bank_name" gorm:"size:100"`
	BankAccountNo    string         `json:"bank_account_no" gorm:"size:50"`
	BankIFSC         string         `json:"bank_ifsc" gorm:"size:15"`
	BankBranch       string         `json:"bank_branch" gorm:"size:100"`
	
	// Status & Verification
	Status           VendorStatus   `json:"status" gorm:"type:varchar(20);default:'PENDING';index"`
	IsVerified       bool           `json:"is_verified" gorm:"default:false"`
	VerifiedAt       *time.Time     `json:"verified_at"`
	VerifiedBy       *uint          `json:"verified_by"` // Admin user ID
	RejectionReason  string         `json:"rejection_reason" gorm:"type:text"`
	
	// Permissions & Settings
	Permissions      JSONB          `json:"permissions" gorm:"type:jsonb"` // What they can do
	Commission       float64        `json:"commission" gorm:"type:decimal(5,2);default:10.00"` // Platform commission %
	
	// Additional Info
	Description      string         `json:"description" gorm:"type:text"`
	Logo             string         `json:"logo"`
	BannerImage      string         `json:"banner_image"`
	Website          string         `json:"website"`
	
	// Ratings
	Rating           float64        `json:"rating" gorm:"type:decimal(3,2);default:0"`
	TotalReviews     int            `json:"total_reviews" gorm:"default:0"`
	
	// Metrics
	TotalProducts    int            `json:"total_products" gorm:"default:0"`
	TotalOrders      int            `json:"total_orders" gorm:"default:0"`
	TotalRevenue     int64          `json:"total_revenue" gorm:"default:0"` // In paise
	
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relationships
	User             *User          `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Country          Country        `json:"country,omitempty" gorm:"foreignKey:CountryID"`
	State            State          `json:"state,omitempty" gorm:"foreignKey:StateID"`
	District         District       `json:"district,omitempty" gorm:"foreignKey:DistrictID"`
	Products         []Product      `json:"products,omitempty" gorm:"foreignKey:VendorID"`
}

func (Vendor) TableName() string {
	return "vendors"
}
