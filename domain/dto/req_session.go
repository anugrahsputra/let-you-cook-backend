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
