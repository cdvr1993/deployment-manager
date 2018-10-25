package services

import (
	"github.com/cdvr1993/deployment-manager/models"
)

type RoleServiceMethods struct {
	GetRole   func(string) models.Role
	ListRoles func() []*models.Role
}

type RoleService struct {
	methods RoleServiceMethods
}

func NewRoleServiceMock(m RoleServiceMethods) (s RoleService) {
	s.methods = m

	return
}

func (s RoleService) GetRole(name string) models.Role {
	if s.methods.GetRole != nil {
		return s.methods.GetRole(name)
	}

	return models.Role{}
}

func (s RoleService) ListRoles() []*models.Role {
	if s.methods.ListRoles != nil {
		return s.methods.ListRoles()
	}

	return nil
}
