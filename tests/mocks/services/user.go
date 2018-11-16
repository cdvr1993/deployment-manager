package services

import (
	"github.com/cdvr1993/deployment-manager/models"
)

type UserServiceMethods struct {
	AddUser        func(*models.User) error
	GetAll         func() ([]*models.User, error)
	GetUser        func(int64) (*models.User, error)
	GetUserByEmail func(string) (*models.User, error)
	DeleteUser     func(int64) error
	UpdateUser     func(*models.User) error
}
type UserService struct {
	methods UserServiceMethods
}

func NewUserServiceMock(m UserServiceMethods) (s UserService) {
	s.methods = m

	return
}

func (s UserService) AddUser(u *models.User) error {
	if s.methods.AddUser != nil {
		return s.methods.AddUser(u)
	}

	panic(ErrNotImplemented)
}

func (s UserService) GetAll() ([]*models.User, error) {
	if s.methods.GetAll != nil {
		return s.methods.GetAll()
	}

	panic(ErrNotImplemented)
}

func (s UserService) GetUser(id int64) (*models.User, error) {
	if s.methods.GetUser != nil {
		return s.methods.GetUser(id)
	}

	panic(ErrNotImplemented)
}

func (s UserService) GetUserByEmail(e string) (*models.User, error) {
	if s.methods.GetUserByEmail != nil {
		return s.methods.GetUserByEmail(e)
	}

	panic(ErrNotImplemented)
}

func (s UserService) DeleteUser(id int64) error {
	if s.methods.DeleteUser != nil {
		return s.methods.DeleteUser(id)
	}

	panic(ErrNotImplemented)
}

func (s UserService) UpdateUser(u *models.User) error {
	if s.methods.UpdateUser != nil {
		return s.methods.UpdateUser(u)
	}

	panic(ErrNotImplemented)
}
