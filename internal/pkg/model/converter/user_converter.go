package converter

import (
	"golang-clean-architecture/internal/pkg/model"
	"golang-clean-architecture/internal/pkg/model/entity"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserToTokenResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		Token: user.Token,
	}
}
