package handlers

import (
	"dhmoney/models"
	"dhmoney/repositories"
	"dhmoney/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAccount(c *gin.Context) {
	accountID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	account, err := services.GetAccount(accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func Deposit(c *gin.Context) {
	accountID, err := strconv.Atoi(c.Param("account_id"))

	// Parse and validate the request body
	var deposit models.Transaction
	if err := c.ShouldBindJSON(&deposit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the account based on the accountID
	account, err := repositories.GetAccount(accountID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	deposit.Destination = account.AccountNumber
	deposit.AccountID = accountID

	// Perform the deposit operation
	err = services.UpdateAccountAmount(accountID, deposit.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Call the service to register the transaction
	transaction := services.CreateTransaction(&deposit)
	if transaction != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": transaction.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deposit successful"})
}

func RetrieveTransactions(c *gin.Context) {
	accountID, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	transactions, err := services.GetAllTransactions(accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}
