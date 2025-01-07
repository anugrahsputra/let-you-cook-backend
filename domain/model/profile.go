package model

type Profile struct {
	Id           string `json:"id" bson:"id"`
	IdAccount    string `json:"id_account" bson:"id_account"`
	Fullname     string `json:"fullname" bson:"fullname"`
	Address      string `json:"address" bson:"address"`
	Email        string `json:"email" bson:"email"`
	Phone        string `json:"phone" bson:"phone"`
	PhotoProfile string `json:"photo_profile" bson:"photo_profile"`
	Bio          string `json:"bio" bson:"bio"`
	UpdatedAt    int    `json:"updated_at" bson:"updated_at"`
	CreatedAt    int    `json:"created_at" bson:"created_at"`
}
