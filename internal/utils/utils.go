package utils

import "github.com/RostislavOrlov/krp_admin/internal/dto"

func GeneratePassword(user dto.AddUserRequest) string {
	return user.FirstName
}
