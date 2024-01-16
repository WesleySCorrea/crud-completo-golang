package view

import (
	"github.com/WesleySCorrea/crud-completo-golang/src/controller/model/response"
	"github.com/WesleySCorrea/crud-completo-golang/src/model"
)

func ConvertDomainToResonse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
