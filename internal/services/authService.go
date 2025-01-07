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

	var role models.Role

	s := config.DB.Where("name = ?", "super-admin").First(&role)
	if s.Error != nil {
		return nil, s.Error
	}

	user := models.User{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Password:  string(hashPassword),
		RoleID:    role.ID,
	}

	result := config.DB.Preload("Role").Create(&user)
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
	existingUser := config.DB.Preload("Role").Where("email", req.Email).First(&user)

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

	token, err := helpers.CreateJWT(user.ID, user.Role.Name)
	if err != nil {
		log.Fatal(err)
	}
	return &user, token, nil

}

func AcceptInvite(req *requests.AcceptInviteRequest, id *string) (*models.User, error) {
	var invite models.Invite
	config.DB.Find(&invite, id).First(&invite)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := models.User{
		Firstname: invite.Firstname,
		Lastname:  invite.Lastname,
		Email:     invite.Email,
		Password:  string(hashPassword),
		RoleID:    invite.RoleID,
		Status:    "verified",
	}

	config.DB.Delete(&invite)

	result := config.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	// go mails.SendWelcomeMail(
	// 	user.Email,
	// 	"Welcome to GOPM",
	// 	invite.Firstname+" "+invite.Lastname)

	return &user, nil
}

func VerifyAccount(id *string) (string, error) {
	var user models.User
	existingUser := config.DB.Find(&user, id).First(&user)

	if existingUser.RowsAffected == 0 {
		return "", errors.New("account doesn't exist")
	}

	config.DB.Model(&user).Update("status", "verified")

	return "Account email verified. Sign in to continue", nil
}
