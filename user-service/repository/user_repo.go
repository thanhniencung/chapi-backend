package repository

import (
	internalModel "chapi-backend/chapi-internal/model"
	"chapi-backend/user-service/model"
	"context"
)

type UserRepository interface {
	CheckLogin(context context.Context, loginReq model.LoginRequest) (internalModel.User, error)
	Save(context context.Context, user internalModel.User) (internalModel.User, error)
	SelectById(context context.Context, userId string) (internalModel.User, error)
}
