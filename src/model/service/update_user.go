package service

import (
	"github.com/WesleySCorrea/crud-completo-golang/src/configuration/logger"
	"github.com/WesleySCorrea/crud-completo-golang/src/configuration/rest_err"
	"github.com/WesleySCorrea/crud-completo-golang/src/model"
	"go.uber.org/zap"
)

func (*userDomainService) UpdateUser(
	userId string, userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init UpdateUser model", zap.String("journey", "upateUser"))

	return nil
}
