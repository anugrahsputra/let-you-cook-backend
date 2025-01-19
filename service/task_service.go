package service

import (
	"let-you-cook/domain/dto"
	"let-you-cook/domain/model"
	"let-you-cook/repository"
	"time"

	"github.com/google/uuid"
)

type ITaskService interface {
	CreateTask(userId string, task dto.ReqTask) error
	GetTasks(userId string) ([]dto.TaskResp, error)
	GetTaskGroupedByCategory(userId string) ([]dto.TaskByCategoryGroupResp, error)
	UpdateTask(id string, userId string, payload dto.ReqPatchTask) (dto.TaskResp, error)
	DeleteTask(id string, userId string) (dto.TaskResp, error)
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
		CategoryId:  task.CategoryId,
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

func (s *taskService) GetTasks(userId string) ([]dto.TaskResp, error) {
	tasks, err := s.repo.GetTasks(userId)
	if err != nil {
		return nil, err
	}

	taskResponse := make([]dto.TaskResp, 0)
	for _, task := range tasks {
		taskResponse = append(taskResponse, task.ToDTO())
	}

	return taskResponse, nil
}

func (s *taskService) GetTaskGroupedByCategory(userId string) ([]dto.TaskByCategoryGroupResp, error) {
	tasks, err := s.repo.GetTaskGroupedByCategory(userId)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *taskService) UpdateTask(id string, userId string, payload dto.ReqPatchTask) (dto.TaskResp, error) {
	existingTask, err := s.repo.FindTask(id, userId)

	if err != nil {
		return dto.TaskResp{}, err
	}

	applyTaskPatch(&existingTask, payload)
	err = s.repo.UpdateTask(id, userId, existingTask)
	if err != nil {
		return dto.TaskResp{}, err
	}

	return existingTask.ToDTO(), nil

}

func (s *taskService) DeleteTask(id string, userId string) (dto.TaskResp, error) {
	existingTask, err := s.repo.FindTask(id, userId)
	if err != nil {
		return dto.TaskResp{}, err
	}

	err = s.repo.DeleteTask(id, userId)
	if err != nil {
		return dto.TaskResp{}, err
	}
	return existingTask.ToDTO(), nil
}

func applyTaskPatch(existingTask *model.Task, payload dto.ReqPatchTask) {
	if payload.Title != nil {
		existingTask.Title = *payload.Title
	}

	if payload.Description != nil {
		existingTask.Description = *payload.Description
	}

	if payload.Priority != nil {
		existingTask.Priority = *payload.Priority
	}

	if payload.Status != nil {
		existingTask.Status = *payload.Status
	}

	if payload.CategoryId != nil {
		existingTask.CategoryId = *payload.CategoryId
	}

}
