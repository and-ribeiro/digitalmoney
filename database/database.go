package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"dhmoney/models"
)

var (
	DB *gorm.DB
)

func Connect() {
	dsn := "root:591213709@tcp(host.docker.internal:3306)/dhmoney?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	db.AutoMigrate(&models.User{}, &models.Account{}, &models.Card{}, &models.Transaction{})

	DB = db

	fmt.Println("Database connected")
}

func DeleteTables() {
	fmt.Println("Receiving the call from the shutdown signal. Deleting tables...")
	DB.Migrator().DropTable(&models.User{}, &models.Account{}, &models.Card{}, &models.Transaction{})
	fmt.Println("Tables deleted successfully")
}
