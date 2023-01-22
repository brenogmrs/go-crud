package routes

import (
	"github.com/brenogmrs/go-crud/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/users/:userId", controller.FindUserById)
	r.GET("/userByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:userId", controller.UpdateUser)
	r.DELETE("/users/:userId", controller.DeleteUser)

}
