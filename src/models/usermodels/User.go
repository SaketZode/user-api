package usermodels

import (
	"fmt"
	"user-api/databaseconnection/postgresdb"
	"user-api/resterrors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

func (user *User) Validate() *resterrors.RestError {
	if user.FirstName == "" {
		return resterrors.NewBadRequestError("First name field is mandatory!")
	}
	if user.LastName == "" {
		return resterrors.NewBadRequestError("Last name field is mandatory!")
	}
	if user.Email == "" {
		return resterrors.NewBadRequestError("Email field is mandatory!")
	}
	return nil
}

func (user *User) CreateUser() error {
	txn := postgresdb.DatabaseConnection.Begin()

	if createError := txn.Create(user).Error; createError != nil {
		txn.Rollback()
		return createError
	}

	txn.Commit()
	return nil
}

func (user *User) GetUser(userid int) (err error) {
	if err := postgresdb.DatabaseConnection.First(user, userid).Error; err != nil {
		fmt.Println("Error in fetching user: ", err)
		return err
	}
	return nil
}

func (user *User) UpdateUser() (err error) {
	return nil
}

func (user *User) DeleteUser() (err error) {
	return nil
}
