package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/cdvr1993/deployment-manager/models"
)

type IUserService interface {
	AddUser(*models.User)
	GetUser(int64) models.User
	GetUserByEmail(string) models.User
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
	}
}

func (s UserService) GetUser(id int64) (u models.User) {
	o := s.ormService.NewOrm()

	u.Id = id

	if err := o.Read(&u); err == orm.ErrNoRows {
		panic(ErrorUserIdNotFound(id))
	}

	return
}

func (s UserService) GetUserByEmail(e string) (u models.User) {
	o := s.ormService.NewOrm()

	u.Email = e

	if err := o.Read(&u, "email"); err == orm.ErrNoRows {
		panic(ErrorUserNotFound(e))
	}

	return
}
