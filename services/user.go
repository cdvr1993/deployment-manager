package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/cdvr1993/deployment-manager/models"
)

type IUserService interface {
	AddUser(*models.User)
	GetAll() []*models.User
	GetUser(int64) models.User
	GetUserByEmail(string) models.User
	DeleteUser(int64)
	UpdateUser(models.User)
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

func (s UserService) GetAll() (users []*models.User) {
	s.ormService.NewOrm().QueryTable(new(models.User)).All(&users)

	return
}

func (s UserService) GetUser(id int64) (u models.User) {
	u.Id = id

	if err := s.ormService.NewOrm().Read(&u); err == orm.ErrNoRows {
		panic(ErrorUserIdNotFound(id))
	}

	return
}

func (s UserService) GetUserByEmail(e string) (u models.User) {
	u.Email = e

	if err := s.ormService.NewOrm().Read(&u, "email"); err == orm.ErrNoRows {
		panic(ErrorUserNotFound(e))
	}

	return
}

func (s UserService) DeleteUser(uid int64) {
	user := s.GetUser(uid)

	if _, err := s.ormService.NewOrm().Delete(&user); err != nil {
		panic(err)
	}
}

func (s UserService) UpdateUser(u models.User) {
	if u.Email != "" {
		panic(ErrorUserCanNotEditEmail(u.Id, u.Email))
	}

	dbUser := s.GetUser(u.Id)

	if u.Name != "" {
		dbUser.Name = u.Name

		if _, err := s.ormService.NewOrm().Update(&dbUser); err != nil {
			panic(err)
		}
	}
}
