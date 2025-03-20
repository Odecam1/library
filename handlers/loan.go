package handlers

import (
	"library/config"
	"library/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET all loans
func GetLoans(c *gin.Context) {
	var loans []models.Loan
	config.DB.Find(&loans)
	c.JSON(http.StatusOK, loans)
}

// POST create a new loan
func CreateLoan(c *gin.Context) {
	var loan models.Loan
	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&loan)
	c.JSON(http.StatusCreated, loan)
}

// DELETE a loan
func DeleteLoan(c *gin.Context) {
	id := c.Param("id")
	var loan models.Loan
	if err := config.DB.First(&loan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}
	config.DB.Delete(&loan)
	c.JSON(http.StatusOK, gin.H{"message": "Loan deleted"})
}
