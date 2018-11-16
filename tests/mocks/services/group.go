package services

import (
	"github.com/cdvr1993/deployment-manager/models"
	svcs "github.com/cdvr1993/deployment-manager/services"
)

type GroupServiceMethods struct {
	AddMember      func(int64, int64, string) error
	CreateGroup    func(*models.Group, string) error
	DeleteGroup    func(int64) error
	GetAllGroups   func(svcs.GetAllGroupsOptions) ([]*models.Group, error)
	GetGroup       func(int64, *svcs.GetGroupOptions) (*models.Group, error)
	GetGroupByName func(string) (*models.Group, error)
	IsAllowed      func(*models.Group, *models.User, bool) bool
	RemoveMember   func(int64, int64) error
	UpdateGroup    func(*models.Group) error
}

type GroupService struct {
	methods GroupServiceMethods
}

func NewGroupServiceMock(m GroupServiceMethods) (s GroupService) {
	s.methods = m

	return
}

func (s GroupService) AddMember(a, b int64, c string) error {
	if s.methods.AddMember != nil {
		return s.methods.AddMember(a, b, c)
	}

	panic(ErrNotImplemented)
}

func (s GroupService) CreateGroup(g *models.Group, e string) error {
	if s.methods.CreateGroup != nil {
		return s.methods.CreateGroup(g, e)
	}

	panic(ErrNotImplemented)
}

func (s GroupService) DeleteGroup(i int64) error {
	if s.methods.DeleteGroup != nil {
		return s.methods.DeleteGroup(i)
	}

	panic(ErrNotImplemented)
}

func (s GroupService) GetAllGroups(opts svcs.GetAllGroupsOptions) ([]*models.Group, error) {
	if s.methods.GetAllGroups != nil {
		return s.methods.GetAllGroups(opts)
	}

	panic(ErrNotImplemented)
}

func (s GroupService) GetGroup(id int64, opts *svcs.GetGroupOptions) (*models.Group, error) {
	if s.methods.GetGroup != nil {
		return s.methods.GetGroup(id, opts)
	}

	panic(ErrNotImplemented)
}

func (s GroupService) GetGroupByName(n string) (*models.Group, error) {
	if s.methods.GetGroupByName != nil {
		return s.methods.GetGroupByName(n)
	}

	panic(ErrNotImplemented)
}

func (s GroupService) IsAllowed(g *models.Group, u *models.User, ro bool) bool {
	if s.methods.IsAllowed != nil {
		return s.methods.IsAllowed(g, u, ro)
	}

	panic(ErrNotImplemented)
}

func (s GroupService) RemoveMember(gid int64, uid int64) error {
	if s.methods.RemoveMember != nil {
		return s.methods.RemoveMember(gid, uid)
	}

	panic(ErrNotImplemented)
}

func (s GroupService) UpdateGroup(g *models.Group) error {
	if s.methods.UpdateGroup != nil {
		return s.methods.UpdateGroup(g)
	}

	panic(ErrNotImplemented)
}
