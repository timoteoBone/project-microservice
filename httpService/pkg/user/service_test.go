package user_test

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
	errors "github.com/timoteoBone/project-microservice/httpService/pkg/errors"
	"github.com/timoteoBone/project-microservice/httpService/pkg/user"
	util "github.com/timoteoBone/project-microservice/httpService/pkg/utils"
)

func TestNewService(t *testing.T) {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	repo := util.NewRepositoryMock()

	srvc := user.NewService(&repo, logger)

	assert.NotNil(t, srvc)

}

func TestCreateUser(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	repo := util.NewRepositoryMock()

	srvc := user.NewService(&repo, logger)

	var (
		correctCreateRequest entities.CreateUserRequest = entities.CreateUserRequest{
			Name: "Timo",
			Age:  19,
			Pass: "123",
		}

		correctCreateResponse entities.CreateUserResponse = entities.CreateUserResponse{
			Status: entities.Status{Message: "created successfully"},
		}
	)

	testCases := []struct {
		Name            string
		Identifier      string
		ServiceRequest  entities.CreateUserRequest
		ServiceResponse entities.CreateUserResponse
		RepoRequest     entities.CreateUserRequest
		RepoResponse    entities.CreateUserResponse
		RepoError       error
		ServiceError    error
	}{
		{

			Name:            "Create user valid case",
			Identifier:      "CreateUser",
			ServiceRequest:  correctCreateRequest,
			ServiceResponse: correctCreateResponse,
			RepoRequest:     correctCreateRequest,
			RepoResponse:    correctCreateResponse,
			RepoError:       nil,
			ServiceError:    nil,
		},
		{
			Name:            "Create User Unvalid case",
			Identifier:      "CreateUser",
			ServiceRequest:  entities.CreateUserRequest{},
			ServiceResponse: entities.CreateUserResponse{},
			RepoRequest:     entities.CreateUserRequest{},
			RepoResponse:    entities.CreateUserResponse{},
			RepoError:       errors.ErrInvalidDataForm,
			ServiceError:    errors.ErrInvalidDataForm,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()
			repo.On(tc.Identifier, ctx, tc.RepoRequest).Return(tc.RepoResponse, tc.RepoError)

			res, err := srvc.CreateUser(ctx, tc.ServiceRequest)
			assert.ErrorIs(t, err, tc.ServiceError)
			assert.Equal(t, tc.ServiceResponse, res)
		})
	}

}

func TestGetUserCreate(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	repo := util.NewRepositoryMock()

	srvc := user.NewService(&repo, logger)

	var (
		correctGetUserRequest entities.GetUserRequest = entities.GetUserRequest{
			UserID: "2abc-323kol",
		}

		correctGetUserResponse entities.GetUserResponse = entities.GetUserResponse{
			Name: "Timo",
			Age:  19,
		}
	)

	testCases := []struct {
		Name       string
		Identifier string
		Request    entities.GetUserRequest
		Response   entities.GetUserResponse
		Error      error
	}{
		{
			Name:       "Get User Valid Case",
			Identifier: "GetUser",
			Request:    correctGetUserRequest,
			Response:   correctGetUserResponse,
			Error:      nil,
		},

		{
			Name:       "Get User Not Found Case",
			Identifier: "GetUser",
			Request:    correctGetUserRequest,
			Response:   entities.GetUserResponse{},
			Error:      errors.ErrUserNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()
			repo.On(tc.Identifier, ctx, tc.Request).Return(tc.Response, tc.Error)

			res, err := srvc.GetUser(ctx, tc.Request)
			assert.Equal(t, tc.Response, res)
			assert.ErrorIs(t, err, tc.Error)

		})
	}

}
