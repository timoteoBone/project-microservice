package errors

import (
	"errors"
	"fmt"
	"net/http"

	pb "github.com/timoteoBone/project-microservice/grpcService/pkg/pb"
)

type UserNotFoundErr struct {
	err error
}

type FieldsMissingErr struct {
	err error
}

type GrpcErr struct {
	err error
}

type DataBaseErr struct {
	err error
}

func (err FieldsMissingErr) Error() string {
	return fmt.Sprint(err.err)
}

func (err UserNotFoundErr) Error() string {
	return fmt.Sprint(err.err)
}

func (err GrpcErr) Error() string {
	return fmt.Sprint(err.err)
}

func (err DataBaseErr) Error() string {
	return fmt.Sprint(err.err)
}

func NewFieldsMissing() FieldsMissingErr {
	return FieldsMissingErr{err: errors.New("all fields are required")}
}

func NewUserNotFound() UserNotFoundErr {
	return UserNotFoundErr{err: errors.New("user not found")}
}

func NewGrpcError() GrpcErr {
	return GrpcErr{err: errors.New("uknown grpc error")}
}

func NewDataBaseError() DataBaseErr {
	return DataBaseErr{err: errors.New("unknown database error")}
}

func CustomToGrpc(err error) *pb.Status {
	var status pb.Status
	switch err.(type) {
	case FieldsMissingErr:
		status.Code = 3
		status.Message = err.Error()
	case UserNotFoundErr:
		status.Code = 5
		status.Message = err.Error()
	}
	return &status
}

func GrpcToCustom(code int32) error {
	var err error
	switch code {
	case 3:
		err = NewFieldsMissing()

	case 5:
		err = NewUserNotFound()
	default:
		err = NewGrpcError()
	}
	return err
}

func CustomToHttp(err error) int {
	switch err.(type) {
	case UserNotFoundErr:
		return http.StatusNotFound
	case FieldsMissingErr:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}

}
