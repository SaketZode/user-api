package userservice

import (
	"user-api/models/usermodels"
	"user-api/resterrors"
)

type UserService interface {
	CreateUser(user *usermodels.User) (*usermodels.User, *resterrors.RestError)
	GetUser(userid int) (*usermodels.User, *resterrors.RestError)
	GetAllUsers() (users []usermodels.User, err *resterrors.RestError)
	UpdateUser(int, *usermodels.User) (*usermodels.User, *resterrors.RestError)
	DeleteUser(userid int) (*usermodels.User, *resterrors.RestError)
}
