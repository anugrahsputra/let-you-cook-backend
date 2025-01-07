package dto

type ReqProfile struct {
	Fullname string `json:"fullname" bson:"fullname"`
	Address  string `json:"address" bson:"address"`
	Phone    string `json:"phone" bson:"phone"`
	Bio      string `json:"bio" bson:"bio"`
}
