package dto

type UserResp struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
}
