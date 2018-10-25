package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/cdvr1993/deployment-manager/models"
)

type IRoleService interface {
	GetRole(string) models.Role
	ListRoles() []*models.Role
}

type RoleService struct {
	ormService IOrmService
}

var (
	roleService = RoleService{
		ormService: NewOrmService(),
	}
)

func NewRoleService() *RoleService {
	return &roleService
}

func (s RoleService) GetRole(name string) (role models.Role) {
	o := s.ormService.NewOrm()

	role.Name = name

	if err := o.Read(&role, "name"); err == orm.ErrNoRows {
		panic(ErrorRoleNotFound(name))
	}

	return
}

func (s RoleService) ListRoles() (roles []*models.Role) {
	o := s.ormService.NewOrm()

	if _, err := o.QueryTable(new(models.Role)).All(&roles); err != nil {
		panic(err)
	}

	return
}
