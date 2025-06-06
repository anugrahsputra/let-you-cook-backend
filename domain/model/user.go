package model

import "let-you-cook/domain/dto"

type User struct {
	Id        string `json:"id" bson:"id"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	Email     string `json:"email" bson:"email"`
	CreatedAt int    `json:"created_at" bson:"created_at"`
	UpdatedAt int    `json:"updated_at" bson:"updated_at"`
}

func (user *User) ToDTO() dto.UserResp {
	return dto.UserResp{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}
