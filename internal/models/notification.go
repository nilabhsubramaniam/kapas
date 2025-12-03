package models

import (
	"time"

	"gorm.io/gorm"
)

type NotificationType string
type NotificationChannel string

const (
	NotificationTypeOrder     NotificationType = "order"
	NotificationTypePayment   NotificationType = "payment"
	NotificationTypeShipping  NotificationType = "shipping"
	NotificationTypePromotion NotificationType = "promotion"
	NotificationTypeSystem    NotificationType = "system"
)

const (
	NotificationChannelEmail NotificationChannel = "email"
	NotificationChannelSMS   NotificationChannel = "sms"
	NotificationChannelInApp NotificationChannel = "in_app"
	NotificationChannelPush  NotificationChannel = "push"
)

// Notification represents user notifications
type Notification struct {
	ID        uint                `gorm:"primaryKey" json:"id"`
	UserID    uint                `gorm:"not null;index" json:"user_id"`
	Type      NotificationType    `gorm:"type:varchar(30);not null" json:"type"`
	Channel   NotificationChannel `gorm:"type:varchar(20);not null" json:"channel"`
	Title     string              `gorm:"not null" json:"title"`
	Message   string              `gorm:"type:text;not null" json:"message"`
	Data      JSONB               `gorm:"type:jsonb" json:"data,omitempty"` // Additional context
	IsRead    bool                `gorm:"default:false;index" json:"is_read"`
	ReadAt    *time.Time          `json:"read_at,omitempty"`
	SentAt    *time.Time          `json:"sent_at,omitempty"`
	CreatedAt time.Time           `json:"created_at"`
	DeletedAt gorm.DeletedAt      `gorm:"index" json:"-"`

	// Relationships
	User User `gorm:"foreignKey:UserID" json:"-"`
}

// ActivityLog represents audit trail for admin actions
type ActivityLog struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null;index" json:"user_id"`
	Action      string    `gorm:"not null;index" json:"action"` // create, update, delete, etc.
	EntityType  string    `gorm:"not null" json:"entity_type"`  // product, order, user, etc.
	EntityID    uint      `json:"entity_id,omitempty"`
	Description string    `gorm:"type:text" json:"description"`
	IPAddress   string    `json:"ip_address,omitempty"`
	UserAgent   string    `gorm:"type:text" json:"user_agent,omitempty"`
	Metadata    JSONB     `gorm:"type:jsonb" json:"metadata,omitempty"`
	CreatedAt   time.Time `json:"created_at"`

	// Relationships
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (Notification) TableName() string {
	return "notifications"
}

func (ActivityLog) TableName() string {
	return "activity_logs"
}
