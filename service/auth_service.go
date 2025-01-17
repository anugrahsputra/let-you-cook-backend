package service

import (
	"errors"
	"let-you-cook/domain/dto"
	"let-you-cook/domain/model"
	"let-you-cook/repository"
	"let-you-cook/utils/jwt"
	"let-you-cook/utils/security"
	"time"

	"github.com/google/uuid"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("main")

type IAuthService interface {
	Register(user dto.ReqUserRegister) (dto.UserResp, error)
	Login(user dto.ReqUserLogin) (string, error)
}

type authService struct {
	repo repository.IAuthRepo
}

func NewAuthService(repo repository.IAuthRepo) *authService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Register(user dto.ReqUserRegister) (dto.UserResp, error) {
	exist, err := s.repo.GetUserExisting(user.Username)
	if err != nil {
		return exist.ToDTO(), err
	}

	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return exist.ToDTO(), err
	}

	newUser := model.User{
		Id:        uuid.New().String(),
		Username:  user.Username,
		Password:  hashedPassword,
		Email:     user.Email,
		CreatedAt: int(time.Now().Unix()),
		UpdatedAt: int(time.Now().Unix()),
	}

	err = s.repo.RegisterRepo(newUser)
	if err != nil {
		return dto.UserResp{}, err
	}

	return newUser.ToDTO(), nil
}

func (s *authService) Login(user dto.ReqUserLogin) (string, error) {

	exist, err := s.repo.CheckUserExistingForLogin(user.Username)
	if err != nil {
		return "", err
	}

	if exist.Username == "" || security.CheckPassword(exist.Password, user.Password) != nil {
		return "", errors.New("invalid credentials")
	}

	return jwt.GenerateToken(exist)

}
