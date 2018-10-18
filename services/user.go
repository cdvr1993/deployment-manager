package services

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/cdvr1993/deployment-manager/models"
)

type IUserService interface {
	AddUser(u *models.User)
	GetUser(e string) models.User
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

func (s UserService) GetUser(e string) (user models.User) {
	o := s.ormService.NewOrm()

	user.Email = e

	if err := o.Read(&user, "email"); err == orm.ErrNoRows {
		panic(ErrorUserNotFound(e))
	} else if err != nil {
		beego.Error(err)
		panic(err)
	}

	return
}
