package handler

import (
	"chapi-backend/chapi-internal/encrypt"
	"chapi-backend/chapi-internal/middleware"
	internalModel "chapi-backend/chapi-internal/model"
	userServiceModel "chapi-backend/user-service/model"
	"chapi-backend/user-service/repository"
	"context"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type UserHandler struct {
	UserRepo repository.UserRepository
}

// Handler sử lý khi user đăng ký tài khoản
// Response trả về sẽ kèm theo token để truy cập các api về sau
func (m *UserHandler) SignUp(c echo.Context) error {
	req := internalModel.User{}
	defer c.Request().Body.Close()

	if err := c.Bind(&req); err != nil {
		return ResponseErr(c, http.StatusBadRequest)
	}

	if _,err := govalidator.ValidateStruct(req); err != nil {
		return ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	req.Password = encrypt.MD5Hash(req.Password)
	ctx, _:= context.WithTimeout(c.Request().Context(), 10 * time.Second)

	user, err := m.UserRepo.Save(ctx, req)
	if err != nil {
		return ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	token, err := middleware.GenToken()
	if err != nil {
		return ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	type userResponse struct {
		internalModel.User
		Token string `json:"token"`
	}

	return ResponseData(c, userResponse{user, token})
}

// Handler sử lý khi user đăng nhập tài khoản
// Response trả về sẽ kèm theo token để truy cập các api về sau
func (u *UserHandler) SignIn(c echo.Context) error {
	req := userServiceModel.LoginRequest{}
	defer c.Request().Body.Close()

	if err := c.Bind(&req); err != nil {
		return ResponseErr(c, http.StatusBadRequest)
	}

	req.Password = encrypt.MD5Hash(req.Password)
	ctx, _:= context.WithTimeout(c.Request().Context(), 10 * time.Second)

	user, err := u.UserRepo.CheckLogin(ctx, req)
	if err != nil {
		return ResponseErr(c, http.StatusNotFound, err.Error())
	}

	token, err := middleware.GenToken()
	if err != nil {
		return ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	type userResponse struct {
		internalModel.User
		Token string `json:"token"`
	}

	return ResponseData(c, userResponse{user, token})
}

func (m *UserHandler) Profile(c echo.Context) error {
	return nil
}

