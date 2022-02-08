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

func (userservice *UserServiceImpl) GetAllUsers() (users []usermodels.User, err *resterrors.RestError) {
	user := usermodels.User{}
	userslist, geterr := user.GetAllUsers()
	if geterr != nil {
		return nil, errorparser.ParseError(geterr)
	}
	return userslist, nil
}

func (userservice *UserServiceImpl) UpdateUser(userid int, user *usermodels.User) (*usermodels.User, *resterrors.RestError) {
	if err := user.UpdateUser(userid); err != nil {
		return nil, errorparser.ParseError(err)
	}
	return user, nil
}

func (userservice *UserServiceImpl) DeleteUser(userid int) (*usermodels.User, *resterrors.RestError) {
	user := &usermodels.User{}
	user.ID = uint(userid)
	if err := user.DeleteUser(); err != nil {
		return nil, errorparser.ParseError(err)
	}
	return user, nil
}
