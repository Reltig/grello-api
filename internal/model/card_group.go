package model

import "gorm.io/gorm"

type CardGroup struct {
	gorm.Model
	Name 	string `gorm:"not null"`
	BoardID	uint
}