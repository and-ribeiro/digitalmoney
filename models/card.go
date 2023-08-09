package models

type Card struct {
	ID             int     `gorm:"primarykey;autoIncrement"`
	FullName       string  `gorm:"not null"`
	Number         int     `gorm:"not null"`
	SecurityCode   int     `gorm:"not null"`
	ExpirationDate string  `gorm:"not null"`
	AccountID      int     `gorm:"not null"`
	Account        Account `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
}
