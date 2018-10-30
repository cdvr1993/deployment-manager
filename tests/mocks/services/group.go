package services

import (
	"github.com/cdvr1993/deployment-manager/models"
	svcs "github.com/cdvr1993/deployment-manager/services"
)

type GroupServiceMethods struct {
	AddMember      func(int64, int64, string)
	CreateGroup    func(*models.Group, string)
	DeleteGroup    func(int64)
	GetAllGroups   func(svcs.GetAllGroupsOptions) []*models.Group
	GetGroup       func(int64, *svcs.GetGroupOptions) models.Group
	GetGroupByName func(string) models.Group
	IsAllowed      func(models.Group, models.User, bool) bool
	RemoveMember   func(int64, int64)
	UpdateGroup    func(models.Group)
}

type GroupService struct {
	methods GroupServiceMethods
}

func NewGroupServiceMock(m GroupServiceMethods) (s GroupService) {
	s.methods = m

	return
}

func (s GroupService) AddMember(a, b int64, c string) {
	if s.methods.AddMember != nil {
		s.methods.AddMember(a, b, c)
	}
}

func (s GroupService) CreateGroup(g *models.Group, e string) {
	if s.methods.CreateGroup != nil {
		s.methods.CreateGroup(g, e)
	}
}

func (s GroupService) DeleteGroup(i int64) {
	if s.methods.DeleteGroup != nil {
		s.methods.DeleteGroup(i)
	}
}

func (s GroupService) GetAllGroups(opts svcs.GetAllGroupsOptions) []*models.Group {
	if s.methods.GetAllGroups != nil {
		return s.methods.GetAllGroups(opts)
	}

	return nil
}

func (s GroupService) GetGroup(id int64, opts *svcs.GetGroupOptions) models.Group {
	if s.methods.GetGroup != nil {
		return s.methods.GetGroup(id, opts)
	}

	return models.Group{}
}

func (s GroupService) GetGroupByName(n string) models.Group {
	if s.methods.GetGroupByName != nil {
		return s.methods.GetGroupByName(n)
	}

	return models.Group{}
}

func (s GroupService) IsAllowed(g models.Group, u models.User, ro bool) bool {
	if s.methods.IsAllowed != nil {
		return s.methods.IsAllowed(g, u, ro)
	}

	return false
}

func (s GroupService) RemoveMember(gid int64, uid int64) {
	if s.methods.RemoveMember != nil {
		s.methods.RemoveMember(gid, uid)
	}
}

func (s GroupService) UpdateGroup(g models.Group) {
	if s.methods.UpdateGroup != nil {
		s.methods.UpdateGroup(g)
	}
}
