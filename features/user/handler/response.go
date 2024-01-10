package handler

import "my-task-api/features/user"

type UserResponse struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Address     string `json:"address" form:"address"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Role        string `json:"role" form:"role"`
}

func CoreToResponse(data user.Core) UserResponse {
	return UserResponse{
		ID:          data.ID,
		Name:        data.Name,
		Email:       data.Email,
		Address:     data.Address,
		PhoneNumber: data.PhoneNumber,
		Role:        data.Role,
	}
}

func CoreToResponseList(data []user.Core) []UserResponse {
	var results []UserResponse
	for _, v := range data {
		results = append(results, CoreToResponse(v))
	}
	return results
}
