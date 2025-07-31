package services

import (
	"errors"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/i-ceu/go-project-manager/internal/config"
	"github.com/i-ceu/go-project-manager/internal/helpers"
	"github.com/i-ceu/go-project-manager/internal/mails"
	"github.com/i-ceu/go-project-manager/internal/models"
	"github.com/i-ceu/go-project-manager/internal/requests"
)

func CreateTeam(req *requests.CreateTeamRequest, user *models.User) (*models.Team, error) {
	var existingTeam models.Team
	v := config.DB.Find(&existingTeam, "name = ? AND creator_id = ?", req.Name, user.ID)
	if v.RowsAffected > 0 {
		return nil, errors.New("you already have a team with this name")
	}

	team := models.Team{
		Name:        req.Name,
		Description: req.Description,
		Size:        req.Size,
		Industry:    req.Industry,
		CreatorID:   user.ID,
	}

	org := config.DB.Create(&team)
	if org.Error != nil {
		return nil, org.Error
	}

	var role models.Role
	s := config.DB.Where("name = ?", "super-admin").First(&role)
	if s.Error != nil {
		return nil, s.Error
	}

	member_role := models.MemberRole{
		UserID: user.ID,
		RoleID: role.ID,
		TeamID: team.ID,
	}
	sr := config.DB.Create(&member_role)
	if sr.Error != nil {
		return nil, sr.Error
	}

	return &team, nil
}

func SignInToTeam(userId string, teamId string) (*models.MemberRole, string, error) {
	var memberRole models.MemberRole
	result := config.DB.Preload("Role").Preload("User").Preload("Team").Where("user_id = ? AND team_id = ?", userId, teamId).First(&memberRole)
	if result.Error != nil {
		return nil, "", errors.New("no team for user with this id")
	}

	token, err := helpers.CreateJWT(userId, teamId)
	if err != nil {
		log.Fatal(err)
	}
	return &memberRole, token, nil
}

func InviteToTeam(req *requests.SendInviteRequest, teamId *string, user *models.User) error {
	email := config.DB.Find(models.User{}, "email = ?", req.Email)
	// user
	if email.RowsAffected > 0 {
		return errors.New("account already exists with this email")

	}

	var org models.Team
	result := config.DB.Find(&org, teamId)
	if result.Error != nil {
		return errors.New("no team with this id")
	}
	invite := models.Invite{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		RoleID:    req.RoleID,
		TeamID:    *teamId,
		SentByID:  user.ID,
	}
	sr := config.DB.Create(&invite)
	if sr.Error != nil {
		return sr.Error
	}

	app_url := os.Getenv("APP_URL")

	go mails.SendInviteMail(
		req.Email,
		"Invite to GOPM",
		org.Name,
		app_url+"/auth/acceptInvite/"+invite.ID,
	)
	return nil

}

func AcceptInvite(req *requests.AcceptInviteRequest, inviteId *string) (*models.MemberRole, error) {
	var invite models.Invite
	err := config.DB.Where("id = ?", inviteId).First(&invite).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invite doesn't exist")
		}
		// Handle other database errors
		return nil, err
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := models.User{
		Firstname: invite.Firstname,
		Lastname:  invite.Lastname,
		Email:     invite.Email,
		Password:  string(hashPassword),
		Status:    "verified",
	}

	result := config.DB.Where(models.User{Email: user.Email}).FirstOrCreate(&user)

	// user

	// user := models.User{}

	// if userexists.RowsAffected > 0 {

	// 	result := config.DB.Create(&user)

	// }

	// user = models.User{
	// 	Firstname: invite.Firstname,
	// 	Lastname:  invite.Lastname,
	// 	Email:     invite.Email,
	// 	Password:  string(hashPassword),
	// 	Status:    "verified",
	// }

	member_role := models.MemberRole{
		UserID: user.ID,
		RoleID: invite.RoleID,
		TeamID: invite.TeamID,
	}
	sr := config.DB.Create(&member_role)
	if sr.Error != nil {
		return nil, sr.Error
	}

	err = config.DB.Preload("User").Preload("Team").Preload("Role").Where("id = ?", member_role.ID).First(&member_role).Error
	if err != nil {
		return nil, err
	}

	config.DB.Delete(&invite)

	if result.Error != nil {
		return nil, result.Error
	}

	return &member_role, nil
}

func GetUserTeams(userId string) (*[]models.MemberRole, error) {
	var allTeams []models.MemberRole
	result := config.DB.Select("team_id", "role_id").Preload("Team").Preload("Role").Where("user_id = ?", userId).Find(&allTeams)
	if result.Error != nil {
		return nil, errors.New("error retrieving teams")
	}

	return &allTeams, nil
}

func GetTeamMembers(teamId string) (*[]models.MemberRole, error) {
	var allMembers []models.MemberRole
	result := config.DB.Select("user_id", "role_id").Preload("User").Preload("Role").Where("team_id = ?", teamId).Find(&allMembers)
	if result.Error != nil {
		return nil, errors.New("error retrieving teams")
	}

	return &allMembers, nil
}

func GetRoles() (*[]models.Role, error) {
	var roles []models.Role
	result := config.DB.Find(&roles)
	if result.Error != nil {
		return nil, errors.New("error retrieving roles")
	}

	return &roles, nil
}
