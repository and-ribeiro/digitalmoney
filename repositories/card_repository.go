package repositories

import (
	"dhmoney/database"
	"dhmoney/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func CreateCard(card *models.Card) error {
	err := database.DB.Create(card).Error
	if err != nil {
		return fmt.Errorf("failed to create card: %w", err)
	}
	return nil
}

func GetAllCardsByAccountID(accountID int) ([]models.Card, error) {
	var cards []models.Card
	err := database.DB.Where("account_id = ?", accountID).Find(&cards).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get cards: %w", err)
	}
	return cards, nil
}

func DeleteCard(accountID, cardID int) error {

	// Delete the card from the database
	result := database.DB.Where("id = ? AND account_id = ?", cardID, accountID).Delete(&models.Card{})
	if result.Error != nil {
		return result.Error
	}

	// Check if any rows were affected
	if result.RowsAffected == 0 {
		return errors.New("card not found or not associated with the account")
	}

	return nil
}

func GetCard(accountID, cardID int) (*models.Card, error) {
	// Retrieve the card from the database based on the account ID and card ID
	var card models.Card
	result := database.DB.First(&card, "account_id = ? AND id = ?", accountID, cardID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("card not found")
		}
		return nil, result.Error
	}

	return &card, nil
}
