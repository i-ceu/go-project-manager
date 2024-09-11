package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(userID string, role string) (string, error) {
	secret := []byte(os.Getenv("jwtSecret"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role":   role,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24 * 10).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, err
}
