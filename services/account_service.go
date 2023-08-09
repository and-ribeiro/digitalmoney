package services

import (
	"dhmoney/models"
	"dhmoney/repositories"
	"fmt"
)

func GetAccount(accountID int) (*models.Account, error) {
	account, err := repositories.GetAccount(accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}
	return account, nil
}

func UpdateAccountAmount(accountID int, amount int) error {
	err := repositories.UpdateAccountAmount(accountID, amount)
	if err != nil {
		return err
	}
	return nil
}

func CreateTransaction(transaction *models.Transaction) error {
	err := repositories.CreateTransaction(transaction)
	if err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}
	return nil
}

func GetAllTransactions(accountID int) ([]models.Transaction, error) {
	_, err := repositories.GetAccount(accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cards: %w", err)
	}

	transactions, err := repositories.GetAllTransactionsByAccountID(accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cards: %w", err)
	}

	return transactions, nil

}
