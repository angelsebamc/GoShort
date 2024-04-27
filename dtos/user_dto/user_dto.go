package user_dto

import "time"

type UserDTO_Info struct {
	Username string    `json:"username" bson:"username"`
	Email    string    `json:"email" bson:"username"`
	Created  time.Time `json:"created" bson:"created"`
}

type UserDTO_Info_Token struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserDTO_Registration struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserDTO_Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
