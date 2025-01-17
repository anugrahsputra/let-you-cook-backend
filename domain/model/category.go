package model

import "let-you-cook/domain/dto"

type Category struct {
	Id        string `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	UserId    string `json:"user_id" bson:"user_id"`
	CreatedAt int    `json:"created_at" bson:"created_at"`
	UpdatedAt int    `json:"updated_at" bson:"updated_at"`
}

func (category *Category) ToDTO() dto.CategoryResp {
	return dto.CategoryResp{
		Id:     category.Id,
		Name:   category.Name,
		UserId: category.UserId,
	}
}
