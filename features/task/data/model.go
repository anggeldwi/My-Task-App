package data

import (
	"my-task-api/features/task"

	"gorm.io/gorm"
)

// struct user gorm model
type Task struct {
	gorm.Model
	// ID          uint `gorm:"primaryKey"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
	// DeletedAt   gorm.DeletedAt `gorm:"index"`
	ID        uint
	Task      string
	ProjectID uint
	Status    string
}

func CoreToModel(input task.Task) Task {
	return Task{
		Task:      input.Task,
		ProjectID: input.ProjectID,
		Status:    input.Status,
	}
}

func (u Task) ModelToCore() task.Task {
	return task.Task{
		ID:        u.ID,
		Task:      u.Task,
		ProjectID: u.ProjectID,
		Status:    u.Status,
	}
}
