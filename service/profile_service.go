package service

import (
	"errors"
	"let-you-cook/config"
	"let-you-cook/domain/dto"
	"let-you-cook/domain/model"
	"let-you-cook/repository"
	"let-you-cook/utils/minio"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type IProfileService interface {
	CreateProfile(userId string, email string, reqProfile dto.ReqProfile) error
	GetProfileByAccountId(userId string) (dto.ProfileResp, error)
	UpdateProfile(userId string, payload dto.ReqPatchProfile) (dto.ProfileResp, error)
	UploadProfilePicture(userId string, file *multipart.FileHeader) (dto.ProfileResp, error)
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
	_, err := s.repoProfile.GetProfileByAccountId(userId)
	if err == nil {
		return errors.New("profile already exist")
	}

	if err != nil && err.Error() != "profile not found" {
		return err
	}

	

	profile := model.Profile{
		Id:           uuid.New().String(),
		UserId:       userId,
		Fullname:     reqProfile.Fullname,
		Address:      reqProfile.Address,
		Email:        email,
		Phone:        reqProfile.Phone,
		Bio:          reqProfile.Bio,
		PhotoProfile: "https://api.dicebear.com/9.x/lorelei/svg",
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
	profileResp := profile.ToDTO()
	profileResp.PhotoProfile = minio_util.TrimMinioURLPrefix(profileResp.PhotoProfile, config.MinioEndpoint, config.MinioUseSSL)
	return profileResp, nil

}

func (s *profileService) UpdateProfile(userId string, payload dto.ReqPatchProfile) (dto.ProfileResp, error) {
	profile, err := s.repoProfile.GetProfileByAccountId(userId)
	if err != nil {
		return dto.ProfileResp{}, err
	}

	applyProfilePatch(&profile, payload)

	err = s.repoProfile.UpdateProfile(userId, profile)
	if err != nil {
		return dto.ProfileResp{}, err
	}

	profileResp := profile.ToDTO()
	profileResp.PhotoProfile = minio_util.TrimMinioURLPrefix(profileResp.PhotoProfile, config.MinioEndpoint, config.MinioUseSSL)
	return profileResp, nil
}

func applyProfilePatch(existingProfile *model.Profile, payload dto.ReqPatchProfile) {
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
}

func (s *profileService) UploadProfilePicture(userId string, file *multipart.FileHeader) (dto.ProfileResp, error) {
	profile, err := s.repoProfile.GetProfileByAccountId(userId)
	if err != nil {
		return dto.ProfileResp{}, err
	}

	// upload to minio
	url, err := minio_util.UploadPhoto(file)
	if err != nil {
		return dto.ProfileResp{}, err
	}

	profile.PhotoProfile = url

	err = s.repoProfile.UpdateProfile(userId, profile)
	if err != nil {
		return dto.ProfileResp{}, err
	}

	profileResp := profile.ToDTO()
	profileResp.PhotoProfile = minio_util.TrimMinioURLPrefix(profileResp.PhotoProfile, config.MinioEndpoint, config.MinioUseSSL)
	return profileResp, nil
}
