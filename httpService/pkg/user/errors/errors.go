package error

import (
	"errors"
	"fmt"
)

type DomainError struct {
	err error
}

var (
	ErrInvalidDataForm error = NewError("invalid data type or form")
	ErrUserNotFound    error = NewError("user not found")
)

func NewError(message string) DomainError {
	return DomainError{err: errors.New(message)}
}

func (err DomainError) Error() string {
	return fmt.Sprint(err.err)
}
