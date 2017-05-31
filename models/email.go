package models

import "github.com/jinzhu/gorm"

//Email Address
type Email struct {
	gorm.Model
	OwnerID   int `gorm:"index"` // Foreign key (belongs to), tag `index` will create index for this column
	OwnerType string
	Email     string `gorm:"type:varchar(100)"` // `type` set sql type, `unique_index` will create unique index for this column
	Primary   bool
	Type      string
}
