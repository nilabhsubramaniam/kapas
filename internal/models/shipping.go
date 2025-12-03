package models

import (
	"time"

	"gorm.io/gorm"
)

type ShipmentStatus string

const (
	ShipmentStatusPending        ShipmentStatus = "pending"
	ShipmentStatusPickedUp       ShipmentStatus = "picked_up"
	ShipmentStatusInTransit      ShipmentStatus = "in_transit"
	ShipmentStatusOutForDelivery ShipmentStatus = "out_for_delivery"
	ShipmentStatusDelivered      ShipmentStatus = "delivered"
	ShipmentStatusFailed         ShipmentStatus = "failed"
	ShipmentStatusReturned       ShipmentStatus = "returned"
)

// LogisticsProvider represents shipping/courier companies
type LogisticsProvider struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null;uniqueIndex" json:"name"`
	Code        string         `gorm:"uniqueIndex;not null" json:"code"` // delhivery, bluedart, shiprocket
	APIEndpoint string         `json:"api_endpoint,omitempty"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Shipments []Shipment `gorm:"foreignKey:ProviderID" json:"shipments,omitempty"`
}

// Shipment represents a shipment/delivery
type Shipment struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	OrderID           uint           `gorm:"not null;uniqueIndex" json:"order_id"`
	ProviderID        uint           `gorm:"not null;index" json:"provider_id"`
	AWBNumber         string         `gorm:"uniqueIndex" json:"awb_number"` // Airway Bill Number
	Status            ShipmentStatus `gorm:"type:varchar(30);default:'pending';index" json:"status"`
	Weight            float64        `json:"weight,omitempty"`                       // in kg
	Dimensions        JSONB          `gorm:"type:jsonb" json:"dimensions,omitempty"` // length, width, height
	ShippingCost      float64        `gorm:"default:0" json:"shipping_cost"`
	EstimatedDelivery *time.Time     `json:"estimated_delivery,omitempty"`
	ActualDelivery    *time.Time     `json:"actual_delivery,omitempty"`
	PickupDate        *time.Time     `json:"pickup_date,omitempty"`
	TrackingURL       string         `json:"tracking_url,omitempty"`
	Metadata          JSONB          `gorm:"type:jsonb" json:"metadata,omitempty"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Order          Order             `gorm:"foreignKey:OrderID" json:"-"`
	Provider       LogisticsProvider `gorm:"foreignKey:ProviderID" json:"provider,omitempty"`
	TrackingEvents []TrackingEvent   `gorm:"foreignKey:ShipmentID" json:"tracking_events,omitempty"`
}

// TrackingEvent represents shipment tracking history
type TrackingEvent struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ShipmentID  uint      `gorm:"not null;index" json:"shipment_id"`
	Status      string    `gorm:"not null" json:"status"`
	Location    string    `json:"location,omitempty"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	EventTime   time.Time `gorm:"not null" json:"event_time"`
	CreatedAt   time.Time `json:"created_at"`

	// Relationships
	Shipment Shipment `gorm:"foreignKey:ShipmentID" json:"-"`
}

func (LogisticsProvider) TableName() string {
	return "logistics_providers"
}

func (Shipment) TableName() string {
	return "shipments"
}

func (TrackingEvent) TableName() string {
	return "tracking_events"
}
