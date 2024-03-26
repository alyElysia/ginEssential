package dto

import (
	"ginEssential/model"
)

// UserDto 设置返回的信息
type UserDto struct {
	Name string `json:"name"`
	Tel  string `json:"tel"`
}

func GetDto(user model.User) UserDto {
	return UserDto{
		Name: user.Name,
		Tel:  user.Tel,
	}
}
