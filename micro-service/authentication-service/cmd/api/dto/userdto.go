package dto

import "authentication/cmd/api/model"

type UserDto struct {
	Name string `json:"name"`
	Mail string `json:"email"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name: user.Name,
		Mail: user.Email,
	}
}
