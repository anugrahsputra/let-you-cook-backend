package service

import (
	"let-you-cook/domain/dto"
	"let-you-cook/repository"
)

type IUserService interface {
	GetAllUsers() ([]dto.UserResp, error)
	GetUserById(userID string) (dto.UserResp, error)
}

type userService struct {
	repo repository.IUserRepo
}

func NewUserService(repo repository.IUserRepo) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetAllUsers() ([]dto.UserResp, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var userResp []dto.UserResp
	for _, user := range users {
		userResp = append(userResp, user.ToDTO())
	}

	return userResp, nil
}

func (s *userService) GetUserById(userId string) (dto.UserResp, error) {
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return dto.UserResp{}, err
	}
	return user.ToDTO(), nil
}
