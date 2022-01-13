package user

import (
	"context"

	"database/sql"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	entities "github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
	errors "github.com/timoteoBone/project-microservice/grpcService/pkg/errors"
	mapper "github.com/timoteoBone/project-microservice/grpcService/pkg/mapper"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/utils"
)

type Repository interface {
	GetUser(ctx context.Context, userId string) (entities.User, error)
	CreateUser(ctx context.Context, user entities.User) (string, error)
	AuthenticateUser(ctx context.Context, email string) (string, error)
}

type service struct {
	Repo   Repository
	Logger log.Logger
}

func NewService(l log.Logger, r Repository) *service {
	return &service{r, l}
}

func (s *service) CreateUser(ctx context.Context, userReq entities.CreateUserRequest) (entities.CreateUserResponse, error) {
	s.Logger.Log(s.Logger, "request", "create user", "recevied")

	response := entities.CreateUserResponse{}
	status := entities.Status{}

	user := mapper.CreateUserRequestToUser(userReq)
	genId, err := s.Repo.CreateUser(ctx, user)

	if err != nil {
		status.Message = "Unable to create user"
		status.Code = 10
		response.Status = status
		return response, errors.NewDataBaseError()
	}

	status.Message = "created successfully"
	response.Status = status
	response.UserId = genId

	return response, nil
}

func (s *service) GetUser(ctx context.Context, user entities.GetUserRequest) (entities.GetUserResponse, error) {
	s.Logger.Log(s.Logger, "request", "get user", "recevied")

	res, err := s.Repo.GetUser(ctx, user.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.GetUserResponse{}, errors.NewUserNotFound()
		}
		return entities.GetUserResponse{}, err
	}

	response := entities.GetUserResponse{
		Name: res.Name,
		Id:   user.UserID,
		Age:  res.Age,
	}

	return response, nil
}

func (s *service) AuthenticateUser(ctx context.Context, rq entities.AuthenticateRequest) (entities.AuthenticateResponse, error) {
	s.Logger.Log(s.Logger, "request", "authenticate user", "recevied")

	res, err := s.Repo.AuthenticateUser(ctx, rq.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			level.Error(s.Logger).Log(err, "sql error")
			return entities.AuthenticateResponse{}, errors.NewUserNotFound()
		}
		return entities.AuthenticateResponse{}, err
	}

	validation := utils.CheckPassword(rq.Pass, res)
	if validation != nil {
		return entities.AuthenticateResponse{}, errors.NewDeniedAuthentication()
	}

	resp := entities.AuthenticateResponse{Status: entities.Status{Message: "authenticated succesfully", Code: 0}}

	return resp, nil
}
