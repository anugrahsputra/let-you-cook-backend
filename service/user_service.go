package service

import (
	"let-you-cook/domain/model"
	"let-you-cook/repository"
)

type IUserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserById(userID string) (model.User, error)
}

type userService struct {
	repo repository.IUserRepo
}

func NewUserService(repo repository.IUserRepo) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserById(userId string) (model.User, error) {
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return user, err
	}
	return user, nil
}
