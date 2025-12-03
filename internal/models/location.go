package models

import (
	"time"

	"gorm.io/gorm"
)

// Country represents a country in the system
type Country struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"size:100;not null;unique"`
	Code      string         `json:"code" gorm:"size:3;not null;unique"` // ISO 3166-1 alpha-2 (IN, US, etc.)
	PhoneCode string         `json:"phone_code" gorm:"size:10"`           // +91, +1, etc.
	Currency  string         `json:"currency" gorm:"size:3"`              // INR, USD, etc.
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relationships
	States []State `json:"states,omitempty" gorm:"foreignKey:CountryID"`
}

// State represents a state/province in a country
type State struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CountryID uint           `json:"country_id" gorm:"not null;index"`
	Name      string         `json:"name" gorm:"size:100;not null"`
	Code      string         `json:"code" gorm:"size:10;not null"` // UP, KL, TN, etc.
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relationships
	Country   Country    `json:"country,omitempty" gorm:"foreignKey:CountryID"`
	Districts []District `json:"districts,omitempty" gorm:"foreignKey:StateID"`
	Regions   []Region   `json:"regions,omitempty" gorm:"foreignKey:StateID"`
}

// District represents a district/city in a state
type District struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	StateID   uint           `json:"state_id" gorm:"not null;index"`
	Name      string         `json:"name" gorm:"size:100;not null"`
	Code      string         `json:"code" gorm:"size:20"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relationships
	State State `json:"state,omitempty" gorm:"foreignKey:StateID"`
}

// Region represents a famous craft/product region for marketing purposes
// This is what products link to for origin display
type Region struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"size:100;not null"` // Kerala, Lucknow, Madhubani, etc.
	Slug        string         `json:"slug" gorm:"size:100;not null;uniqueIndex"`
	Type        string         `json:"type" gorm:"size:50"`                        // State, City, Region
	StateID     *uint          `json:"state_id" gorm:"index"`                      // Optional - may span multiple states
	Description string         `json:"description" gorm:"type:text"`               // Famous for Kasavu sarees, Chikankari work, etc.
	FamousFor   string         `json:"famous_for" gorm:"type:text"`                // Comma-separated: Kasavu Sarees, Silk Sarees, etc.
	ImageURL    string         `json:"image_url"`                                  // Region banner image
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	DisplayOrder int           `json:"display_order" gorm:"default:0"`             // For sorting on UI
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relationships
	State    *State    `json:"state,omitempty" gorm:"foreignKey:StateID"`
	Products []Product `json:"products,omitempty" gorm:"foreignKey:RegionID"`
}

// TableName overrides
func (Country) TableName() string {
	return "countries"
}

func (State) TableName() string {
	return "states"
}

func (District) TableName() string {
	return "districts"
}

func (Region) TableName() string {
	return "regions"
}
