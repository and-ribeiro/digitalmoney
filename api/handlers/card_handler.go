package handlers

import (
	"dhmoney/models"
	"dhmoney/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateCard(c *gin.Context) {
	// Get the account ID from the request parameters
	accountID, err := strconv.Atoi(c.Param("account_id"))

	// Parse the card data from the request body
	var card models.Card
	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card data"})
		return
	}

	// Call the corresponding service function to create the card
	err = services.CreateCard(accountID, &card)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"card": card})
}

func GetAllCards(c *gin.Context) {
	// Get the account ID from the request parameters
	accountID, err := strconv.Atoi(c.Param("account_id"))

	// Call the corresponding service function to get all cards for the account
	cards, err := services.GetAllCards(accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cards": cards})
}

func DeleteCard(c *gin.Context) {
	// Get the account ID from the URL parameter
	accountID, err := strconv.Atoi(c.Param("account_id"))

	// Get the card ID from the URL parameter
	cardID, err := strconv.Atoi(c.Param("card_id"))

	// Call the corresponding service function to delete the card
	err = services.DeleteCard(accountID, cardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Card deleted successfully"})
}

func GetCard(c *gin.Context) {
	accountID := c.Param("account_id")
	cardID := c.Param("card_id")

	// Convert the accountID and cardID to integers
	accountIDInt, err := strconv.Atoi(accountID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	cardIDInt, err := strconv.Atoi(cardID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ID"})
		return
	}

	// Call the corresponding service function to get the card
	card, err := services.GetCard(accountIDInt, cardIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the card in the response
	c.JSON(http.StatusOK, card)
}
