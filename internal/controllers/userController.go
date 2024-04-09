package controllers

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
)

func Register(c *gin.Context) {
	var req struct {
		FirstName       string
		LastName        string
		Email           string
		Password        string
		ConfirmPassword string
		Role            string
	}

	c.Bind(&req)
	// fmt.Println(c.PostForm("password"))
	// fmt.Println(c.PostForm("confirmPassword"))
	if req.Password != req.ConfirmPassword {
		c.JSON(422, gin.H{
			"message": "Password and Confrim password dont match",
		})
		return
	}
	var existingUser models.User
	email := config.DB.Where("email = ?", req.Email).First(&existingUser)
	// fmt.Println(email)
	if email.RowsAffected > 0 {
		c.JSON(403, gin.H{
			"message": "Account exists with this email",
		})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return
	}
	user := models.User{
		Firstname: req.FirstName,
		Lastname:  req.LastName,
		Email:     req.Email,
		Password:  string(hashPassword),
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		// c.Status(400)
		c.JSON(400, gin.H{
			"message": result,
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "User registered succefully",
		"User":    user,
	})
	// return
}

func SignIn(c *gin.Context) {
	var req struct {
		Email    string
		Password string
	}
	c.Bind(&req)
	var user models.User
	existingUser := config.DB.Where("email = ?", req.Email).First(&user)

	if existingUser.RowsAffected == 0 {
		c.JSON(403, gin.H{
			"message": "Account doesn't exist",
		})

	}
	//compare password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	token, err := createJWT(int64(user.ID))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"message": "Signed in Successfuly",
		"user":    user,
		"token":   token,
	})

}

func createJWT(userID int64) (string, error) {
	secret := []byte(os.Getenv("jwtSecret"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(time.Hour * 24 * 120).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, err
}
