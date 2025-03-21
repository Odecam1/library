package config

import (
	"library/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:00700700@tcp(localhost:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.User{}, &models.Book{}, &models.Loan{})

	DB = database
	log.Println("✅ Connexion à la base de données réussie !")
}
