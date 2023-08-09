package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func VerifyPassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func GenerateToken(claims jwt.MapClaims, secretKey string) (string, error) {

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return signedToken, nil
}

const accountNumberLength = 22

func GenerateAccountNumber() string {
	rand.Seed(time.Now().UnixNano())

	var accountNumber string
	for len(accountNumber) < accountNumberLength {
		accountNumber += strconv.Itoa(rand.Intn(10))
	}

	return accountNumber
}
