package models

import (
  "time"
  "github.com/jinzhu/gorm"
)

type Person struct {
    gorm.Model
    Birthday          time.Time
    FirstName         string  `gorm:"size:255"` // Default size for string is 255, reset it with this tag
    LastName          string  `gorm:"size:255"` // Default size for string is 255, reset it with this tag
    Emails            []Email `gorm:"polymorphic:owner" `       // One-To-Many relationship (has many - use Email's UserID as foreign key)
    Addresses         []Address `gorm:"polymorphic:owner" `
    ProfilePic        string
}
