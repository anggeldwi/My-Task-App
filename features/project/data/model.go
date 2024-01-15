package data

import (
	"my-task-api/features/project"

	"gorm.io/gorm"
)

// User is the gorm model for the user
type User struct {
	gorm.Model
	Name        string
	Email       string
	Address     string
	PhoneNumber string
	Role        string
}

type Task struct {
	gorm.Model
	Task      string `json:"task" form:"task"`
	ProjectID uint   `json:"project_id" form:"project_id"`
	Status    string `json:"status" form:"status"`
}

// Project is the gorm model for the project
type Project struct {
	gorm.Model
	Name        string
	UserID      uint
	Description string
	User        User
	Tasks       []Task `gorm:"foreignKey:ProjectID"`
}

func CoreToModelProject(core project.Core) Project {
	return Project{
		Name:        core.Name,
		Description: core.Description,
		UserID:      core.UserID,
	}
}
