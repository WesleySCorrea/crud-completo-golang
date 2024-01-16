package service

import (
	"fmt"

	"github.com/WesleySCorrea/crud-completo-golang/src/configuration/logger"
	"github.com/WesleySCorrea/crud-completo-golang/src/configuration/rest_err"
	"github.com/WesleySCorrea/crud-completo-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(ud)

	return nil
}
