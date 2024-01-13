package controller

import (
	"github.com/WesleySCorrea/crud-completo-golang/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func FindUserByID(c *gin.Context) {

	err := rest_err.NewBadRequestError("VocÊ chamou essa rota")
	c.JSON(err.Code, err)
}

func FindUserByEmail(c *gin.Context) {

}
