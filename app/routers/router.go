package routers

import (
	"my-task-api/app/middlewares"
	"my-task-api/features/user/data"
	_userHandler "my-task-api/features/user/handler"
	_userService "my-task-api/features/user/service"
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

	// Definisikan rute untuk entitas User
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.CreateUser)
	e.GET("/users", userHandlerAPI.GetUsers, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.Update, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.DeleteUser, middlewares.JWTMiddleware())

}
