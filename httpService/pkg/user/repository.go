package user

import (
	"context"
	"fmt"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/entities"

	proto "github.com/timoteoBone/project-microservice/grpcService/pkg/pb"
	"google.golang.org/grpc"
)

type grpcClient struct {
	server *grpc.ClientConn
	logger log.Logger
}

func NewgRPClient(log log.Logger, sv *grpc.ClientConn) *grpcClient {
	return &grpcClient{sv, log}
}

func (repo *grpcClient) CreateUser(ctx context.Context, rq entities.CreateUserRequest) (entities.CreateUserResponse, error) {
	logger := log.With(repo.logger, "create  user", "recevied")

	client := proto.NewUserServiceClient(repo.server)

	protoReq := proto.CreateUserRequest{
		Name: rq.Name,
		Age:  rq.Age,
		Pass: rq.Pass,
	}

	resp, err := client.CreateUser(ctx, &protoReq)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return entities.CreateUserResponse{}, err
	}

	status := entities.Status{
		Message: resp.Status.Message,
	}

	res := entities.CreateUserResponse{
		Status: status,
		UserId: resp.User_Id,
	}

	return res, nil

}

func (repo *grpcClient) GetUser(ctx context.Context, rq entities.GetUserRequest) (entities.GetUserResponse, error) {
	logger := log.With(repo.logger, "get user request", "received")

	client := proto.NewUserServiceClient(repo.server)

	protoReq := proto.GetUserRequest{User_Id: rq.UserID}

	protoRes, err := client.GetUser(ctx, &protoReq)
	fmt.Println("hasta 1")
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return entities.GetUserResponse{}, err
	}

	fmt.Println("hasta 2")

	resp := entities.GetUserResponse{
		Name: protoRes.Name,
		Id:   protoReq.User_Id,
		Age:  protoRes.Age,
	}

	return resp, nil
}
