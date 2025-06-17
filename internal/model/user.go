package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string 	     `gorm:"uniqueIndex;not null"`
	Email           string 	     `gorm:"uniqueIndex;not null"`
	Password        string 	     `gorm:"not null"`
	FirstName       *string
	SecondName      *string
	OwnedWorkspaces []Workspace  `gorm:"foreignKey:OwnerID"`
	Workspaces 		[]*Workspace `gorm:"many2many:user_workspaces;"`
	Boards			[]*Board     `gorm:"many2many:user_boards;"`
}
