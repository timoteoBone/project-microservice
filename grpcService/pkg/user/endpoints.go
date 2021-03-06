package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
)

type Service interface {
	GetUser(ctx context.Context, userReq entities.GetUserRequest) (entities.GetUserResponse, error)
	CreateUser(ctx context.Context, userReq entities.CreateUserRequest) (entities.CreateUserResponse, error)
	AuthenticateUser(ctx context.Context, userReq entities.AuthenticateRequest) (entities.AuthenticateResponse, error)
}

type Endpoints struct {
	CreateUser       endpoint.Endpoint
	GetUser          endpoint.Endpoint
	AuthenticateUser endpoint.Endpoint
}

func MakeEndpoint(s Service) Endpoints {
	return Endpoints{
		CreateUser:       MakeCreateUserEndpoint(s),
		GetUser:          MakeGetUserEndpoint(s),
		AuthenticateUser: MakeAuthenticateUserEndpoint(s),
	}
}

func MakeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		req := request.(entities.CreateUserRequest)
		c, err := s.CreateUser(ctx, req)
		if err != nil {
			return nil, err
		}

		return c, nil

	}
}

func MakeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		req := request.(entities.GetUserRequest)
		c, err := s.GetUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return c, nil

	}
}

func MakeAuthenticateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entities.AuthenticateRequest)
		resp, err := s.AuthenticateUser(ctx, req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}
