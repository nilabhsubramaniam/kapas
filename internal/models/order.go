package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus string
type PaymentStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusConfirmed  OrderStatus = "confirmed"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusReturned   OrderStatus = "returned"
)

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
)

// Order represents a customer order
type Order struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	OrderNumber     string         `gorm:"uniqueIndex;not null" json:"order_number"`
	UserID          uint           `gorm:"not null;index" json:"user_id"`
	Status          OrderStatus    `gorm:"type:varchar(20);default:'pending';index" json:"status"`
	PaymentStatus   PaymentStatus  `gorm:"type:varchar(20);default:'pending'" json:"payment_status"`
	PaymentMethod   string         `gorm:"type:varchar(50)" json:"payment_method"`
	SubtotalAmount  float64        `gorm:"not null" json:"subtotal_amount"`
	DiscountAmount  float64        `gorm:"default:0" json:"discount_amount"`
	TaxAmount       float64        `gorm:"default:0" json:"tax_amount"`
	ShippingAmount  float64        `gorm:"default:0" json:"shipping_amount"`
	TotalAmount     float64        `gorm:"not null" json:"total_amount"`
	CouponCode      string         `json:"coupon_code,omitempty"`
	ShippingAddress JSONB          `gorm:"type:jsonb" json:"shipping_address"`
	BillingAddress  JSONB          `gorm:"type:jsonb" json:"billing_address,omitempty"`
	CustomerNotes   string         `gorm:"type:text" json:"customer_notes,omitempty"`
	AdminNotes      string         `gorm:"type:text" json:"admin_notes,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	User          User                 `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Items         []OrderItem          `gorm:"foreignKey:OrderID" json:"items,omitempty"`
	StatusHistory []OrderStatusHistory `gorm:"foreignKey:OrderID" json:"status_history,omitempty"`
	Payment       *Payment             `gorm:"foreignKey:OrderID" json:"payment,omitempty"`
	Shipment      *Shipment            `gorm:"foreignKey:OrderID" json:"shipment,omitempty"`
}

// OrderItem represents line items in an order
type OrderItem struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	OrderID     uint           `gorm:"not null;index" json:"order_id"`
	ProductID   uint           `gorm:"not null;index" json:"product_id"`
	ProductName string         `gorm:"not null" json:"product_name"`
	Quantity    int            `gorm:"not null" json:"quantity"`
	UnitPrice   float64        `gorm:"not null" json:"unit_price"`
	TotalPrice  float64        `gorm:"not null" json:"total_price"`
	Metadata    JSONB          `gorm:"type:jsonb" json:"metadata,omitempty"` // Store product snapshot
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Order   Order   `gorm:"foreignKey:OrderID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

// OrderStatusHistory tracks order status changes
type OrderStatusHistory struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	OrderID   uint        `gorm:"not null;index" json:"order_id"`
	Status    OrderStatus `gorm:"type:varchar(20);not null" json:"status"`
	Comment   string      `gorm:"type:text" json:"comment,omitempty"`
	ChangedBy uint        `json:"changed_by,omitempty"` // User ID who changed the status
	CreatedAt time.Time   `json:"created_at"`

	// Relationships
	Order Order `gorm:"foreignKey:OrderID" json:"-"`
}

func (Order) TableName() string {
	return "orders"
}

func (OrderItem) TableName() string {
	return "order_items"
}

func (OrderStatusHistory) TableName() string {
	return "order_status_history"
}
