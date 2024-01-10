package databases

import (
	"gorm.io/gorm"
)

// struct user gorm model
type User struct {
	gorm.Model
	// ID          uint `gorm:"primaryKey"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
	// DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string `json:"name" form:"name"`
	Email       string `gorm:"unique" json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Address     string `json:"address" form:"address"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Role        string `json:"role" form:"role"`
}

// struct project gorm model
type Project struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	UserID      uint   `json:"user_id" form:"user_id" gorm:"index"`
	Description string `json:"description" form:"description"`
	User        User   `gorm:"foreignKey:UserID"`
}

type TaskStatus string

const (
	Completed    TaskStatus = "completed"
	NotCompleted TaskStatus = "not completed"
)

// struct task gorm model
type Task struct {
	gorm.Model
	Task      string     `json:"name" form:"name"`
	ProjectID uint       `json:"project_id" form:"project_id" gorm:"index"`
	Status    TaskStatus `json:"status" form:"status"`
	Project   Project    `gorm:"foreignKey:ProjectID"`
}
