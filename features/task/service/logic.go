package service

import (
	"errors"
	"my-task-api/features/task"
)

type taskService struct {
	taskData task.TaskDataInterface
}

func NewTaskService(taskData task.TaskDataInterface) task.TaskServiceInterface {
	return &taskService{
		taskData: taskData,
	}
}

// Create implements task.TaskServiceInterface.
func (service *taskService) Create(input task.Task) error {
	if input.Task == "" {
		return errors.New("[validation] Task must be filled")
	}

	return service.taskData.Insert(input)
}

// Update implements task.TaskServiceInterface.
func (service *taskService) Update(id int, input task.Task) error {
	// validasi
	if id <= 0 {
		return errors.New("invalid id")
	}
	// validasi inputan
	// ...
	err := service.taskData.Update(id, input)
	return err
}

// Delete implements task.TaskServiceInterface.
func (service *taskService) Delete(id int) error {
	// validasi
	if id <= 0 {
		return errors.New("invalid id")
	}

	err := service.taskData.Delete(id)
	return err
}
