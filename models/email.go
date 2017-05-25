package models

type Email struct {
    ID      int
    OwnerID  int     `gorm:"index"` // Foreign key (belongs to), tag `index` will create index for this column
    OwnerType string
    Email   string  `gorm:"type:varchar(100);unique_index"` // `type` set sql type, `unique_index` will create unique index for this column
    Primary bool
    Type string
}
