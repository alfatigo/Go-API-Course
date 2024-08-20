package models

import "gorm.io/gorm"

type Tasks struct {
	gorm.Model

	Tittle      string `gorm:"type:varchar(200);not null;unique_index"`
	Description string
	Done        bool `gorm:"default:false"`
	UserId      uint
}
