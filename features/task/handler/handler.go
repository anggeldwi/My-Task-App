package handler

import (
	"log"
	"my-task-api/features/task"
	"my-task-api/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskService task.TaskServiceInterface
}

func NewTaskHandler(taskService task.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (handler *TaskHandler) CreateProject(c echo.Context) error {
	// // Extract ID user from JWT token
	// userID := middlewares.ExtractTokenUserId(c)
	// log.Println("UserID:", userID)

	// Bind data dari request
	newTask := TaskRequest{}
	errBind := c.Bind(&newTask)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	// Mapping request ke core
	taskCore := RequestToCore(newTask)

	// // Set ProjectID pada proyek yang akan dibuat
	// taskCore.ProjectID = uint(projectID)

	// Memanggil service untuk membuat proyek
	errInsert := handler.taskService.Create(taskCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data"+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *TaskHandler) UpdateTask(c echo.Context) error {
	// Extract projectID from URL parameter
	taskID := c.Param("taskid")
	log.Println("TaskID:", taskID)

	// Convert projectID to integer (assuming it's an integer)
	taskIDInt, err := strconv.Atoi(taskID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Invalid project ID", nil))
	}

	// Parse the request body to get the updated project details
	var updatedTask task.Task
	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Invalid request payload", nil))
	}

	// Call the Update method of the projectService
	if err := handler.taskService.Update(taskIDInt, updatedTask); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Failed to update project", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Success updating task.", nil))
}

func (handler *TaskHandler) DeleteTask(c echo.Context) error {
	// Extract projectID from URL parameter
	taskID := c.Param("taskid")
	log.Println("TaskID:", taskID)

	// Convert projectID to integer (assuming it's an integer)
	taskIDInt, err := strconv.Atoi(taskID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Invalid project ID", nil))
	}

	// Call the service to delete the user profile
	err = handler.taskService.Delete(taskIDInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data"+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}
