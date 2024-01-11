package data

import (
	"errors"
	"my-task-api/features/project"

	"gorm.io/gorm"
)

type projectQuery struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) project.ProjectDataInterface {
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
func (*projectQuery) SelectAll() ([]project.Core, error) {
	panic("unimplemented")
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
