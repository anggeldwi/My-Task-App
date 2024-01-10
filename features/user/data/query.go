package data

import (
	"errors"
	"my-task-api/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

// Insert implements user.UserDataInterface.
func (repo *userQuery) Insert(input user.Core) error {
	// proses mapping dari struct entities core ke model gorm
	userInputGorm := User{
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
		Role:        input.Role,
	}
	// simpan ke DB
	tx := repo.db.Create(&userInputGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

// SelectUser implements user.UserDataInterface.
func (repo *userQuery) SelectUser(id int) ([]user.Core, error) {
	var userDataGorm User
	tx := repo.db.Where("id = ?", id).First(&userDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// proses mapping dari struct gorm model ke struct core
	userDataCore := []user.Core{
		{
			ID:          userDataGorm.ID,
			Name:        userDataGorm.Name,
			Email:       userDataGorm.Email,
			Password:    userDataGorm.Password,
			Address:     userDataGorm.Address,
			PhoneNumber: userDataGorm.PhoneNumber,
			Role:        userDataGorm.Role,
			CreatedAt:   userDataGorm.CreatedAt,
			UpdatedAt:   userDataGorm.UpdatedAt,
		},
	}

	return userDataCore, nil
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (data *user.Core, err error) {
	var userGorm User
	tx := repo.db.Where("email = ?", email).First(&userGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}
	result := userGorm.ModelToCore()
	return &result, nil
}

// Update implements user.UserDataInterface.
func (repo *userQuery) Update(id int, input user.Core) error {
	dataGorm := CoreToModel(input)
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(dataGorm)
	if tx.Error != nil {
		// fmt.Println("err:", tx.Error)
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

// Delete implements user.UserDataInterface.
func (repo *userQuery) Delete(id int) error {
	tx := repo.db.Delete(&User{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}
	return nil
}
