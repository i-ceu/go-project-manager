package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/i-ceu/go-project-manager/internal/config"
	"github.com/i-ceu/go-project-manager/internal/helpers"
	"github.com/i-ceu/go-project-manager/internal/models"
)

func Auth() gin.HandlerFunc {

	var secret = []byte(os.Getenv("jwtSecret"))

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.ParseWithClaims(tokenString, &helpers.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*helpers.MyClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to extract token claims"})
			c.Abort()
		}

		c.Set("userID", claims.UserID)
		c.Set("teamID", claims.TeamID)

		c.Next()
	}
}

func Guest() gin.HandlerFunc {

	var secret = []byte(os.Getenv("jwtSecret"))
	type MyClaims struct {
		UserID string `json:"userID"`
		jwt.RegisteredClaims
	}

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*MyClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to extract token claims"})
			c.Abort()
		}

		c.Set("userID", claims.UserID)

		c.Next()
	}
}

func CheckRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		userID, exists := c.Get("userID")
		teamID, exists := c.Get("teamID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No claims found"})
			c.Abort()
			return
		}

		// jwtClaims := claims.(*helpers.MyClaims)

		var memberRole models.MemberRole
		err := config.DB.Preload("Role").Where("user_id = ? AND team_id = ?",
			userID, teamID).First(&memberRole).Error

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "User not authorized for this team"})
			c.Abort()
			return
		}

		roleAllowed := false
		for _, role := range allowedRoles {
			if memberRole.Role.Name == "super-admin" {
				roleAllowed = true
				break
			} else if memberRole.Role.Name == role {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Set("userRole", memberRole.Role)
		c.Set("teamID", teamID)
		c.Set("userID", userID)
		c.Next()
	}
}
