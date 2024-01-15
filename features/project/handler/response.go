package handler

import (
	"my-task-api/features/project"
)

type ProjectResponse struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

type TaskResponse struct {
	ID        uint   `json:"id" form:"id"`
	Task      string `json:"task" form:"task"`
	ProjectID uint   `json:"project_id" form:"project_id"`
	Status    string `json:"status" form:"status"`
}

type ProjectResponseID struct {
	ID          uint           `json:"id" form:"id"`
	Name        string         `json:"name" form:"name"`
	Description string         `json:"description" form:"description"`
	Task        []TaskResponse `gorm:"foreignKey:ProjectID" json:"task" form:"task"`
}

func CoreToResponse(data project.Core) ProjectResponse {
	return ProjectResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
	}
}

func CoreToResponseID(data project.Core) ProjectResponseID {
	var tasksResponse []TaskResponse
	for _, task := range data.Task {
		taskResponse := TaskResponse{
			ID:        task.ID,
			Task:      task.Task,
			ProjectID: task.ProjectID,
			Status:    task.Status,
		}
		tasksResponse = append(tasksResponse, taskResponse)
	}

	return ProjectResponseID{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Task:        tasksResponse,
	}
}
