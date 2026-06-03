package dto

import "bemunair2026/server/database/entities"

type UserCreateRequest struct {
	Name     string  `json:"name" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,password"`
	Role     string  `json:"role" validate:"required,oneof=ADMIN MENTRI"`
	Ministry *string `json:"ministry"`
	Phone    *string `json:"phone"`
}

type UserUpdateRequest struct {
	Name     string  `json:"name"`
	Email    string  `json:"email" validate:"omitempty,email"`
	Role     string  `json:"role" validate:"omitempty,oneof=ADMIN MENTRI"`
	Ministry *string `json:"ministry"`
	Phone    *string `json:"phone"`
}

type UserResponse struct {
	ID       uint64  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Role     string  `json:"role"`
	Ministry *string `json:"ministry,omitempty"`
	Phone    *string `json:"phone,omitempty"`
}

func NewUserResponse(user *entities.User) UserResponse {
	if user == nil {
		return UserResponse{}
	}

	return UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		Ministry: user.Ministry,
		Phone:    user.Phone,
	}
}

func NewUserResponses(users []entities.User) []UserResponse {
	responses := make([]UserResponse, 0, len(users))
	for i := range users {
		responses = append(responses, NewUserResponse(&users[i]))
	}
	return responses
}
