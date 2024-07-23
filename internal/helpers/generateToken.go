package helpers

import (
	"math/rand"

	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
)

func GenerateToken() int {
	token := rand.Intn(999999-100000+1) + 100000
	if existing := config.DB.Find(models.Token{}, "token = ?", token); existing.RowsAffected > 0 {
		GenerateToken()
	}
	return token
}
