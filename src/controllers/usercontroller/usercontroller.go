package usercontroller

import (
	"net/http"
	"strconv"
	"user-api/models/usermodels"
	"user-api/resterrors"
	"user-api/services/userservice"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service userservice.UserService
}

func New() *UserController {
	return &UserController{
		service: userservice.New(),
	}
}

func (usercontroller *UserController) CreateUser(c *gin.Context) {
	var user = usermodels.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		parseerr := resterrors.NewBadRequestError("Unable to parse user from JSON!")
		c.JSON(parseerr.HttpStatus, parseerr)
		return
	}

	result, creationerr := usercontroller.service.CreateUser(&user)
	if creationerr != nil {
		c.JSON(creationerr.HttpStatus, creationerr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (usercontroller *UserController) GetUser(c *gin.Context) {
	userid := c.Param("user_id")
	userId, err := strconv.ParseInt(userid, 10, 32)
	if err != nil {
		parseErr := resterrors.NewBadRequestError("Unable to parse userID from params!")
		c.JSON(parseErr.HttpStatus, parseErr)
		return
	}
	user, fetchErr := usercontroller.service.GetUser(int(userId))
	if fetchErr != nil {
		c.JSON(fetchErr.HttpStatus, fetchErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (usercontroller *UserController) GetAllUsers(c *gin.Context) {
	users, err := usercontroller.service.GetAllUsers()
	if err != nil {
		c.JSON(err.HttpStatus, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (usercontroller *UserController) UpdateUser(c *gin.Context) {
	userid := c.Param("user_id")
	userId, err := strconv.ParseInt(userid, 10, 32)
	if err != nil {
		parseErr := resterrors.NewBadRequestError("Unable to parse userID from params!")
		c.JSON(parseErr.HttpStatus, parseErr)
		return
	}
	var user = usermodels.User{}
	if parserr := c.ShouldBindJSON(&user); parserr != nil {
		resterr := resterrors.NewBadRequestError("Unable to parse user from JSON!")
		c.JSON(resterr.HttpStatus, resterr)
		return
	}

	result, resterr := usercontroller.service.UpdateUser(int(userId), &user)
	if resterr != nil {
		c.JSON(resterr.HttpStatus, resterr)
		return
	}
	c.JSON(http.StatusAccepted, result)
}

func (usercontroller *UserController) DeleteUser(c *gin.Context) {
	userid := c.Param("user_id")
	userId, converr := strconv.ParseInt(userid, 10, 32)
	if converr != nil {
		parseErr := resterrors.NewBadRequestError("Unable to parse user ID from params!")
		c.JSON(parseErr.HttpStatus, parseErr)
		return
	}
	user, err := usercontroller.service.DeleteUser(int(userId))
	if err != nil {
		c.JSON(err.HttpStatus, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
