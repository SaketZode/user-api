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

func (user *User) Copy(newuser *User) {
	user.Age = newuser.Age
	user.Email = newuser.Email
	user.FirstName = newuser.FirstName
	user.LastName = newuser.LastName
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

func (user *User) GetAllUsers() (userslist []User, err error) {
	if err := postgresdb.DatabaseConnection.Find(&userslist).Error; err != nil {
		fmt.Println("Error while fetching users list:", err)
		return nil, err
	}
	return userslist, nil
}

func (user *User) UpdateUser(userid int) error {
	existing := &User{}
	txn := postgresdb.DatabaseConnection.Begin()
	if err := txn.First(existing, userid).Error; err != nil {
		fmt.Println("Error in fetching user while update:", err)
		txn.Rollback()
		return err
	}

	existing.Copy(user)

	if updateErr := txn.Save(existing).Error; updateErr != nil {
		fmt.Println("Error in updating user:", updateErr)
		txn.Rollback()
		return updateErr
	}
	txn.Commit()
	return nil
}

func (user *User) DeleteUser() (err error) {
	txn := postgresdb.DatabaseConnection.Begin()
	if geterr := txn.First(user, user.ID).Error; geterr != nil {
		fmt.Println("Error while getting record for deletion:", geterr)
		txn.Rollback()
		return geterr
	}
	if err := txn.Delete(user, user.ID).Error; err != nil {
		fmt.Println("Error while deleting record:", err)
		txn.Rollback()
		return err
	}
	txn.Commit()
	return nil
}
