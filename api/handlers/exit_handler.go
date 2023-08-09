package handlers

import (
	"dhmoney/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExitApplication(c *gin.Context) {
	fmt.Println("Shutdown signal received!")
	database.DeleteTables()
	c.Status(http.StatusOK)

	c.JSON(http.StatusOK, gin.H{
		"message": "Cleanup done",
	})
}
