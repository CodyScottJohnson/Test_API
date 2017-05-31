package models

import "github.com/jinzhu/gorm"

type Recruit struct {
	gorm.Model
	Person      Person
	PersonID    int
	Tags        []Tag `gorm:"many2many:recruit_tags;"`
	Pop         int
	ColorRed    int `json:"Color_Red"`
	ColorBlue   int `json:"Color_Blue"`
	ColorWhite  int `json:"Color_White"`
	ColorYellow int `json:"Color_Yellow"`
	Source      string
	Archived    bool `sql:"Default:false"`
	Old_ID      int
}
