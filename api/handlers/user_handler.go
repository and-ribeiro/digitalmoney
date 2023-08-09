package handlers

import (
	"dhmoney/models"
	"dhmoney/services"
	"dhmoney/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func LoginUser(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind the JSON request body to the loginRequest struct
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate the login credentials
	user, err := services.LoginUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login credentials"})
		return
	}

	// Create the claims for the JWT token
	claims := jwt.MapClaims{
		"userID":    user.ID,
		"userEmail": user.Email,
		"userCPF":   user.CPF,
	}

	// Set secret key
	secretKey := "secret-dhmoney-key"

	// Generate the token
	token, err := utils.GenerateToken(claims, secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the token in the response
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func LogoutUser(c *gin.Context) {
	// Get the JWT token from the request header
	token := c.GetHeader("Authorization")
	// Validate and invalidate the token
	// Perform any necessary logout actions
	_ = token
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful / not implemented yet"})
}

func GetUser(c *gin.Context) {
	userID := c.Param("id")
	// Convert the userID to an integer value
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := services.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	// Extract the user ID from the request URL parameters
	userID := c.Param("id")

	// Convert the userID to an integer value
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Parse the request body to get the updated fields for the user
	var updateRequest struct {
		FullName string `json:"FullName"`
		Email    string `json:"Email"`
		Phone    string `json:"Phone"`
	}

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call the corresponding service function to update the user
	err = services.UpdateUser(id, updateRequest.FullName, updateRequest.Email, updateRequest.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
