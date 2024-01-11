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

// Project is the gorm model for the project
type Project struct {
	gorm.Model
	Name        string
	UserID      uint
	Description string
	User        User
}

func CoreToModelProduct(core project.Core) Project {
	return Project{
		Name:        core.Name,
		Description: core.Description,
		UserID:      core.UserID,
	}
}
