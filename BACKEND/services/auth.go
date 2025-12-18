package services

import (
	"amass-test/models"
	"amass-test/repositories"
	"amass-test/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (map[string]interface{}, error)
	Register(user models.User) error
}

type authService struct {
	authRepo repositories.AuthRepository
}

func NewAuthService(
	authRepo repositories.AuthRepository,
) AuthService {
	return &authService{
		authRepo: authRepo,
	}
}

func (s *authService) Login(username, password string) (map[string]interface{}, error) {
	user, err := s.authRepo.GetByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	token, err := utils.GenerateToken(user.UserID, user.Username)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"token": token}, nil
}

func (s *authService) Register(user models.User) error {
	user.UserID = utils.GenerateUUID()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	exists, err := s.authRepo.CheckUsernameExists(user.Username)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("username already exists")
	}

	_, err = s.authRepo.Register(user)
	return err
}
