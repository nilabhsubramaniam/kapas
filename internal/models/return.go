package models

import (
	"time"

	"gorm.io/gorm"
)

type ReturnStatus string
type ReturnReason string

const (
	ReturnStatusRequested ReturnStatus = "requested"
	ReturnStatusApproved  ReturnStatus = "approved"
	ReturnStatusRejected  ReturnStatus = "rejected"
	ReturnStatusPickedUp  ReturnStatus = "picked_up"
	ReturnStatusReceived  ReturnStatus = "received"
	ReturnStatusRefunded  ReturnStatus = "refunded"
)

const (
	ReturnReasonDefective      ReturnReason = "defective"
	ReturnReasonWrongItem      ReturnReason = "wrong_item"
	ReturnReasonNotAsDescribed ReturnReason = "not_as_described"
	ReturnReasonSizeIssue      ReturnReason = "size_issue"
	ReturnReasonOther          ReturnReason = "other"
)

// Return represents a return request
type Return struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	OrderID       uint           `gorm:"not null;index" json:"order_id"`
	UserID        uint           `gorm:"not null;index" json:"user_id"`
	ReturnNumber  string         `gorm:"uniqueIndex;not null" json:"return_number"`
	Reason        ReturnReason   `gorm:"type:varchar(50);not null" json:"reason"`
	ReasonDetails string         `gorm:"type:text" json:"reason_details,omitempty"`
	Status        ReturnStatus   `gorm:"type:varchar(20);default:'requested';index" json:"status"`
	RefundAmount  float64        `json:"refund_amount,omitempty"`
	AdminNotes    string         `gorm:"type:text" json:"admin_notes,omitempty"`
	ApprovedBy    *uint          `json:"approved_by,omitempty"`
	ApprovedAt    *time.Time     `json:"approved_at,omitempty"`
	RefundedAt    *time.Time     `json:"refunded_at,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Order Order        `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	User  User         `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Items []ReturnItem `gorm:"foreignKey:ReturnID" json:"items,omitempty"`
}

// ReturnItem represents items being returned
type ReturnItem struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ReturnID    uint      `gorm:"not null;index" json:"return_id"`
	OrderItemID uint      `gorm:"not null;index" json:"order_item_id"`
	ProductID   uint      `gorm:"not null" json:"product_id"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	Reason      string    `gorm:"type:text" json:"reason,omitempty"`
	CreatedAt   time.Time `json:"created_at"`

	// Relationships
	Return    Return    `gorm:"foreignKey:ReturnID" json:"-"`
	OrderItem OrderItem `gorm:"foreignKey:OrderItemID" json:"order_item,omitempty"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

func (Return) TableName() string {
	return "returns"
}

func (ReturnItem) TableName() string {
	return "return_items"
}
