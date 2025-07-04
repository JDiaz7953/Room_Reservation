package services

import (
	"appt_booking/db"
	model "appt_booking/models"
	"fmt"
)

func CreatUser(u model.User) error {
	fmt.Printf("Signing up user: %v\n", u)
	result := db.DB.Create(&u)
	return result.Error
}