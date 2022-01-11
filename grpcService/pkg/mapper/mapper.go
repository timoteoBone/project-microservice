package mapper

import "github.com/timoteoBone/project-microservice/grpcService/pkg/entities"

func CreateUserRequestToUser(userReq entities.CreateUserRequest) entities.User {

	user := entities.User{
		userReq.Name,
		userReq.Pass,
		userReq.Age,
	}
	return user
}
