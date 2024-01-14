package data

import (
	"errors"
	"my-task-api/features/task"

	"gorm.io/gorm"
)

type taskQuery struct {
	db *gorm.DB
}

// Delete implements task.TaskDataInterface.
func NewTask(db *gorm.DB) task.TaskDataInterface {
	return &taskQuery{
		db: db,
	}
}

// Insert implements task.TaskDataInterface.
func (repo *taskQuery) Insert(input task.Task) error {
	// proses mapping dari struct entities core ke model gorm
	taskInputGorm := Task{
		Task:      input.Task,
		ProjectID: input.ProjectID,
	}
	// simpan ke DB
	tx := repo.db.Create(&taskInputGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

// Update implements task.TaskDataInterface.
func (repo *taskQuery) Update(id int, input task.Task) error {
	dataGorm := CoreToModel(input)
	tx := repo.db.Model(&Task{}).Where("id = ?", id).Updates(dataGorm)
	if tx.Error != nil {
		// fmt.Println("err:", tx.Error)
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

func (repo *taskQuery) Delete(id int) error {
	tx := repo.db.Delete(&Task{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}
	return nil
}
