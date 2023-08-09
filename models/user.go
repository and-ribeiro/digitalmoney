package models

import "time"

type User struct {
	ID        int       `gorm:"primarykey;autoIncrement"`
	FullName  string    `gorm:"not null"`
	CPF       string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Phone     string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Alias     string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
