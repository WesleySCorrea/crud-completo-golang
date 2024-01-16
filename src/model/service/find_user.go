package service

import (
	"github.com/WesleySCorrea/crud-completo-golang/src/configuration/logger"
	"github.com/WesleySCorrea/crud-completo-golang/src/configuration/rest_err"
	"github.com/WesleySCorrea/crud-completo-golang/src/model"
	"go.uber.org/zap"
)

func (*userDomainService) FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init FindUser model", zap.String("journey", "findUser"))

	return nil, nil
}
