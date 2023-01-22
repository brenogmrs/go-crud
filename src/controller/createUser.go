package controller

import (
	"net/http"

	"github.com/brenogmrs/go-crud/src/configuration/logger"
	"github.com/brenogmrs/go-crud/src/configuration/validation"
	"github.com/brenogmrs/go-crud/src/controller/model/request"
	"github.com/brenogmrs/go-crud/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", 
		zap.String("journey", "create-user"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error validating user info", err,
			zap.String("journey", "create-user"),
		)

		restError := validation.ValidateUserError(err)
		c.JSON(restError.Code, restError)
		return
	}
	
	
	response := response.UserResponse{
		ID: "teste",
		Email: userRequest.Email,
		Name: userRequest.Name,
		Age: userRequest.Age,
	}

	logger.Info("User created successfully",
		zap.String("journey", "create-user"),
	)

	c.JSON(http.StatusOK, response)
}
