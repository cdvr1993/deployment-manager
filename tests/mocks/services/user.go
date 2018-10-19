package services

import (
	"github.com/cdvr1993/deployment-manager/models"
)

type UserServiceMethods struct {
	AddUser        func(*models.User)
	GetUser        func(int64) models.User
	GetUserByEmail func(string) models.User
}
type UserService struct {
	methods UserServiceMethods
}

func NewUserServiceMock(m UserServiceMethods) (s UserService) {
	s.methods.AddUser = m.AddUser
	s.methods.GetUserByEmail = m.GetUserByEmail

	return
}

func (s UserService) AddUser(u *models.User) {
	s.methods.AddUser(u)
}

func (s UserService) GetUser(id int64) models.User {
	return s.methods.GetUser(id)
}

func (s UserService) GetUserByEmail(e string) models.User {
	return s.methods.GetUserByEmail(e)
}
