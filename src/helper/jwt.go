package helper

import (
	"time"

	"github.com/afiefafian/go-pos/src/config"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWToken(userID string) (string, error) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Hour)
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SECRET))
}
