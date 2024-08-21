package models

import "gorm.io/gorm"

type Tasks struct {
	gorm.Model

	Tittle      string `gorm:"type:varchar(200);not null;unique_index" json:"title"`
	Description string `json:"description"`
	Done        bool   `gorm:"default:false" json:"done"`
	UserId      uint   `json:"user_id"`
}
