package main

import (
	"library/config"
	"library/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connexion à la base de données
	config.ConnectDatabase()

	// Initialisation du routeur
	r := gin.Default()

	bookRoutes := r.Group("/books")
	{
		bookRoutes.GET("", handlers.GetBooks)
		bookRoutes.GET("/:id", handlers.GetBookByID)
		bookRoutes.POST("", handlers.CreateBook)
		bookRoutes.PUT("/:id", handlers.UpdateBook)
		bookRoutes.DELETE("/:id", handlers.DeleteBook)
	}

	loanRoutes := r.Group("/loans")
	{
		loanRoutes.GET("", handlers.GetLoans)
		loanRoutes.POST("", handlers.CreateLoan)
		loanRoutes.DELETE("/:id", handlers.DeleteLoan)
	}

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("", handlers.GetUsers)
		userRoutes.POST("", handlers.CreateUser)
	}

	r.Run(":8080")
}
