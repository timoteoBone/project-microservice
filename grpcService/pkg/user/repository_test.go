package user_test

import (
	"context"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/user"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/utils"
)

func TestNewRepo(t *testing.T) {
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

	db, _ := utils.NewMock(logger)

	repo := user.NewSQL(db, logger)

	assert.NotNil(t, repo)
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

	db, mock := utils.NewMock(logger)
	defer db.Close()

	userId := utils.GenerateId()

	repo := user.NewSQL(db, logger)

	testCases := []struct {
		Name              string
		Identifier        string
		User              entities.User
		UserID            string
		Query             string
		ExpectedRespError error
	}{

		{
			Name:       "Create User Valid Case",
			Identifier: "CreateUser",
			User: entities.User{
				Name:  "Timo",
				Age:   19,
				Pass:  "1234",
				Email: "timoteo@globant.com",
			},
			UserID:            userId,
			Query:             "INSERT INTO USER (first_name, id, age, pass) VALUES (?,?,?,?)",
			ExpectedRespError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			ctx := context.Background()

			prep := mock.ExpectPrepare(tc.Query)
			prep.ExpectExec().WithArgs(tc.User.Name, userId, tc.User.Pass, tc.User.Age, tc.User.Email).WillReturnResult(sqlmock.NewResult(int64(1), int64(1)))

			_, err := repo.CreateUser(ctx, tc.User, userId)
			assert.NoError(t, err)
		})
	}
}

func TestGetUser(t *testing.T) {
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

	userMock := entities.User{
		Name:  "Timoteo",
		Pass:  "sdsfodemf",
		Age:   19,
		Email: "timoteo@globant.com",
	}

	db, mock := utils.NewMock(logger)
	defer db.Close()

	userId := utils.GenerateId()

	repo := user.NewSQL(db, logger)

	testCases := []struct {
		Name           string
		UserID         string
		buildMock      func(mock sqlmock.Sqlmock, userId string)
		assertResponse func(t *testing.T, response entities.User, err error)
	}{
		{
			Name:   "Get Existing User",
			UserID: userId,
			buildMock: func(mock sqlmock.Sqlmock, userId string) {
				res := sqlmock.NewRows([]string{"first_name", "age", "email"}).AddRow(userMock.Name, userMock.Age, userMock.Email)
				mock.ExpectPrepare(utils.GetUser)
				mock.ExpectQuery(utils.GetUser).WithArgs(userId).WillReturnRows(res)
			},
			assertResponse: func(t *testing.T, resp entities.User, err error) {
				assert.Nil(t, err)
				assert.Equal(t, userMock.Email, resp.Email)
			},
		},

		{
			Name:   "Get non existing user",
			UserID: userId,
			buildMock: func(mock sqlmock.Sqlmock, userId string) {
				res := sqlmock.NewRows([]string{"first_name", "age", "email"})
				mock.ExpectPrepare(utils.GetUser)
				mock.ExpectQuery(utils.GetUser).WithArgs(userId).WillReturnRows(res)
			},
			assertResponse: func(t *testing.T, resp entities.User, err error) {
				assert.Error(t, err)
				assert.Empty(t, resp)
			},
		},
	}

	for _, tc := range testCases {

		t.Run(tc.Name, func(t *testing.T) {
			ctx := context.Background()
			tc.buildMock(mock, userId)

			res, err := repo.GetUser(ctx, tc.UserID)
			tc.assertResponse(t, res, err)
		})
	}

}
