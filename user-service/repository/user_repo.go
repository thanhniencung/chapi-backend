package repository

import (
	internalModel "chapi-backend/chapi-internal/model"
	"chapi-backend/user-service/model"
	"context"
)

type UserRepository interface {
	CheckLogin(context context.Context, loginReq model.LoginRequest) (internalModel.User, error)
}
