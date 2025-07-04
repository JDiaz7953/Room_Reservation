package handlers

import (
	"appt_booking/db"
	"appt_booking/models"
	"net/http"

	"github.com/gin-gonic/gin"
)



func Login(c *gin.Context){

	//holds the users login information
	var input struct {
		Username string `json: "username"`
		Password string `json: "password"`
	}

	//Store users information in struct
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid input"})
		return
	}

	//check if the user exist in the database
	var user models.User
	result := db.DB.Where("username =?", input.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "Username not found"})
		return
	} 

	//Check Password
	if input.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "Incorrect password"})
		return
	}

	//create token
	token, err := CreateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Login successful",
		"token" : token,
	})
}