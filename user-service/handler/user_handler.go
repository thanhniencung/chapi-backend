package handler

import (
	"chapi-backend/chapi-internal/middleware"
	internalModel "chapi-backend/chapi-internal/model"
	userServiceModel "chapi-backend/user-service/model"
	"chapi-backend/user-service/repository"
	"context"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type UserHandler struct {
	UserRepo repository.UserRepository
}

func (m *UserHandler) Register(c echo.Context) error {
	return nil
}

func (u *UserHandler) Login(c echo.Context) error {
	req := userServiceModel.LoginRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, internalModel.Response{
			StatusCode:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		})
	}
	fmt.Println("Request data ", req)

	ctx, _:= context.WithTimeout(c.Request().Context(), 10 * time.Second)
	user, err := u.UserRepo.CheckLogin(ctx, req)
	if err != nil {
		fmt.Println("ahihi")
		return c.JSON(http.StatusNotFound, internalModel.Response{
			StatusCode:  http.StatusNotFound,
			Message: err.Error(),
		})
	}

	token, err := middleware.GenToken()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internalModel.Response{
			StatusCode:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	type userResponse struct {
		internalModel.User
		Token string
	}

	return c.JSON(http.StatusOK, internalModel.Response{
		StatusCode: http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data: userResponse{user, token},
	})
}

func (m *UserHandler) Profile(c echo.Context) error {
	return nil
}

