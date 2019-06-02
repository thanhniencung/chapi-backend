package repo_impl

import (
	"chapi-backend/chapi-internal/db"
	internalModel "chapi-backend/chapi-internal/model"
	"chapi-backend/user-service/model"
	"chapi-backend/user-service/repository"
	"context"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepository {
	return &UserRepoImpl{
		sql: sql,
	}
}

func (u *UserRepoImpl) CheckLogin(context context.Context, loginReq model.LoginRequest) (internalModel.User, error) {
	var user internalModel.User

	row := u.sql.Db.QueryRowxContext(context, "SELECT phone, password FROM users WHERE phone=? AND password=?",
						loginReq.Phone, loginReq.Password)

	err := row.Err()
	if err != nil {
		return user, err
	}

	err = row.StructScan(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

