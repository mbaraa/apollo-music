package entities

import "github.com/mbaraa/apollo-music/enums"

type User struct {
	FullName string           `json:"fullName"`
	PublicId string           `json:"publicId"`
	Email    string           `json:"email"`
	Password string           `json:"password"`
	Status   enums.UserStatus `json:"status"`
}
