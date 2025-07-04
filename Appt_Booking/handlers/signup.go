package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"appt_booking/models"
	"appt_booking/services"
	
)

//? why gin.context
func SignUpHandler(c *gin.Context){
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid input"})
		return
	}

	err := services.CreatUser(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "User created succesfully"})
}