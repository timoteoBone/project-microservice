package utils

import (
	"context"
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/stretchr/testify/mock"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
)

func NewMock(logger log.Logger) (*sql.DB, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		level.Error(logger).Log("error opnening a sql connection", err)
	}

	return db, mock
}

type RepoSitoryMock struct {
	mock.Mock
	logger log.Logger
}

func NewRepoMock(logger log.Logger, mock mock.Mock) RepoSitoryMock {
	return RepoSitoryMock{Mock: mock, logger: logger}
}

func (repo *RepoSitoryMock) CreateUser(ctx context.Context, user entities.User) (string, error) {
	args := repo.Called(ctx, user)

	return (args.String(0)), args.Error(1)
}

func (repo *RepoSitoryMock) GetUser(ctx context.Context, userId string) (entities.User, error) {
	args := repo.Called(ctx, userId)
	id := args[0]

	return id.(entities.User), nil

}
