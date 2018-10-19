package services

import (
	"github.com/cdvr1993/deployment-manager/models"
)

type UserServiceMethods struct {
	AddUser        func(*models.User)
	GetAll         func() []*models.User
	GetUser        func(int64) models.User
	GetUserByEmail func(string) models.User
	DeleteUser     func(int64)
	UpdateUser     func(models.User)
}
type UserService struct {
	methods UserServiceMethods
}

func NewUserServiceMock(m UserServiceMethods) (s UserService) {
	s.methods = m

	return
}

func (s UserService) AddUser(u *models.User) {
	s.methods.AddUser(u)
}

func (s UserService) GetAll() []*models.User {
	return s.methods.GetAll()
}

func (s UserService) GetUser(id int64) models.User {
	return s.methods.GetUser(id)
}

func (s UserService) GetUserByEmail(e string) models.User {
	return s.methods.GetUserByEmail(e)
}

func (s UserService) DeleteUser(id int64) {
	if s.methods.DeleteUser != nil {
		s.methods.DeleteUser(id)
	}
}

func (s UserService) UpdateUser(u models.User) {
	if s.methods.UpdateUser != nil {
		s.methods.UpdateUser(u)
	}
}
