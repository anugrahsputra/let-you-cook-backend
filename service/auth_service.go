package service

import (
	"errors"
	"let-you-cook/domain/model"
	"let-you-cook/repository"
	"let-you-cook/utils/jwt"
	"let-you-cook/utils/security"

	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("main")

type IAuthService interface {
	Register(user model.User) error
	Login(username, password string) (string, error)
}

type authService struct {
	repo repository.IAuthRepo
}

func NewAuthService(repo repository.IAuthRepo) *authService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Register(user model.User) error {
	exist, err := s.repo.GetUserExisting(user.Username)
	if err != nil {
		return err
	}

	if exist.Username != "" {
		return errors.New("user already exist")
	}

	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	err = s.repo.RegisterRepo(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(username, password string) (string, error) {

	user, err := s.repo.CheckUserExistingForLogin(username)
	if err != nil {
		return "", err
	}

	if user.Username == "" || security.CheckPassword(user.Password, password) != nil {
		return "", errors.New("invalid credentials")
	}

	return jwt.GenerateToken(user)

}
