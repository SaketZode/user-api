package userservice

import (
	"user-api/models/usermodels"
	"user-api/resterrors"
	"user-api/resterrors/errorparser"
)

type UserServiceImpl struct {
}

func (userservice *UserServiceImpl) CreateUser(user *usermodels.User) (*usermodels.User, *resterrors.RestError) {
	validationErr := user.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	err := user.CreateUser()
	if err != nil {
		return nil, resterrors.NewInternalServerError("Something went wrong while saving user details!")
	}

	return user, nil
}

func (userservice *UserServiceImpl) GetUser(userid int) (user *usermodels.User, err *resterrors.RestError) {
	user = &usermodels.User{}
	if err := user.GetUser(userid); err != nil {
		return nil, errorparser.ParseError(err)
	}
	return user, nil
}

func (userservice *UserServiceImpl) GetAllUsers() (err *resterrors.RestError) {
	return resterrors.NewNotImplementedError("Method not implemented!")
}

func (userservice *UserServiceImpl) UpdateUser() (err *resterrors.RestError) {
	return resterrors.NewNotImplementedError("Method not implemented!")
}

func (userservice *UserServiceImpl) DeleteUser() (err *resterrors.RestError) {
	return resterrors.NewNotImplementedError("Method not implemented!")
}
