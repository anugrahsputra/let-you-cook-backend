package dto

type ReqCreateSession struct {
	Name          string `json:"name" bson:"name"`
	TaskId        string `json:"task_id" bson:"task_id"`
	FocusDuration int    `json:"focus_duration" bson:"focus_duration"`
	BreakDuration int    `json:"break_duration" bson:"break_duration"`
}

type ReqStartSession struct {
	Status string `json:"status" bson:"status"`
}

type ReqEndSession struct {
	Status string `json:"status" bson:"status"`
}

type ReqPatchSession struct {
	Name      *string `json:"name" bson:"name" binding:"omitempty"`
	Status    *string `json:"status" bson:"status" binding:"omitempty"`
	StartTime *int    `json:"start_time" bson:"start_time" binding:"omitempty"`
	EndTime   *int    `json:"end_time" bson:"end_time" binding:"omitempty"`
}

type PomodoroSessionResp struct {
	Id            string `json:"id" bson:"id"`
	Name          string `json:"name" bson:"name"`
	TaskId        string `json:"task_id" bson:"task_id"`
	StartTime     int    `json:"start_time" bson:"start_time"`
	EndTime       int    `json:"end_time" bson:"end_time"`
	Status        string `json:"status" bson:"status"`
	FocusDuration int    `json:"focus_duration" bson:"focus_duration"`
	BreakDuration int    `json:"break_duration" bson:"break_duration"`
}

type SessionEndSessionResp struct {
	PomodoroSessionResp
	Calculate int `json:"calculate" bson:"calculate"`
}
