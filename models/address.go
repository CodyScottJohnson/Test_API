package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	OwnerID   int `gorm:"index"` // Foreign key (belongs to), tag `index` will create index for this column
	OwnerType string
	Address   string `gorm:"not null;"` // Set field as not nullable and unique
	City      string
	State     string
	Post      sql.NullString //`gorm:"not null"`
	Primary   bool
	Type      string
}
