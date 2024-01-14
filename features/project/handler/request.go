package handler

import "my-task-api/features/project"

type ProjectRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

func RequestToCore(input ProjectRequest) project.Core {
	return project.Core{
		Name:        input.Name,
		Description: input.Description,
	}
}
