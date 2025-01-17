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
	GetProfileByAccountId(userId string) (dto.ProfileResp, error)
	UpdateProfile(userId string, payload dto.ReqPatchProfile) (dto.ProfileResp, error)
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

	if exist.PhotoProfile != "" {

	}

	profile := model.Profile{
		Id:           uuid.New().String(),
		UserId:       userId,
		Fullname:     reqProfile.Fullname,
		Address:      reqProfile.Address,
		Email:        email,
		Phone:        reqProfile.Phone,
		Bio:          reqProfile.Bio,
		PhotoProfile: reqProfile.PhotoProfile,
		UpdatedAt:    int(time.Now().Unix()),
		CreatedAt:    int(time.Now().Unix()),
	}

	err = s.repoProfile.CreateProfile(profile)
	if err != nil {
		return err
	}

	return nil
}

func (s *profileService) GetProfileByAccountId(userId string) (dto.ProfileResp, error) {
	profile, err := s.repoProfile.GetProfileByAccountId(userId)
	if err != nil {
		return dto.ProfileResp{}, err
	}
	return profile.ToDTO(), nil

}

func (s *profileService) UpdateProfile(userId string, payload dto.ReqPatchProfile) (dto.ProfileResp, error) {
	existingProfile, err := s.repoProfile.GetProfileByAccountId(userId)
	if err != nil {
		return dto.ProfileResp{}, err
	}

	if payload.Fullname != nil {
		existingProfile.Fullname = *payload.Fullname
	}

	if payload.Address != nil {
		existingProfile.Address = *payload.Address
	}

	if payload.Phone != nil {
		existingProfile.Phone = *payload.Phone
	}

	if payload.Bio != nil {
		existingProfile.Bio = *payload.Bio
	}

	if payload.PhotoProfile != nil {
		existingProfile.PhotoProfile = *payload.PhotoProfile
	}

	err = s.repoProfile.UpdateProfile(existingProfile.Id, payload)
	if err != nil {
		return dto.ProfileResp{}, err
	}

	return existingProfile.ToDTO(), nil
}
