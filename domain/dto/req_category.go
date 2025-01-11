package dto

type ReqCategory struct {
	Id     string `json:"id" bson:"id"`
	UserId string `json:"user_id" bson:"user_id"`
	Name   string `json:"name" bson:"name"`
}

type ReqUpdateCategory struct {
	Name string `json:"name" bson:"name"`
}
