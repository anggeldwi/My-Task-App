package handler

import (
	"log"
	"my-task-api/app/middlewares"
	"my-task-api/features/project"
	"my-task-api/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectService project.ProjectServiceInterface
}

func NewProjectHandler(projectService project.ProjectServiceInterface) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

func (handler *ProjectHandler) CreateProject(c echo.Context) error {
	// Extract ID user from JWT token
	userID := middlewares.ExtractTokenUserId(c)
	log.Println("UserID:", userID)

	// Bind data dari request
	newProject := ProjectRequest{}
	errBind := c.Bind(&newProject)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	// Mapping request ke core
	projectCore := RequestToCore(newProject)

	// Set UserID pada proyek yang akan dibuat
	projectCore.UserID = uint(userID)

	// Memanggil service untuk membuat proyek
	errInsert := handler.projectService.Create(projectCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data"+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}
