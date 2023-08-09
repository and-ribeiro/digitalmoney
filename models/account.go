package models

type Account struct {
	ID              int    `gorm:"primarykey;autoIncrement"`
	UserID          int    `gorm:"not null"`
	AccountNumber   string `gorm:"unique;not null;size:22"`
	AvailableAmount int    `gorm:"not null;default:0"`
	User            User   `gorm:"constraint:OnDelete:CASCADE;"`
}
