package user_test

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/errors"
	err "github.com/timoteoBone/project-microservice/grpcService/pkg/errors"
	service "github.com/timoteoBone/project-microservice/grpcService/pkg/user"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/utils"
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
	repo := utils.NewRepoMock(logger, mock.Mock{})

	srvc := service.NewService(logger, &repo)

	assert.False(t, srvc == nil)
}

func TestServiceCreateUser(t *testing.T) {
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

	user := entities.User{
		Name: "Timo",
		Pass: "123",
		Age:  19,
	}

	userId := utils.GenerateId()

	correctCreateUserRequest := entities.CreateUserRequest{
		Name: user.Name,
		Pass: user.Pass,
		Age:  user.Age,
	}

	Succesfullresponse := entities.CreateUserResponse{
		Status: entities.Status{
			Message: "created successfully",
		}, UserId: userId,
	}

	testsCases := []struct {
		Name            string
		Identifier      string
		User            entities.User
		Request         entities.CreateUserRequest
		RepoError       error
		RepoId          string
		ServiceResponse entities.CreateUserResponse
		Error           error
	}{
		{
			Name:            "Create User Valid Case",
			Identifier:      "CreateUser",
			User:            user,
			Request:         correctCreateUserRequest,
			RepoError:       nil,
			RepoId:          userId,
			ServiceResponse: Succesfullresponse,
			Error:           nil,
		},

		{
			Name:       "Create User Empty Fields",
			Identifier: "CreateUser",
			User:       entities.User{Pass: "sdsd"},
			Request:    entities.CreateUserRequest{Pass: "sdsd"},
			RepoError:  err.ErrAllFieldsRequired,
			RepoId:     userId,
			ServiceResponse: entities.CreateUserResponse{
				Status: entities.Status{
					Message: "Unable to create user",
				},
			},
			Error: errors.ErrAllFieldsRequired,
		},
	}

	repo := utils.NewRepoMock(logger, mock.Mock{})
	srvc := service.NewService(logger, &repo)

	for _, tc := range testsCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()
			repo.On(tc.Identifier, ctx, tc.User).Return((tc.RepoId), tc.RepoError)

			res, err := srvc.CreateUser(ctx, tc.Request)
			assert.ErrorIs(t, err, tc.RepoError)
			assert.Equal(t, tc.ServiceResponse, res)

		})

	}

}

func TestServiceGetUser(t *testing.T) {

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

	user := entities.User{
		Name: "Timo",
		Pass: "123",
		Age:  19,
	}

	userId := utils.GenerateId()

	correctGetUserRequest := entities.GetUserRequest{
		UserID: userId,
	}

	correctGetUserResponse := entities.GetUserResponse{
		Name: user.Name,
		Id:   userId,
		Age:  user.Age,
	}

	repo := new(utils.RepoSitoryMock)
	srvc := service.NewService(logger, repo)

	ctx := context.Background()

	testCases := []struct {
		Name                    string
		Identifier              string
		UserId                  string
		Request                 entities.GetUserRequest
		RepoReturn              entities.User
		RepoError               error
		ServiceExpectedResponse entities.GetUserResponse
		Error                   error
	}{
		{
			Name:                    "Get User Valid Request",
			Identifier:              "GetUser",
			UserId:                  userId,
			Request:                 correctGetUserRequest,
			RepoReturn:              user,
			RepoError:               nil,
			ServiceExpectedResponse: correctGetUserResponse,
			Error:                   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			repo.Mock.On(tc.Identifier, ctx, tc.UserId).Return(tc.RepoReturn, tc.RepoError)

			res, err := srvc.GetUser(ctx, tc.Request)
			assert.Equal(t, tc.ServiceExpectedResponse, res)
			assert.ErrorIs(t, err, tc.Error)

		})
	}

}

func TestAuthenticateUser(t *testing.T) {

}
