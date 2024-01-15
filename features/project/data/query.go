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
func (repo *projectQuery) SelectByProjectID(id int) ([]project.Core, error) {
	var projectsDataGorm []Project

	// Menggunakan Preload untuk menggabungkan data dari tabel Task
	tx := repo.db.Preload("Tasks").Preload("User").Where("id = ?", id).Find(&projectsDataGorm)

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
			Task:        make([]*project.Task, len(value.Tasks)),
		}

		// Konversi data Task dari model ke core
		for i, task := range value.Tasks {
			projectCore.Task[i] = &project.Task{
				ID:        task.ID,
				Task:      task.Task,
				ProjectID: task.ProjectID,
				Status:    task.Status,
			}
		}

		projectsDataCore = append(projectsDataCore, projectCore)
	}

	return projectsDataCore, nil
}

// Update implements project.ProjectDataInterface.
func (repo *projectQuery) Update(id int, input project.Core) error {
	// Update implements user.UserDataInterface.
	dataGorm := CoreToModelProject(input)
	tx := repo.db.Model(&Project{}).Where("id = ?", id).Updates(dataGorm)
	if tx.Error != nil {
		// fmt.Println("err:", tx.Error)
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

// Delete implements project.ProjectDataInterface.
func (repo *projectQuery) Delete(id int) error {
	// Mulai transaksi
	tx := repo.db.Begin()

	// Hapus tugas terlebih dahulu yang terkait dengan proyek
	if err := tx.Where("project_id = ?", id).Delete(&Task{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Hapus proyek setelah berhasil menghapus tugas
	if err := tx.Delete(&Project{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaksi
	return tx.Commit().Error
}
