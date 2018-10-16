package services

import (
	"fmt"
)

type IServiceError interface {
	ID() string
	Error() string
	Status() int
}

type ServiceError struct {
	id     string
	msg    string
	status int
}

func (s ServiceError) ID() string {
	return s.id
}

func (s ServiceError) Error() string {
	return s.msg
}

func (s ServiceError) Status() int {
	return s.status
}

func ErrorUserEmailExists(email string) IServiceError {
	return ServiceError{
		id:     "EmailExists",
		msg:    fmt.Sprintf("User email '%s' is currently in use", email),
		status: 400,
	}
}
