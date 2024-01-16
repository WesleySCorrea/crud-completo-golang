package controller

import (
	"net/http"

	"github.com/WesleySCorrea/crud-completo-golang/src/configuration/logger"
	"github.com/WesleySCorrea/crud-completo-golang/src/configuration/validation"
	"github.com/WesleySCorrea/crud-completo-golang/src/controller/model/request"
	"github.com/WesleySCorrea/crud-completo-golang/src/model"
	"github.com/WesleySCorrea/crud-completo-golang/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "CreateUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "CreateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User Create successfuly", zap.String("journey", "CreateUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResonse(domain))
}
