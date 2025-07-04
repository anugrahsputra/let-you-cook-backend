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
	UpdateSession(id string, userId string, payload dto.ReqPatchSession) (dto.PomodoroSessionResp, error)
	StartSession(id string, userId string) (dto.PomodoroSessionResp, error)
	EndSession(id string, userId string) (dto.PomodoroSessionResp, error)
	GetAllSessions(userId string) ([]dto.PomodoroSessionResp, error)
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

func (s *sessionService) UpdateSession(id string, userId string, payload dto.ReqPatchSession) (dto.PomodoroSessionResp, error) {
	session, err := s.repoSession.GetSessionById(id, userId)
	if err != nil {
		return dto.PomodoroSessionResp{}, err
	}

	applySessionPatch(&session, payload)
	err = s.repoSession.UpdateSession(id, userId, session)
	if err != nil {
		return dto.PomodoroSessionResp{}, err
	}

	return session.ToDTO(), nil
}

func (s *sessionService) GetAllSessions(userId string) ([]dto.PomodoroSessionResp, error) {
	sessions, err := s.repoSession.GetAllSessions(userId)
	if err != nil {
		return nil, err
	}

	sessionResp := make([]dto.PomodoroSessionResp, 0)
	for _, session := range sessions {
		sessionResp = append(sessionResp, session.ToDTO())
	}

	return sessionResp, nil
}

func (s *sessionService) StartSession(id string, userId string) (dto.PomodoroSessionResp, error) {
	session, err := s.repoSession.GetSessionById(id, userId)
	if err != nil {
		return dto.PomodoroSessionResp{}, err
	}

	if session.Status != "PENDING" {
		return dto.PomodoroSessionResp{}, nil // or an error indicating the session cannot be started
	}

	session.Status = "ACTIVE"
	session.StartTime = int(time.Now().Unix())
	session.UpdatedAt = int(time.Now().Unix())

	err = s.repoSession.StartSession(id, userId)
	if err != nil {
		return dto.PomodoroSessionResp{}, err
	}

	return session.ToDTO(), nil
}

func (s *sessionService) EndSession(id string, userId string) (dto.PomodoroSessionResp, error) {
	session, err := s.repoSession.GetSessionById(id, userId)
	if err != nil {
		return dto.PomodoroSessionResp{}, err
	}

	if session.Status != "ACTIVE" {
		return dto.PomodoroSessionResp{}, nil // or an error indicating the session cannot be ended
	}

	session.Status = "COMPLETED"
	session.EndTime = int(time.Now().Unix())
	session.UpdatedAt = int(time.Now().Unix())

	err = s.repoSession.EndSession(id, userId)
	if err != nil {
		return dto.PomodoroSessionResp{}, err
	}

	return session.ToDTO(), nil
}

func applySessionPatch(session *model.PomodoroSession, payload dto.ReqPatchSession) {

	if payload.Name != nil {
		session.Name = *payload.Name
	}

	if payload.Status != nil {
		session.Status = *payload.Status
	}

	if payload.TaskId != nil {
		session.TaskId = *payload.TaskId
	}

	if payload.StartTime != nil {
		session.StartTime = *payload.StartTime
	}

	if payload.EndTime != nil {
		session.EndTime = *payload.EndTime
	}
}
