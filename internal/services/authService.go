package services

import (
	"errors"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/i-ceu/go-project-manager/internal/config"
	"github.com/i-ceu/go-project-manager/internal/helpers"
	"github.com/i-ceu/go-project-manager/internal/mails"
	"github.com/i-ceu/go-project-manager/internal/models"
	"github.com/i-ceu/go-project-manager/internal/requests"
)

func RegisterUser(req *requests.RegisterUserRequest) (*models.User, error) {
	var existingUser models.User
	email := config.DB.Find(&existingUser, "email = ?", req.Email)
	if email.RowsAffected > 0 {
		return nil, errors.New("user already exists with this email")
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := models.User{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Password:  string(hashPassword),
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	app_url := os.Getenv("APP_URL")

	verification_link := app_url + "/auth/verify/" + user.ID

	go mails.SendWelcomeMail(
		user.Email,
		"Welcome to GOPM",
		req.Firstname+" "+req.Lastname,
		verification_link)

	return &user, nil
}

func SignIn(req *requests.SignInRequest) (*models.User, string, error) {
	var user models.User
	existingUser := config.DB.Preload("MemberRoles").Preload("MemberRoles.Role").Preload("MemberRoles.Team").Where("email", req.Email).First(&user)

	if existingUser.RowsAffected == 0 {
		return nil, "", errors.New("account doesn't exist")
	}

	if user.Status != "verified" {
		return nil, "", errors.New("please verify your account to login")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, "", errors.New("invalid credentials")

	}

	token, err := helpers.CreateJWT(user.ID, "nil")
	if err != nil {
		log.Fatal(err)
	}
	return &user, token, nil
}

func VerifyAccount(verificationId *string) (string, error) {
	var user models.User
	existingUser := config.DB.Find(&user, verificationId).First(&user)

	if existingUser.RowsAffected == 0 {
		return "", errors.New("account doesn't exist")
	}

	config.DB.Model(&user).Update("status", "verified")

	return "Account email verified. Sign in to continue", nil
}
