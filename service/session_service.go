package service

import (
	"let-you-cook/domain/dto"
	"let-you-cook/domain/model"
	"let-you-cook/repository"
	"time"

	"github.com/google/uuid"
)

type ISessionService interface {
	CreateSession(userId string, req dto.ReqCreateSession) error
	StartSession(id string, userId string) error
	EndSession(id string, userId string) error
	GetAllSessions(userId string) ([]model.PomodoroSession, error)
}

type sessionService struct {
	repoSession repository.ISessionRepo
	repoUser    repository.IUserRepo
}

func NewSessionService(repoSession repository.ISessionRepo, repoUser repository.IUserRepo) *sessionService {
	return &sessionService{
		repoSession: repoSession,
		repoUser:    repoUser,
	}
}

func (s *sessionService) CreateSession(userId string, req dto.ReqCreateSession) error {
	user, err := s.repoUser.GetUserById(userId)
	if err != nil {
		return err
	}

	newSession := model.PomodoroSession{
		Id:            uuid.New().String(),
		UserId:        user.Id,
		TaskId:        req.TaskId,
		Name:          req.Name,
		FocusDuration: req.FocusDuration,
		BreakDuration: req.BreakDuration,
		StartTime:     0,
		EndTime:       0,
		Status:        "PENDING",
		CreatedAt:     int(time.Now().Unix()),
		UpdatedAt:     int(time.Now().Unix()),
	}

	if err = s.repoSession.CreateSession(newSession); err != nil {
		return err
	}

	return nil
}

func (s *sessionService) StartSession(id string, userId string) error {
	err := s.repoSession.StartSession(id, userId)
	if err != nil {
		return err
	}
	return nil

}

func (s *sessionService) EndSession(id string, userId string) error {
	err := s.repoSession.EndSession(id, userId)

	if err != nil {
		return err
	}

	return nil

}

func (s *sessionService) GetAllSessions(userId string) ([]model.PomodoroSession, error) {
	sessions, err := s.repoSession.GetAllSessions(userId)

	if err != nil {
		return nil, err
	}

	return sessions, nil
}
