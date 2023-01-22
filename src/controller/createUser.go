package controller

import (
	"fmt"

	"github.com/brenogmrs/go-crud/src/configuration/rest_err"
	"github.com/brenogmrs/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, error=%s\n", err.Error()))

		c.JSON(restError.Code, restError)
		return
	}

	fmt.Println(userRequest)

}
