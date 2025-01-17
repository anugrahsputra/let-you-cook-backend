package model

import (
	"let-you-cook/domain/dto"
)

type PomodoroSession struct {
	Id            string `json:"id" bson:"id"`
	UserId        string `json:"user_id" bson:"user_id"`
	TaskId        string `json:"task_id" bson:"task_id"`
	Name          string `json:"name" bson:"name"`
	StartTime     int    `json:"start_time" bson:"start_time"`
	EndTime       int    `json:"end_time" bson:"end_time"`
	Status        string `json:"status" bson:"status"`
	FocusDuration int    `json:"focus_duration" bson:"focus_duration"`
	BreakDuration int    `json:"break_duration" bson:"break_duration"`
	CreatedAt     int    `json:"created_at" bson:"created_at"`
	UpdatedAt     int    `json:"updated_at" bson:"updated_at"`
}

func (session *PomodoroSession) ToDTO() dto.PomodoroSessionResp {
	return dto.PomodoroSessionResp{
		Name:          session.Name,
		TaskId:        session.TaskId,
		StartTime:     session.StartTime,
		EndTime:       session.EndTime,
		Status:        session.Status,
		FocusDuration: session.FocusDuration,
		BreakDuration: session.BreakDuration,
	}
}
