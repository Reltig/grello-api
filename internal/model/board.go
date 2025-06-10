package model

import "gorm.io/gorm"

type Board struct {
	gorm.Model
	Name 		string
	Description string
	WorkspaceID uint
}