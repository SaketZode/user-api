package userservice

import (
	"user-api/models/usermodels"
	"user-api/resterrors"
)

type UserService interface {
	CreateUser(user *usermodels.User) (*usermodels.User, *resterrors.RestError)
	GetUser(userid int) (*usermodels.User, *resterrors.RestError)
	GetAllUsers() (err *resterrors.RestError)
	UpdateUser() (err *resterrors.RestError)
	DeleteUser() (err *resterrors.RestError)
}
