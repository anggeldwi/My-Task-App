package data

import (
	"errors"
	"my-task-api/features/project"

	"gorm.io/gorm"
)

type projectQuery struct {
	db *gorm.DB
}

func NewProject(db *gorm.DB) project.ProjectDataInterface {
	return &projectQuery{
		db: db,
	}
}

// Insert implements project.ProjectDataInterface.
func (repo *projectQuery) Insert(input project.Core) error {
	projectInputGorm := Project{
		Name:        input.Name,
		Description: input.Description,
		UserID:      input.UserID,
	}

	tx := repo.db.Create(&projectInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil

}

// SelectAll implements project.ProjectDataInterface.
func (repo *projectQuery) SelectAll(userID int) ([]project.Core, error) {
	var projectsDataGorm []Project

	tx := repo.db.Where("user_id = ?", userID).Find(&projectsDataGorm)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var projectsDataCore []project.Core
	for _, value := range projectsDataGorm {
		var projectCore = project.Core{
			ID:          value.ID,
			Name:        value.Name,
			UserID:      value.UserID,
			Description: value.Description,
		}
		projectsDataCore = append(projectsDataCore, projectCore)
	}

	return projectsDataCore, nil
}

// SelectByProjecttID implements project.ProjectDataInterface.
func (*projectQuery) SelectByProjecttID(id int) ([]project.Core, error) {
	panic("unimplemented")
}

// Update implements project.ProjectDataInterface.
func (*projectQuery) Update(id int, input project.Core) error {
	panic("unimplemented")
}

// Delete implements project.ProjectDataInterface.
func (*projectQuery) Delete(id int) error {
	panic("unimplemented")
}
