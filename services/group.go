package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/cdvr1993/deployment-manager/models"
)

type IGroupService interface {
	AddMember(int64, int64)
	CreateGroup(*models.Group)
	DeleteGroup(int64)
	GetAllGroups() []*models.Group
	GetGroup(int64, *GetGroupOptions) models.Group
	GetGroupByName(string) models.Group
	RemoveMember(int64, int64)
	UpdateGroup(models.Group)
}

type GroupService struct {
	ormService  IOrmService
	userService IUserService
}

var (
	groupService = GroupService{
		ormService:  NewOrmService(),
		userService: NewUserService(),
	}
)

func NewGroupService() *GroupService {
	return &groupService
}

func (s GroupService) AddMember(gid, uid int64) {
	user := s.userService.GetUser(uid)
	group := s.GetGroup(gid, nil)

	memb := models.GroupMember{User: &user, Group: &group}

	s.ormService.NewOrm().ReadOrCreate(&memb, "user_id", "group_id")
}

func (s GroupService) CreateGroup(g *models.Group) {
	o := s.ormService.NewOrm()

	if created, id, err := o.ReadOrCreate(g, "name"); err == nil {
		if created {
			g.Id = id
		} else {
			panic(ErrorGroupNameExists(g.Name))
		}
	}
}

func (s GroupService) DeleteGroup(gid int64) {
	group := s.GetGroup(gid, nil)

	o := s.ormService.NewOrm()

	memb := models.GroupMember{Group: &group}

	// Remove all the user group relationships
	if _, err := o.Delete(&memb, "group_id"); err != nil {
		panic(err)
	}

	if _, err := o.Delete(&group); err != nil {
		panic(err)
	}
}

func (s GroupService) GetAllGroups() (groups []*models.Group) {
	o := s.ormService.NewOrm()

	if _, err := o.QueryTable(new(models.Group)).All(&groups); err != nil {
		panic(err)
	}

	return
}

type GetGroupOptions struct {
	LoadRelated bool
}

func (s GroupService) GetGroup(gid int64, opts *GetGroupOptions) (g models.Group) {
	o := s.ormService.NewOrm()

	g.Id = gid

	if err := o.Read(&g); err == orm.ErrNoRows {
		panic(ErroGroupIdNotFound(gid))
	}

	if opts != nil {
		if opts.LoadRelated {
			o.LoadRelated(&g, "Members")
		}
	}

	return
}

func (s GroupService) GetGroupByName(n string) (g models.Group) {
	o := s.ormService.NewOrm()

	g.Name = n

	if err := o.Read(&g, "name"); err == orm.ErrNoRows {
		panic(ErroGroupNotFound(n))
	}

	o.LoadRelated(&g, "Members")

	return
}

func (s GroupService) RemoveMember(gid int64, uid int64) {
	group := s.GetGroup(gid, nil)

	user := models.User{Id: uid}
	memb := models.GroupMember{User: &user, Group: &group}

	s.ormService.NewOrm().Delete(&memb, "user_id")
}

func (s GroupService) UpdateGroup(g models.Group) {
	if g.Name == "" {
		panic(ErrorNothingToUpdate(g))
	}

	// Check if group actually exists
	group := s.GetGroup(g.Id, nil)
	group.Name = g.Name

	if _, err := s.ormService.NewOrm().Update(&group); err != nil {
		panic(err)
	}
}
