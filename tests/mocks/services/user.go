package services

import (
	"github.com/cdvr1993/deployment-manager/models"
)

type UserServiceMethods struct {
	AddUser func(*models.User)
	GetUser func(string) models.User
}
type UserService struct {
	methods UserServiceMethods
}

func NewUserServiceMock(m UserServiceMethods) (s UserService) {
	s.methods.AddUser = m.AddUser
	s.methods.GetUser = m.GetUser

	return
}

func (s UserService) AddUser(u *models.User) {
	s.methods.AddUser(u)
}

func (s UserService) GetUser(e string) models.User {
	return s.methods.GetUser(e)
}
