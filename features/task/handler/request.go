package handler

import "my-task-api/features/task"

type TaskRequest struct {
	Task      string `json:"task" form:"task"`
	ProjectID uint   `json:"project_id" form:"project_id"`
	// Status    string `json:"status" form:"status"`
}

func RequestToCore(input TaskRequest) task.Task {
	return task.Task{
		Task:      input.Task,
		ProjectID: input.ProjectID,
		// Status:    input.Status,
	}
}
