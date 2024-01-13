package controller

import (
	"fmt"

	"github.com/WesleySCorrea/crud-completo-golang/src/configuration/rest_err"
	"github.com/WesleySCorrea/crud-completo-golang/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		rest_err := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorect filds, erro=%s\n", err.Error()))

		c.JSON(rest_err.Code, rest_err)
		return
	}

	fmt.Println(userRequest)
}
