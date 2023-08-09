package main

import (
	"dhmoney/api/routes"
	"dhmoney/database"
	"github.com/gin-gonic/gin"
)

func main() {
	//Create Gin Router
	router := gin.Default()

	//Register the routes
	routes.RegisterRoutes(router)

	//Attempt to connect to database
	database.Connect()

	//Start the server
	router.Run(":8080")
}
