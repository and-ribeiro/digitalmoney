package services

import (
	"dhmoney/models"
	"dhmoney/repositories"
	"dhmoney/utils"
	"errors"
	"fmt"
)

func RegisterUser(user *models.User) error {

	err := repositories.CreateUser(user)
	if err != nil {
		return fmt.Errorf("Failed to register user: %w", err)
	}

	// Generate the account number and unique alias
	//accountAlias := utils.GenerateAccountAlias()
	accountNumber := utils.GenerateAccountNumber()
	// Create the account
	account := &models.Account{
		UserID:          user.ID,
		AccountNumber:   accountNumber,
		AvailableAmount: 0.0,
	}

	err = repositories.CreateAccount(account)
	if err != nil {
		return fmt.Errorf("Failed to create account: %w", err)
	}
	return nil
}

func LoginUser(email, password string) (*models.User, error) {
	// Retrieve the user from the database by email
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Check if the password is correct
	if !utils.VerifyPassword(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func GetUser(userID int) (*models.User, error) {
	// Call the appropriate repository function to fetch the user from the database
	user, err := repositories.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(userID int, fullName string, email string, phone string) error {
	// Check if the user exists in the database
	user, err := repositories.GetUser(userID)
	if err != nil {
		return fmt.Errorf("Failed to update user: %w", err)
	}

	// Update the user's fields
	user.FullName = fullName
	user.Email = email
	user.Phone = phone

	// Call the repository function to update the user in the database
	err = repositories.UpdateUser(user)
	if err != nil {
		return fmt.Errorf("Failed to update user: %w", err)
	}

	return nil
}
