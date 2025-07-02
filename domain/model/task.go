package model

import (
	"let-you-cook/domain/dto"
)

type Task struct {
	Id          string   `json:"id" bson:"id"`
	UserId      string   `json:"user_id" bson:"user_id"`
	CategoryId  string   `json:"category_id" bson:"category_id"`
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	Status      string   `json:"status" bson:"status"`
	Priority    string   `json:"priority" bson:"priority"`
	CreatedAt   int      `json:"created_at" bson:"created_at"`
	UpdatedAt   int      `json:"updated_at" bson:"updated_at"`
	CompletedAt int      `json:"completed_at" bson:"completed_at"`
	Tags        []string `json:"tags" bson:"tags"`
}

func (t *Task) ToDTO() dto.TaskResp {
	return dto.TaskResp{
		Id:          t.Id,
		Title:       t.Title,
		Description: t.Description,
		CategoryId:  t.CategoryId,
		Status:      t.Status,
		Priority:    t.Priority,
		Tags:        t.Tags,
	}
}
