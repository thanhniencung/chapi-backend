package handler

import (
	"chapi-backend/chapi-internal/encrypt"
	internalEncryt "chapi-backend/chapi-internal/encrypt"
	"chapi-backend/chapi-internal/helper"
	"chapi-backend/chapi-internal/middleware"
	internalModel "chapi-backend/chapi-internal/model"
	userServiceModel "chapi-backend/user-service/model"
	"chapi-backend/user-service/repository"
	"context"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/lib/pq"
	"net/http"
	"reflect"
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
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	if _,err := govalidator.ValidateStruct(req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	if !helper.IsValidPhoneNumber(req.Phone) {
		return helper.ResponseErr(c, http.StatusBadRequest, "Số điện thoại không hợp lệ")
	}

	req.Password = encrypt.MD5Hash(req.Password)
	req.UserId = internalEncryt.UUID()
	// Trường hợp muốn tạo user có role là Admin thì có thể truyền thêm 1 param đặc biết đã quy ước rồi kiểm tra
	// Hoặc backend sẽ cấp cho user 1 token đặc biệt để đăng ký thành user Admin
	// Ở đây chúng ta để mặc định là MEMBER
	if req.Phone == "0973901736" {
		req.Role = internalModel.ADMIN.String()
	} else {
		req.Role = internalModel.MEMBER.String()
	}

	ctx, _:= context.WithTimeout(c.Request().Context(), 10 * time.Second)
	user, err := m.UserRepo.Save(ctx, req)
	if err != nil {
		// Chú ý khi sử dụng cách này, bởi chúng ta đang hiện một lệnh write vào database
		// Cân nhắc select record để check record tồn tại hay chưa, chưa thì hãy insert
		if reflect.TypeOf(err).String() == reflect.TypeOf(&pq.Error{}).String() {
			pqErr := err.(*pq.Error)
			if pqErr.Code == "23505" { //duplicate key value violates unique constraint "users_phone_key"
				return helper.ResponseErr(c, http.StatusConflict, pqErr.Message)
			}
		}
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	token, err := middleware.GenToken(user)
	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	type userResponse struct {
		internalModel.User
		Token string `json:"token"`
	}

	helper.FormatUserResponse(&user)
	return helper.ResponseData(c, userResponse{user, token})
}

// Handler sử lý khi user đăng nhập tài khoản
// Response trả về sẽ kèm theo token để truy cập các api về sau
func (u *UserHandler) SignIn(c echo.Context) error {
	//fmt.Println(">>>", c.RealIP())
	req := userServiceModel.LoginRequest{}
	defer c.Request().Body.Close()

	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	req.Password = encrypt.MD5Hash(req.Password)
	ctx, _:= context.WithTimeout(c.Request().Context(), 10 * time.Second)
	user, err := u.UserRepo.CheckLogin(ctx, req)
	if err != nil {
		return helper.ResponseErr(c, http.StatusUnauthorized, err.Error())
	}

	token, err := middleware.GenToken(user)
	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	type userResponse struct {
		internalModel.User
		Token string `json:"token"`
	}

	helper.FormatUserResponse(&user)
	return helper.ResponseData(c, userResponse{user, token})
}

func (u *UserHandler) Profile(c echo.Context) error {
	defer c.Request().Body.Close()

	// Lấy thông tin user_id từ token
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*internalModel.JwtCustomClaims)

	ctx, _:= context.WithTimeout(c.Request().Context(), 10 * time.Second)
	user, err := u.UserRepo.SelectById(ctx, claims.UserId)

	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	if user == (internalModel.User{}) {
		return helper.ResponseErr(c, http.StatusNotFound, "Người dùng này không tồn tại")
	}

	helper.FormatUserResponse(&user)
	return helper.ResponseData(c, user)
}

func (u *UserHandler) List(c echo.Context) error {
	defer c.Request().Body.Close()

	// Lấy thông tin user_id từ token
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*internalModel.JwtCustomClaims)

	ctx, _:= context.WithTimeout(c.Request().Context(), 10 * time.Second)
	user, err := u.UserRepo.SelectAll(ctx, claims.UserId)

	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseData(c, user)
}


