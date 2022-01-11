package util

import (
	err "github.com/timoteoBone/project-microservice/httpService/pkg/errors"

	"github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
)

func ValidateCreateUserRequest(user entities.CreateUserRequest) error {
	if user.Age < 1 || len(user.Name) < 2 || len(user.Pass) < 5 {
		return err.ErrInvalidDataForm
	}
	return nil
}
