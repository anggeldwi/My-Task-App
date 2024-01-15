package handler

import (
	"log"
	"my-task-api/app/middlewares"
	"my-task-api/features/project"
	"my-task-api/utils/responses"
	"net/http"
	"strconv"

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

func (handler *ProjectHandler) GetProjects(c echo.Context) error {
	// Extract ID user from JWT token
	userID := middlewares.ExtractTokenUserId(c)
	log.Println("UserID:", userID)

	// Memanggil fungsi logic untuk mendapatkan semua proyek milik pengguna
	projects, err := handler.projectService.SelectAll(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Failed to get projects by user ID", nil))
	}

	// Transform projects from project.Core to ProjectResponse
	var projectsResponse []ProjectResponse
	for _, projectCore := range projects {
		projectsResponse = append(projectsResponse, CoreToResponse(projectCore))
	}

	// Mengembalikan proyek dalam respons JSON
	return c.JSON(http.StatusOK, responses.WebResponse("Success read data.", projectsResponse))
}

func (handler *ProjectHandler) GetProjectByID(c echo.Context) error {
	// Extract project ID from path parameter
	projectIDStr := c.Param("projectid")
	log.Println("ProjectID:", projectIDStr)

	// Mengonversi projectID dari string ke int
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Invalid project ID", nil))
	}

	// Memanggil fungsi logic untuk mendapatkan satu proyek berdasarkan ID
	projects, err := handler.projectService.SelectByProjectID(projectID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Failed to get project by ID", nil))
	}

	// Transform satu proyek dari project.Core menjadi satu ProjectResponseID
	var projectsResponse []ProjectResponseID
	for _, project := range projects {
		projectResponse := CoreToResponseID(project)
		projectsResponse = append(projectsResponse, projectResponse)
	}

	// Mengembalikan proyek dalam respons JSON
	return c.JSON(http.StatusOK, responses.WebResponse("Success read data.", projectsResponse))
}

func (handler *ProjectHandler) Update(c echo.Context) error {
	// Extract project ID from path parameter
	projectIDStr := c.Param("projectid")
	log.Println("ProjectID:", projectIDStr)

	// Mengonversi projectID dari string ke int
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Invalid project ID", nil))
	}

	// Bind the request data
	var projectData = ProjectRequest{}
	errBind := c.Bind(&projectData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	// Convert request data to core model
	projectCore := RequestToCore(projectData)

	// Call the service to update the user profile
	err = handler.projectService.Update(projectID, projectCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data"+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *ProjectHandler) DeleteProject(c echo.Context) error {
	// Extract project ID from path parameter
	projectIDStr := c.Param("projectid")
	log.Println("ProjectID:", projectIDStr)

	// Mengonversi projectID dari string ke int
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Invalid project ID", nil))
	}

	// Call the service to delete the project
	err = handler.projectService.Delete(projectID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Failed to delete project", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Success delete project", nil))
}
