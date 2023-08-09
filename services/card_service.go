package services

import (
	"dhmoney/models"
	"dhmoney/repositories"
	"fmt"
)

func CreateCard(accountID int, card *models.Card) error {
	// Check if the account exists
	_, err := repositories.GetAccount(accountID)
	if err != nil {
		return fmt.Errorf("failed to create card: %w", err)
	}

	// Set the account ID for the card
	card.AccountID = accountID

	// Call the repository function to create the card
	err = repositories.CreateCard(card)
	if err != nil {
		return fmt.Errorf("failed to create card: %w", err)
	}

	return nil
}

func GetAllCards(accountID int) ([]models.Card, error) {
	// Check if the account exists
	_, err := repositories.GetAccount(accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cards: %w", err)
	}

	// Call the repository function to get all cards for the account
	cards, err := repositories.GetAllCardsByAccountID(accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cards: %w", err)
	}

	return cards, nil
}

func DeleteCard(accountID, cardID int) error {
	// Call the repository function to delete the card
	err := repositories.DeleteCard(accountID, cardID)
	if err != nil {
		return fmt.Errorf("failed to delete card: %w", err)
	}
	return nil
}

func GetCard(accountID, cardID int) (*models.Card, error) {
	// Call the repository function to get the card
	card, err := repositories.GetCard(accountID, cardID)
	if err != nil {
		return nil, fmt.Errorf("failed to get card: %w", err)
	}

	return card, nil
}
