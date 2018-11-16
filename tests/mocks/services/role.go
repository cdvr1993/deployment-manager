package services

import (
	"github.com/cdvr1993/deployment-manager/models"
)

type RoleServiceMethods struct {
	GetRole   func(string) (*models.Role, error)
	ListRoles func() ([]*models.Role, error)
}

type RoleService struct {
	methods RoleServiceMethods
}

func NewRoleServiceMock(m RoleServiceMethods) (s RoleService) {
	s.methods = m

	return
}

func (s RoleService) GetRole(name string) (*models.Role, error) {
	if s.methods.GetRole != nil {
		return s.methods.GetRole(name)
	}

	panic(ErrNotImplemented)
}

func (s RoleService) ListRoles() ([]*models.Role, error) {
	if s.methods.ListRoles != nil {
		return s.methods.ListRoles()
	}

	panic(ErrNotImplemented)
}
