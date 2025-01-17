package model

import "let-you-cook/domain/dto"

type Profile struct {
	Id           string `json:"id" bson:"id"`
	UserId       string `json:"user_id" bson:"user_id"`
	Fullname     string `json:"fullname" bson:"fullname"`
	Address      string `json:"address" bson:"address"`
	Email        string `json:"email" bson:"email"`
	Phone        string `json:"phone" bson:"phone"`
	PhotoProfile string `json:"photo_profile" bson:"photo_profile"`
	Bio          string `json:"bio" bson:"bio"`
	UpdatedAt    int    `json:"updated_at" bson:"updated_at"`
	CreatedAt    int    `json:"created_at" bson:"created_at"`
}

func (profile *Profile) ToDTO() dto.ProfileResp {
	return dto.ProfileResp{
		Fullname:     profile.Fullname,
		Address:      profile.Address,
		Phone:        profile.Phone,
		Bio:          profile.Bio,
		PhotoProfile: profile.PhotoProfile,
	}
}
