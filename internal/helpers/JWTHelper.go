package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	UserID string `json:"userID"`
	TeamID string `json:"teamID"`
	jwt.RegisteredClaims
}

func CreateJWT(userID string, teamID string) (string, error) {
	secret := []byte(os.Getenv("jwtSecret"))
	claims := MyClaims{
		UserID: userID,
		TeamID: teamID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, err
}
