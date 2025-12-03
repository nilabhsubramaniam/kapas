package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod string
type PaymentProvider string

const (
	PaymentMethodCard       PaymentMethod = "card"
	PaymentMethodUPI        PaymentMethod = "upi"
	PaymentMethodNetBanking PaymentMethod = "netbanking"
	PaymentMethodWallet     PaymentMethod = "wallet"
	PaymentMethodCOD        PaymentMethod = "cod"
)

const (
	PaymentProviderRazorpay PaymentProvider = "razorpay"
	PaymentProviderStripe   PaymentProvider = "stripe"
)

// Payment represents a payment transaction
type Payment struct {
	ID                uint            `gorm:"primaryKey" json:"id"`
	OrderID           uint            `gorm:"not null;uniqueIndex" json:"order_id"`
	PaymentProvider   PaymentProvider `gorm:"type:varchar(50);not null" json:"payment_provider"`
	ProviderOrderID   string          `gorm:"uniqueIndex" json:"provider_order_id"`             // Razorpay order_id
	ProviderPaymentID string          `gorm:"uniqueIndex" json:"provider_payment_id,omitempty"` // Razorpay payment_id
	PaymentMethod     PaymentMethod   `gorm:"type:varchar(50)" json:"payment_method,omitempty"`
	Amount            float64         `gorm:"not null" json:"amount"`
	Currency          string          `gorm:"default:'INR'" json:"currency"`
	Status            PaymentStatus   `gorm:"type:varchar(20);default:'pending';index" json:"status"`
	PaymentSignature  string          `json:"payment_signature,omitempty"` // For verification
	ErrorCode         string          `json:"error_code,omitempty"`
	ErrorDescription  string          `gorm:"type:text" json:"error_description,omitempty"`
	Metadata          JSONB           `gorm:"type:jsonb" json:"metadata,omitempty"`
	PaidAt            *time.Time      `json:"paid_at,omitempty"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	DeletedAt         gorm.DeletedAt  `gorm:"index" json:"-"`

	// Relationships
	Order Order `gorm:"foreignKey:OrderID" json:"-"`
}

// Coupon represents a discount coupon
type Coupon struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	Code           string         `gorm:"uniqueIndex;not null" json:"code"`
	Description    string         `gorm:"type:text" json:"description,omitempty"`
	DiscountType   string         `gorm:"type:varchar(20);not null" json:"discount_type"` // percentage, fixed
	DiscountValue  float64        `gorm:"not null" json:"discount_value"`
	MinOrderAmount float64        `gorm:"default:0" json:"min_order_amount"`
	MaxDiscount    float64        `json:"max_discount,omitempty"`       // For percentage discounts
	UsageLimit     int            `gorm:"default:0" json:"usage_limit"` // 0 = unlimited
	UsedCount      int            `gorm:"default:0" json:"used_count"`
	ValidFrom      time.Time      `json:"valid_from"`
	ValidUntil     time.Time      `json:"valid_until"`
	IsActive       bool           `gorm:"default:true" json:"is_active"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Usages []CouponUsage `gorm:"foreignKey:CouponID" json:"usages,omitempty"`
}

// CouponUsage tracks coupon usage by users
type CouponUsage struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	CouponID uint      `gorm:"not null;index" json:"coupon_id"`
	UserID   uint      `gorm:"not null;index" json:"user_id"`
	OrderID  uint      `gorm:"not null;uniqueIndex" json:"order_id"`
	UsedAt   time.Time `json:"used_at"`

	// Relationships
	Coupon Coupon `gorm:"foreignKey:CouponID" json:"-"`
	User   User   `gorm:"foreignKey:UserID" json:"-"`
	Order  Order  `gorm:"foreignKey:OrderID" json:"-"`
}

func (Payment) TableName() string {
	return "payments"
}

func (Coupon) TableName() string {
	return "coupons"
}

func (CouponUsage) TableName() string {
	return "coupon_usages"
}
