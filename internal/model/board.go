package model

import "gorm.io/gorm"

type Board struct {
	gorm.Model
	Name 		 string  `gorm:"not null"`
	Description  string  `gorm:"not null"`
	Public       bool
	WorkspaceID  uint
	Users 		 []*User `gorm:"many2many:user_boards;"`
}