package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int64) (string, error) {

	// jwt -> heaer algm and typ()jwt, payload(user info adn expiry time), signature (secret key)

	secret := os.Getenv("JWT_SECRET")

	if secret == ""{
		return "", errors.New("JWT is not configured")
	}

	claims := Claims {
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "go_task_api",
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))

	if err!= nil {
		return "", fmt.Errorf("Sign jwt : %w", err)
	}

	return signedToken, nil

}