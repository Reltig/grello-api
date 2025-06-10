package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"

	"grello-api/config"
	"grello-api/internal/model"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
		config.Config("DB_HOST"),
		config.Config("DB_USERNAME"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
		config.Config("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}
	//db.Logger = logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&model.User{}, &model.Workspace{}, &model.Board{})
	DB = db
}