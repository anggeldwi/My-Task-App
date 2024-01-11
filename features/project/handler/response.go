package handler

import "my-task-api/features/project"

type ProjectResponse struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	UserID      uint   `json:"user_id" form:"user_id" gorm:"index"`
	Description string `json:"description" form:"description"`
}

func CoreToResponse(data project.Core) ProjectResponse {
	return ProjectResponse{
		ID:          data.ID,
		Name:        data.Name,
		UserID:      data.UserID,
		Description: data.Description,
	}
}

func CoreToResponseList(data []project.Core) []ProjectResponse {
	var results []ProjectResponse
	for _, v := range data {
		results = append(results, CoreToResponse(v))
	}
	return results
}
