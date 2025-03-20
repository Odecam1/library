package handlers

import (
	"library/config"
	"library/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET all users
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Preload("Loans").Find(&users)
	c.JSON(http.StatusOK, users)
}

// POST create a new user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}
