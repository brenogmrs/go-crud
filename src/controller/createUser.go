package controller

import (
	"fmt"

	"github.com/brenogmrs/go-crud/src/configuration/validation"
	"github.com/brenogmrs/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := validation.ValidateUserError(err)

		c.JSON(restError.Code, restError)
		return
	}

	fmt.Println(userRequest)

}
