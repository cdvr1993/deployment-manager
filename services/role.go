package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/cdvr1993/deployment-manager/models"
)

type IRoleService interface {
	GetRole(string) *models.Role
	ListRoles() ([]*models.Role, error)
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

func (s RoleService) GetRole(name string) (role *models.Role) {
	o := s.ormService.NewOrm()

	role.Name = name

	if err := o.Read(role, "name"); err == orm.ErrNoRows {
		panic(ErrorRoleNotFound(name))
	}

	return
}

func (s RoleService) ListRoles() ([]*models.Role, error) {
	qs := s.ormService.NewOrm().QueryTable(new(models.Role))

	var roles []*models.Role
	if _, err := qs.All(&roles); err != nil && err != orm.ErrNoRows {
		return nil, err
	}

	return roles, nil
}
