package main

import (
	"fmt"
	"grello-api/api/handler"
	"grello-api/database"
	"grello-api/internal/model"
)

func main() {
	database.ConnectDb()
	db := database.DB
	hash, err := handler.HashPassword("admin")
	if err != nil {
		fmt.Printf("Password hashing error %s", err.Error())
	}
	if err := db.Create(&model.User{Username: "admin", Password: hash}).Error; err != nil {
		fmt.Printf("User creation error %s", err.Error())
	}
}