package errors

import (
	"errors"
	"fmt"
)

type GrpcError struct {
	err error
}

var (
	ErrUserNotFound      error = NewGrpcErr("user not found")
	ErrInvalidData       error = NewGrpcErr("invalid data type")
	ErrAllFieldsRequired error = NewGrpcErr("all fields are required")
)

func (err GrpcError) Error() string {
	return fmt.Sprint(err.err)
}

func NewGrpcErr(message string) GrpcError {
	return GrpcError{err: errors.New(message)}
}
