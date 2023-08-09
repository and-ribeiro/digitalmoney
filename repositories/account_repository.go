package repositories

import (
	"dhmoney/database"
	"dhmoney/models"
	"fmt"
	"gorm.io/gorm"
)

// CreateAccount creates a new account in the database.
func CreateAccount(account *models.Account) error {
	err := database.DB.Create(account).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAccount(accountID int) (*models.Account, error) {
	account := &models.Account{}
	if err := database.DB.Joins("User").First(account, accountID).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func UpdateAccountAmount(accountID int, amount int) error {
	account := &models.Account{}
	err := database.DB.Model(account).Where("id = ?", accountID).Update("available_amount", gorm.Expr("available_amount + ?", amount)).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateTransaction(transaction *models.Transaction) error {
	err := database.DB.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllTransactionsByAccountID(accountID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := database.DB.Where("account_id = ?", accountID).Find(&transactions).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}
	return transactions, nil
}
