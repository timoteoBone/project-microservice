package util

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
)

type RepositoryMock struct {
	mock.Mock
}

func NewRepositoryMock() RepositoryMock {
	return RepositoryMock{mock.Mock{}}
}

func (repo *RepositoryMock) CreateUser(ctx context.Context, rq entities.CreateUserRequest) (entities.CreateUserResponse, error) {

	args := repo.Mock.Called(ctx, rq)
	response := args[0]

	return response.(entities.CreateUserResponse), args.Error(1)
}

func (repo *RepositoryMock) GetUser(ctx context.Context, rq entities.GetUserRequest) (entities.GetUserResponse, error) {
	args := repo.Mock.Called(ctx, rq)
	response := args[0]

	return response.(entities.GetUserResponse), args.Error(1)

}
