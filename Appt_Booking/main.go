package main

import (
	"appt_booking/db"
	"appt_booking/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.POST("/signup", handlers.SignUpHandler) // or however you're importing it
	r.POST("/login", handlers.Login)
	//Connects to the Database
	err := db.Connect()
	if err != nil {
		fmt.Println("Failed to connect", err)
	} else {
		fmt.Println("Database Connected")
	}
	r.Run(":8080")
	

}