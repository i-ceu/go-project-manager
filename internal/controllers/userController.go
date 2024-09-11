package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/helpers"
	"github.com/ubaniIsaac/go-project-manager/internal/mails"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
)

func RegisterUser(c *gin.Context) {
	var req struct {
		FirstName       string `validate:"required"`
		LastName        string `validate:"required"`
		Email           string `validate:"required,email"`
		Password        string `validate:"required"`
		ConfirmPassword string `validate:"required,eqfield=Password"`
		Role            string
	}

	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}

	var existingUser models.User
	email := config.DB.Find(&existingUser, "email = ?", req.Email)
	if email.RowsAffected > 0 {
		c.JSON(403, gin.H{
			"message": "Account exists with this email",
		})
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
		Role:      req.Role,
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result,
		})
		return
	}

	go mails.SendWelcomeMail(
		user.Email,
		"Welcome to GOPM",
		req.FirstName+" "+req.LastName)

	c.JSON(201, gin.H{
		"message": "User registered succefully",
		"User":    user,
	})
}

func SignIn(c *gin.Context) {
	var req struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required"`
	}
	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}
	var user models.User
	existingUser := config.DB.Where("email", req.Email).First(&user)

	if existingUser.RowsAffected == 0 {
		c.JSON(403, gin.H{
			"message": "Account doesn't exist",
		})
		return
	}
	if user.Status != "verified" {
		c.JSON(403, gin.H{
			"message": "Please verify your account to login",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Invalid credentials",
		})
		return
	}
	token, err := helpers.CreateJWT(user.ID, user.Role)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"message": "Signed in Successfuly",
		"user":    user,
		"token":   token,
	})

}
