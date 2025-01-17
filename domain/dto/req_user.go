package dto

type ReqUserRegister struct {
	Username string `json:"username" validate:"required,min=6,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type ReqUserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResp struct {
	Id       string `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
}
