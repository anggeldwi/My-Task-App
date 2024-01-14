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

// SelectAllAll implements project.ProjectServiceInterface.
func (*productService) SelectAllAll() ([]project.Core, error) {
	panic("unimplemented")
}

// SelectByProjecttID implements project.ProjectServiceInterface.
func (*productService) SelectByProjecttID(id int) ([]project.Core, error) {
	panic("unimplemented")
}

// Update implements project.ProjectServiceInterface.
func (*productService) Update(id int, input project.Core) error {
	panic("unimplemented")
}

// Delete implements project.ProjectServiceInterface.
func (*productService) Delete(id int) error {
	panic("unimplemented")
}
