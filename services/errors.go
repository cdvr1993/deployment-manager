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

func ErrorNothingToUpdate(obj interface{}) IServiceError {
	return ServiceError{
		id:     "NothingToUpdate",
		msg:    fmt.Sprintf("There is nothing to update '(%x)'", obj),
		status: 400,
	}
}

func ErroGroupIdNotFound(id int64) IServiceError {
	return ServiceError{
		id:     "GroupIdNotFound",
		msg:    fmt.Sprintf("Group with id '%d' doesn't exist", id),
		status: 404,
	}
}

func ErroGroupNotFound(name string) IServiceError {
	return ServiceError{
		id:     "GroupNotFound",
		msg:    fmt.Sprintf("Group '%s' doesn't exist", name),
		status: 404,
	}
}

func ErrorGroupNameExists(name string) IServiceError {
	return ServiceError{
		id:     "GroupNameExists",
		msg:    fmt.Sprintf("Group '%s' already exists", name),
		status: 400,
	}
}

func ErrorRoleNotFound(role string) IServiceError {
	return ServiceError{
		id:     "RoleNotFound",
		msg:    fmt.Sprintf("The role '%s' doesn't exist", role),
		status: 404,
	}
}

func ErrorUserCanNotEditEmail(id int64, email string) IServiceError {
	return ServiceError{
		id:     "UserCanNotEditEmail",
		msg:    fmt.Sprintf("Can't change email '%s' for user with id '%d'", email, id),
		status: 400,
	}
}

func ErrorUserEmailExists(email string) IServiceError {
	return ServiceError{
		id:     "EmailExists",
		msg:    fmt.Sprintf("User email '%s' is currently in use", email),
		status: 400,
	}
}

func ErrorUserIdNotFound(id int64) IServiceError {
	return ServiceError{
		id:     "UserIdNotFound",
		msg:    fmt.Sprintf("User with id '%d' doesn't exist", id),
		status: 404,
	}
}

func ErrorUserNotFound(email string) IServiceError {
	return ServiceError{
		id:     "UserNotFound",
		msg:    fmt.Sprintf("User with email '%s' doesn't exist", email),
		status: 404,
	}
}
