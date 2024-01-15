package service

import (
	"errors"
	"my-task-api/features/project"
)

type productService struct {
	projectData project.ProjectDataInterface
}

func NewProjectService(projectData project.ProjectDataInterface) project.ProjectServiceInterface {
	return &productService{
		projectData: projectData,
	}
}

// Create implements project.ProjectServiceInterface.
func (service *productService) Create(input project.Core) error {
	if input.Name == "" {
		return errors.New("[validation] Name must be filled")
	}

	return service.projectData.Insert(input)
}

// SelectAll implements project.ProjectServiceInterface.
func (service *productService) SelectAll(userID int) ([]project.Core, error) {
	// Validasi userID
	if userID <= 0 {
		return nil, errors.New("invalid userID")
	}

	return service.projectData.SelectAll(userID)
}

// SelectByProjecttID implements project.ProjectServiceInterface.
func (service *productService) SelectByProjectID(id int) ([]project.Core, error) {
	// Validasi projectID
	if id <= 0 {
		return nil, errors.New("invalid projectID")
	}

	return service.projectData.SelectByProjectID(id)
}

// Update implements project.ProjectServiceInterface.
func (service *productService) Update(id int, input project.Core) error {
	// Validasi projectID
	if id <= 0 {
		return errors.New("invalid projectID")
	}

	return service.projectData.Update(id, input)
}

// Delete implements project.ProjectServiceInterface.
func (service *productService) Delete(id int) error {
	// Validasi projectID
	if id <= 0 {
		return errors.New("invalid projectID")
	}

	return service.projectData.Delete(id)
}
