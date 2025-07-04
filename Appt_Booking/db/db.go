package db

import (
	"appt_booking/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "Gez10"
	dbname = "Rooms"
)

var DB *gorm.DB

func Connect () error {

	// Gets the information to connect to the database
	psqInfo := fmt.Sprintf("host=%s port=%d user=%s " +
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname,
	)

	//Connects to the database
	db, err := gorm.Open(postgres.Open(psqInfo), &gorm.Config{})
	if err != nil {
		 return fmt.Errorf("failed to open DB: %w")
	}

	DB = db

	fmt.Println("Database connection successful!")
    return DB.AutoMigrate(&models.User{})
}