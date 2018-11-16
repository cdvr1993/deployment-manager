package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/cdvr1993/deployment-manager/models"
)

type IUserService interface {
	AddUser(*models.User) error
	GetAll() ([]*models.User, error)
	GetUser(int64) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	DeleteUser(int64) error
	UpdateUser(*models.User) error
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

func (s UserService) AddUser(u *models.User) error {
	o := s.ormService.NewOrm()

	created, id, err := o.ReadOrCreate(u, "email")

	if err != nil {
		return err
	}

	if !created {
		return ErrorUserEmailExists(u.Email)
	}

	u.Id = id

	return nil
}

func (s UserService) GetAll() ([]*models.User, error) {
	qs := s.ormService.NewOrm().QueryTable(new(models.User))

	users := make([]*models.User, 0)
	if _, err := qs.All(&users); err != nil && err != orm.ErrNoRows {
		return nil, err
	}

	return users, nil
}

func (s UserService) GetUser(id int64) (*models.User, error) {
	u := models.User{Id: id}

	if err := s.ormService.NewOrm().Read(&u); err == orm.ErrNoRows {
		return nil, ErrorUserIdNotFound(id)
	}

	return &u, nil
}

func (s UserService) GetUserByEmail(e string) (*models.User, error) {
	u := models.User{Email: e}

	if err := s.ormService.NewOrm().Read(&u, "email"); err == orm.ErrNoRows {
		return nil, ErrorUserNotFound(e)
	}

	return &u, nil
}

func (s UserService) DeleteUser(uid int64) error {
	user, err := s.GetUser(uid)

	if err != nil {
		return err
	}

	if _, err := s.ormService.NewOrm().Delete(&user); err != nil {
		return err
	}

	return nil
}

func (s UserService) UpdateUser(u *models.User) error {
	if u.Email != "" {
		return ErrorUserCanNotEditEmail(u.Id, u.Email)
	}

	dbUser, err := s.GetUser(u.Id)

	if err != nil {
		return err
	}

	if u.Name == "" {
		return nil
	}

	dbUser.Name = u.Name

	if _, err := s.ormService.NewOrm().Update(&dbUser); err != nil {
		return err
	}

	return nil
}
