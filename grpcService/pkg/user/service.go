package user

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"

	entities "github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
	mapper "github.com/timoteoBone/project-microservice/grpcService/pkg/mapper"
)

type Repository interface {
	GetUser(ctx context.Context, userId string) (entities.User, error)
	CreateUser(ctx context.Context, user entities.User) (string, error)
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
		response.Status = status
		return response, err
	}
	fmt.Println(genId)

	status.Message = "created successfully"
	response.Status = status
	response.UserId = genId

	return response, nil
}

func (s *service) GetUser(ctx context.Context, user entities.GetUserRequest) (entities.GetUserResponse, error) {
	s.Logger.Log(s.Logger, "request", "get user", "recevied")

	res, err := s.Repo.GetUser(ctx, user.UserID)
	if err != nil {
		return entities.GetUserResponse{}, err
	}

	response := entities.GetUserResponse{
		Name: res.Name,
		Id:   user.UserID,
		Age:  res.Age,
	}

	return response, nil
}
