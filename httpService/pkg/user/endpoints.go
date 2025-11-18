package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
	errs "github.com/timoteoBone/project-microservice/httpService/pkg/errors"
)

type Service interface {
	CreateUser(ctx context.Context, rq entities.CreateUserRequest) (entities.CreateUserResponse, error)
	GetUser(ctx context.Context, rq entities.GetUserRequest) (entities.GetUserResponse, error)
	Authenticate(ctx context.Context, rq entities.AuthenticateRequest) (entities.AuthenticateResponse, error)
}

type Endpoints struct {
	CreateUs     endpoint.Endpoint
	GetUs        endpoint.Endpoint
	Authenticate endpoint.Endpoint
}

func MakeEndpoints(s Service) *Endpoints {

	return &Endpoints{
		CreateUs:     MakeCreateUserEndpoint(s),
		GetUs:        MakeGetUserEndpoint(s),
		Authenticate: MakeAuthenticateEndpoint(s),
	}
}

func MakeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, rq interface{}) (interface{}, error) {
		request, valid := rq.(entities.CreateUserRequest)

		if !valid {
			return nil, errs.ErrInvalidDataForm
		}

		res, err := s.CreateUser(ctx, request)
		if err != nil {
			return nil, err
		}

		return res, nil

	}
}

func MakeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, rq interface{}) (interface{}, error) {
		request, valid := rq.(entities.GetUserRequest)
		if !valid {
			return nil, errs.ErrInvalidDataForm
		}

		res, err := s.GetUser(ctx, request)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func MakeAuthenticateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, rq interface{}) (interface{}, error) {
		request, valid := rq.(entities.AuthenticateRequest)
		if !valid {
			return nil, errs.ErrInvalidDataForm
		}

		res, err := s.Authenticate(ctx, request)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}
