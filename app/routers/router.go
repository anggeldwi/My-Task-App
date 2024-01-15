package routers

import (
	"my-task-api/app/middlewares"
	"my-task-api/features/user/data"
	_userHandler "my-task-api/features/user/handler"
	_userService "my-task-api/features/user/service"

	_projectData "my-task-api/features/project/data"
	_projectHandler "my-task-api/features/project/handler"
	_projectService "my-task-api/features/project/service"

	_taskData "my-task-api/features/task/data"
	_taskHandler "my-task-api/features/task/handler"
	_taskService "my-task-api/features/task/service"

	"my-task-api/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	//factory
	hashService := encrypts.NewHashService()
	userData := data.New(db)
	// userData := _userData.NewRaw(db)
	userService := _userService.New(userData, hashService)
	userHandlerAPI := _userHandler.New(userService)

	// Inisialisasi data dan service untuk entitas Project
	projectData := _projectData.NewProject(db)
	projectService := _projectService.NewProjectService(projectData)
	projectHandlerAPI := _projectHandler.NewProjectHandler(projectService)

	// Inisialisasi data dan service untuk entitas Task
	taskData := _taskData.NewTask(db)
	taskService := _taskService.NewTaskService(taskData)
	taskHandlerAPI := _taskHandler.NewTaskHandler(taskService)

	// Definisikan rute untuk entitas User
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.CreateUser)
	e.GET("/users", userHandlerAPI.GetUsers, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.Update, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.DeleteUser, middlewares.JWTMiddleware())

	// Definisikan rute untuk entitas Project
	e.POST("/projects", projectHandlerAPI.CreateProject, middlewares.JWTMiddleware())
	e.GET("/projects", projectHandlerAPI.GetProjects, middlewares.JWTMiddleware())
	e.GET("/projects/:projectid", projectHandlerAPI.GetProjectByID, middlewares.JWTMiddleware())
	e.PUT("/projects/:projectid", projectHandlerAPI.Update, middlewares.JWTMiddleware())
	e.DELETE("/projects/:projectid", projectHandlerAPI.DeleteProject, middlewares.JWTMiddleware())

	// Definisikan rute untuk entitas Task
	e.POST("/tasks", taskHandlerAPI.CreateTask, middlewares.JWTMiddleware())
	e.PUT("/tasks/:taskid", taskHandlerAPI.UpdateTask, middlewares.JWTMiddleware())
	e.DELETE("/tasks/:taskid", taskHandlerAPI.DeleteTask, middlewares.JWTMiddleware())
}
