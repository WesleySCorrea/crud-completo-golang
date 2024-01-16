package service

import (
	"github.com/WesleySCorrea/crud-completo-golang/src/configuration/logger"
	"github.com/WesleySCorrea/crud-completo-golang/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*userDomainService) DeleteUser(string) *rest_err.RestErr {

	logger.Info("Init DeleteUser model", zap.String("journey", "deleteUser"))

	return nil
}
