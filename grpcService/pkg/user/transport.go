package user

import (
	"context"

	gr "github.com/go-kit/kit/transport/grpc"

	"github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
	customErr "github.com/timoteoBone/project-microservice/grpcService/pkg/errors"
	proto "github.com/timoteoBone/project-microservice/grpcService/pkg/pb"
)

type gRPCSv struct {
	createUs gr.Handler
	getUs    gr.Handler
	authUs   gr.Handler
	proto.UnimplementedUserServiceServer
}

func NewGrpcServer(end Endpoints) proto.UserServiceServer {

	return &gRPCSv{
		createUs: gr.NewServer(
			end.CreateUser,
			decodeCreateUserRequest,
			encodeCreateUserResponse,
		),

		getUs: gr.NewServer(
			end.GetUser,
			decodeGetUserRequest,
			encodeGetUserResponse,
		),

		authUs: gr.NewServer(
			end.AuthenticateUser,
			decodeAuthenticateRequest,
			encodeAuthenticateResponse,
		),
	}
}

func (g *gRPCSv) CreateUser(ctx context.Context, rq *proto.CreateUserRequest) (rs *proto.CreateUserResponse, err error) {
	_, resp, err := g.createUs.ServeGRPC(ctx, rq)

	if err != nil {
		status := customErr.CustomToGrpc(err)
		resp := proto.CreateUserResponse{Status: status}
		return &resp, err
	}

	return resp.(*proto.CreateUserResponse), nil
}

func (g *gRPCSv) GetUser(ctx context.Context, rq *proto.GetUserRequest) (rs *proto.GetUserResponse, err error) {
	_, resp, err := g.getUs.ServeGRPC(ctx, rq)

	if err != nil {
		return nil, err
	}

	return resp.(*proto.GetUserResponse), nil
}

func (g *gRPCSv) Authenticate(ctx context.Context, rq *proto.AuthenticateRequest) (*proto.AuthenticateResponse, error) {
	_, resp, err := g.authUs.ServeGRPC(ctx, rq)

	if err != nil {
		status := customErr.CustomToGrpc(err)
		resp := &proto.AuthenticateResponse{Status: status}
		return resp, err
	}

	return resp.(*proto.AuthenticateResponse), nil
}

func decodeCreateUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := request.(*proto.CreateUserRequest)

	if !err {
		return nil, customErr.NewGrpcError()
	}

	return entities.CreateUserRequest{
		Name: res.Name,
		Age:  res.Age,
		Pass: res.Pass,
	}, nil

}

func encodeCreateUserResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(entities.CreateUserResponse)
	status := &proto.Status{Message: res.Status.Message, Code: res.Status.Code}
	protoResp := &proto.CreateUserResponse{User_Id: res.UserId, Status: status}
	return protoResp, nil
}

func decodeGetUserRequest(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := request.(*proto.GetUserRequest)

	if !err {
		return nil, customErr.NewGrpcError()
	}

	return entities.GetUserRequest{
		UserID: res.User_Id,
	}, nil

}

func encodeGetUserResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(entities.GetUserResponse)
	protoResp := &proto.GetUserResponse{Id: res.Id, Name: res.Name, Age: res.Age}
	return protoResp, nil
}

func decodeAuthenticateRequest(ctx context.Context, rq interface{}) (interface{}, error) {
	req, err := rq.(*proto.AuthenticateRequest)

	if !err {
		return nil, customErr.NewGrpcError()
	}

	return entities.AuthenticateRequest{
		Email: req.Email,
		Pass:  req.Pass,
	}, nil

}

func encodeAuthenticateResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(entities.AuthenticateResponse)

	status := proto.Status{Message: resp.Status.Message, Code: resp.Status.Code}

	protoResp := *&proto.AuthenticateResponse{Status: &status}

	return protoResp, nil

}
