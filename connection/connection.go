package connection

import (
	"fmt"
	"log"
	"os"
	"uts-golang-laundry/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatalln("Eror to Load Env")
	}

	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error connect To Database")
	}

	db.AutoMigrate(&models.Outlet{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.LaundryType{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.LaundryPrices{})
	db.AutoMigrate(&models.Transaction{})

	return db
}
