package application

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"user-api/controllers/usercontroller"
)

var (
	router = gin.Default()
)

func mapURLS() {
	fmt.Println("Mapping urls....")
	usercontroller := usercontroller.New()
	router.POST("/user", usercontroller.CreateUser)
	router.GET("/user/:user_id", usercontroller.GetUser)
	router.PUT("/user/:user_id", usercontroller.UpdateUser)
	router.DELETE("/user/:user_id", usercontroller.DeleteUser)
}
