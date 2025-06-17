package model

import "gorm.io/gorm"

type Workspace struct {
	gorm.Model
	Name 		string
	Description string
	OwnerID		uint
	Users		[]*User `gorm:"many2many:user_workspaces;"`
	Boards		[]Board
}