package models

import (
  "database/sql"
)

type Address struct {
    ID        int
    OwnerID   int     `gorm:"index"` // Foreign key (belongs to), tag `index` will create index for this column
    OwnerType string
    Address1  string         `gorm:"not null;unique"` // Set field as not nullable and unique
    Address2  string
    City      string
    State     string
    Post      sql.NullString `gorm:"not null"`
    Primary   bool
    Type      string
}
