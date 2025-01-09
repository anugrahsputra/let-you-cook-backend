package service

import (
	"errors"
	"let-you-cook/domain/dto"
	"let-you-cook/domain/model"
	"let-you-cook/repository"
	"time"

	"github.com/google/uuid"
)

type ITaskService interface {
	CreateTask(userId string, task dto.ReqTask) error
	GetTasks(userId string) ([]model.Task, error)
	UpdateTask(id string, userId string, update map[string]interface{}) (model.Task, error)
	DeleteTask(id string, userId string) (model.Task, error)
}

type taskService struct {
	repo     repository.ITaskRepo
	repoUser repository.IUserRepo
}

func NewTaskService(repo repository.ITaskRepo, repoUser repository.IUserRepo) *taskService {
	return &taskService{
		repo:     repo,
		repoUser: repoUser,
	}
}

func (s *taskService) CreateTask(userId string, task dto.ReqTask) error {
	idUser, err := s.repoUser.GetUserById(userId)

	if err != nil {
		return err
	}

	newTask := model.Task{
		Id:          uuid.New().String(),
		UserId:      idUser.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		CreatedAt:   int(time.Now().Unix()),
		UpdatedAt:   int(time.Now().Unix()),
		CompletedAt: 0,
		Tags:        task.Tags,
	}

	if err = s.repo.CreateTask(newTask); err != nil {
		return err
	}

	return nil

}

func (s *taskService) GetTasks(userId string) ([]model.Task, error) {
	tasks, err := s.repo.GetTasks(userId)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *taskService) UpdateTask(id string, userId string, update map[string]interface{}) (model.Task, error) {
	existingTask, err := s.repo.UpdateTask(id, userId, update)
	if err != nil {
		return model.Task{}, err
	}

	if existingTask.Id == "" {
		return model.Task{}, errors.New("task not found")
	}

	update["updated_at"] = int(time.Now().Unix())

	updatedTask, err := s.repo.UpdateTask(existingTask.Id, existingTask.UserId, update)

	if err != nil {
		return model.Task{}, err
	}

	return updatedTask, nil

}

func (s *taskService) DeleteTask(id string, userId string) (model.Task, error) {
	task, err := s.repo.DeleteTask(id, userId)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}
