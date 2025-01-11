package service

import (
	"errors"
	"let-you-cook/domain/dto"
	"let-you-cook/domain/model"
	"let-you-cook/repository"
	"time"

	"github.com/google/uuid"
)

type ICategoryService interface {
	CreateCategory(userId string, category dto.ReqCategory) error
	GetCategories(userId string) ([]model.Category, error)
	GetCategoryById(id string, userId string) (model.Category, error)
	UpdateCategory(id string, userId string, update dto.ReqUpdateCategory) (model.Category, error)
	DeleteCategory(id string, userId string) (model.Category, error)
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

func (s *categoryService) CreateCategory(userId string, category dto.ReqCategory) error {
	user, err := s.repoUser.GetUserById(userId)

	if err != nil {
		return err
	}

	newCategory := model.Category{
		Id:        uuid.New().String(),
		Name:      category.Name,
		UserId:    user.Id,
		CreatedAt: int(time.Now().Unix()),
		UpdatedAt: int(time.Now().Unix()),
	}

	if err = s.repoCategory.CreateCategory(newCategory); err != nil {
		return err
	}
	return nil
}

func (s *categoryService) GetCategories(userId string) ([]model.Category, error) {
	categories, err := s.repoCategory.GetCategories(userId)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *categoryService) GetCategoryById(id string, userId string) (model.Category, error) {
	category, err := s.repoCategory.GetCategoryById(id, userId)

	if err != nil {
		return model.Category{}, err
	}

	return category, nil

}

func (s *categoryService) UpdateCategory(id string, userId string, update dto.ReqUpdateCategory) (model.Category, error) {
	updatedCategory, err := s.repoCategory.UpdateCategory(id, userId, update)

	if err != nil {
		if err.Error() == "category not found" {
			return model.Category{}, errors.New("category not found")
		}
		return model.Category{}, err
	}

	return updatedCategory, nil

}

func (s *categoryService) DeleteCategory(id string, userId string) (model.Category, error) {
	deletedCategory, err := s.repoCategory.DeleteCategory(id, userId)
	if err != nil {
		if err.Error() == "category not found" {
			return model.Category{}, errors.New("category not found")
		}
		return model.Category{}, err
	}
	return deletedCategory, nil
}
