package models

import "time"

type Transaction struct {
	ID          uint      `gorm:"primarykey;autoIncrement"`
	Origin      string    `gorm:"not null"`
	Destination string    `gorm:"default:null;size:22"`
	Amount      int       `gorm:"not null"`
	Date        time.Time `gorm:"autoCreateTime"`
	Info        string    `gorm:"default:null"`
	AccountID   int       `gorm:"not null"`
	Account     Account   `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
}
