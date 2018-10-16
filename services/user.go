package services

import (
	"github.com/cdvr1993/deployment-manager/models"
)

type IUserService interface {
	AddUser(u *models.User)
}

type UserService struct {
	ormService IOrmService
}

var (
	userService = UserService{
		ormService: NewOrmService(),
	}
)

func NewUserService() *UserService {
	return &userService
}

func (s UserService) AddUser(u *models.User) {
	o := s.ormService.NewOrm()

	if created, id, err := o.ReadOrCreate(u, "email"); err == nil {
		if created {
			u.Id = id
		} else {
			panic(ErrorUserEmailExists(u.Email))
		}
	} else {
		panic(err)
	}
}
