package handler

import (
	"log"
	"my-task-api/app/middlewares"
	"my-task-api/features/user"
	"my-task-api/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	//mapping dari request ke core
	userCore := RequestToCore(newUser)
	errInsert := handler.userService.Create(userCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data"+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *UserHandler) GetUsers(c echo.Context) error {
	// extract id user from jwt token
	userID := middlewares.ExtractTokenUserId(c)
	log.Println("UserID:", userID)

	// Menggunakan idToken sebagai parameter fungsi SelectUser
	results, errSelect := handler.userService.SelectUser(userID)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}

	// proses mapping dari core ke response
	usersResult := CoreToResponseList(results)
	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", usersResult))
}

func (handler *UserHandler) Login(c echo.Context) error {
	var reqData = LoginRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}
	result, token, err := handler.userService.Login(reqData.Email, reqData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error login "+err.Error(), nil))
	}
	responseData := map[string]any{
		"token": token,
		"nama":  result.Name,
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success login", responseData))
}

func (handler *UserHandler) Update(c echo.Context) error {
	// Extract user ID from JWT token
	userID := middlewares.ExtractTokenUserId(c)
	log.Println("UserID:", userID)

	// Check if the user ID is valid
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("unauthorized", nil))
	}

	// Use the extracted user ID from the token
	idParam := userID

	// Bind the request data
	var userData = UserRequest{}
	errBind := c.Bind(&userData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	// Convert request data to core model
	userCore := RequestToCore(userData)

	// Call the service to update the user profile
	err := handler.userService.Update(idParam, userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data"+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *UserHandler) DeleteUser(c echo.Context) error {
	// Extract user ID from JWT token
	userID := middlewares.ExtractTokenUserId(c)
	log.Println("UserID:", userID)

	// Check if the user ID is valid
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("unauthorized", nil))
	}

	// Use the extracted user ID from the token
	idParam := userID

	// Call the service to delete the user profile
	err := handler.userService.Delete(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data"+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}
