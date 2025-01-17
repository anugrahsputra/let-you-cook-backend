package service

import (
	"let-you-cook/domain/dto"
	"let-you-cook/domain/model"
	"let-you-cook/repository"
	"time"

	"github.com/google/uuid"
)

type ICategoryService interface {
	CreateCategory(userId string, category dto.ReqCategory) error
	GetCategories(userId string, reqCategory dto.ReqCategory) ([]dto.CategoryResp, error)
	GetCategoryById(id string, userId string) (dto.CategoryResp, error)
	UpdateCategory(id string, userId string, update dto.ReqPatchCategory) (dto.CategoryResp, error)
	DeleteCategory(id string, userId string) (dto.CategoryResp, error)
}

type categoryService struct {
	repoCategory repository.ICategoryRepo
	repoUser     repository.IUserRepo
}

func NewCategoryService(repoCategory repository.ICategoryRepo, repoUser repository.IUserRepo) *categoryService {
	return &categoryService{
		repoCategory: repoCategory,
		repoUser:     repoUser,
	}
}

func (s *categoryService) CreateCategory(userId string, reqCategory dto.ReqCategory) error {
	user, err := s.repoUser.GetUserById(userId)

	if err != nil {
		return err
	}

	newCategory := model.Category{
		Id:        uuid.New().String(),
		Name:      reqCategory.Name,
		UserId:    user.Id,
		CreatedAt: int(time.Now().Unix()),
		UpdatedAt: int(time.Now().Unix()),
	}

	if err = s.repoCategory.CreateCategory(newCategory); err != nil {
		return err
	}
	return nil
}

func (s *categoryService) GetCategories(userId string, reqCategory dto.ReqCategory) ([]dto.CategoryResp, error) {
	categories, err := s.repoCategory.GetCategories(userId, reqCategory)
	if err != nil {
		return nil, err
	}

	categoryResp := make([]dto.CategoryResp, 0)
	for _, category := range categories {
		categoryResp = append(categoryResp, category.ToDTO())
	}

	return categoryResp, nil
}

func (s *categoryService) GetCategoryById(id string, userId string) (dto.CategoryResp, error) {
	category, err := s.repoCategory.GetCategoryById(id, userId)
	if err != nil {
		return category.ToDTO(), err
	}

	return category.ToDTO(), nil

}

func (s *categoryService) UpdateCategory(id string, userId string, payload dto.ReqPatchCategory) (dto.CategoryResp, error) {
	existingCategory, err := s.repoCategory.GetCategoryById(id, userId)
	if err != nil {
		return dto.CategoryResp{}, err
	}

	if payload.Name != nil {
		existingCategory.Name = *payload.Name
	}

	err = s.repoCategory.UpdateCategory(id, userId, existingCategory)
	if err != nil {
		return dto.CategoryResp{}, err
	}

	return existingCategory.ToDTO(), nil

}

func (s *categoryService) DeleteCategory(id string, userId string) (dto.CategoryResp, error) {
	deletedCategory, err := s.repoCategory.GetCategoryById(id, userId)
	if err != nil {
		return dto.CategoryResp{}, err
	}

	err = s.repoCategory.DeleteCategory(id, userId)
	if err != nil {
		return dto.CategoryResp{}, err
	}

	return deletedCategory.ToDTO(), nil
}
