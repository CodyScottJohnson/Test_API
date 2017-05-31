package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"size:255"` // Default size for string is 255, reset it with this tag
}
