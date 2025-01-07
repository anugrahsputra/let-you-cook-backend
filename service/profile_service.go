package service

import (
	"errors"
	"let-you-cook/domain/dto"
	"let-you-cook/domain/model"
	"let-you-cook/repository"
	"time"

	"github.com/google/uuid"
)

type IProfileService interface {
	CreateProfile(userId string, email string, reqProfile dto.ReqProfile) error
	GetProfileByAccountId(userId string) (model.Profile, error)
}

type profileService struct {
	repoProfile repository.IProfileRepo
	repoUser    repository.IUserRepo
}

func NewProfileService(repoProfile repository.IProfileRepo, repoUser repository.IUserRepo) *profileService {
	return &profileService{
		repoProfile: repoProfile,
		repoUser:    repoUser,
	}
}

func (s *profileService) CreateProfile(userId string, email string, reqProfile dto.ReqProfile) error {
	exist, err := s.repoProfile.GetProfileByAccountId(userId)
	if err != nil {
		return err
	}
	if exist.Id != "" {
		return errors.New("profile already exist")
	}

	profile := model.Profile{
		Id:        uuid.New().String(),
		IdAccount: userId,
		Fullname:  reqProfile.Fullname,
		Address:   reqProfile.Address,
		Email:     email,
		Phone:     reqProfile.Phone,
		Bio:       reqProfile.Bio,
		UpdatedAt: int(time.Now().Unix()),
		CreatedAt: int(time.Now().Unix()),
	}

	err = s.repoProfile.CreateProfile(profile)
	if err != nil {
		return err
	}

	return nil
}

func (s *profileService) GetProfileByAccountId(userId string) (model.Profile, error) {
	profile, err := s.repoProfile.GetProfileByAccountId(userId)
	if err != nil {
		return model.Profile{}, err
	}
	return profile, nil

}
