package dto

type ReqProfile struct {
	Fullname     string `json:"fullname" bson:"fullname"`
	Address      string `json:"address" bson:"address"`
	Phone        string `json:"phone" bson:"phone"`
	Bio          string `json:"bio" bson:"bio"`
	PhotoProfile string `json:"photo_profile" bson:"photo_profile"`
}

type ReqPatchProfile struct {
	Fullname     *string `json:"fullname" bson:"fullname" binding:"omitempty"`
	Address      *string `json:"address" bson:"address" binding:"omitempty"`
	Phone        *string `json:"phone" bson:"phone" binding:"omitempty"`
	Bio          *string `json:"bio" bson:"bio" binding:"omitempty"`
	PhotoProfile *string `json:"photo_profile" bson:"photo_profile" binding:"omitempty"`
}

type ProfileResp struct {
	Fullname     string `json:"fullname" bson:"fullname"`
	Address      string `json:"address" bson:"address"`
	Phone        string `json:"phone" bson:"phone"`
	Bio          string `json:"bio" bson:"bio"`
	PhotoProfile string `json:"photo_profile" bson:"photo_profile"`
}
