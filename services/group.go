package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/cdvr1993/deployment-manager/models"
)

type IGroupService interface {
	AddMember(int64, int64, string)
	CreateGroup(*models.Group, string)
	DeleteGroup(int64)
	GetAllGroups(GetAllGroupsOptions) []*models.Group
	GetGroup(int64, *GetGroupOptions) models.Group
	GetGroupByName(string) models.Group
	IsAllowed(models.Group, models.User, bool) bool
	RemoveMember(int64, int64)
	UpdateGroup(models.Group)
}

type GroupService struct {
	ormService  IOrmService
	roleService IRoleService
	userService IUserService
}

var (
	groupService = GroupService{
		ormService:  NewOrmService(),
		roleService: NewRoleService(),
		userService: NewUserService(),
	}
)

func NewGroupService() *GroupService {
	return &groupService
}

func (s GroupService) AddMember(gid, uid int64, rName string) {
	role := s.roleService.GetRole(rName)

	user := s.userService.GetUser(uid)
	group := s.GetGroup(gid, nil)

	memb := models.GroupMember{User: &user, Group: &group, Role: &role}

	s.ormService.NewOrm().ReadOrCreate(&memb, "user_id", "group_id")
}

func (s GroupService) CreateGroup(g *models.Group, e string) {
	o := s.ormService.NewOrm()

	user := s.userService.GetUserByEmail(e)

	if created, id, err := o.ReadOrCreate(g, "name"); err == nil {
		if created {
			g.Id = id

			s.AddMember(g.Id, user.Id, "Owner")
		} else {
			panic(ErrorGroupNameExists(g.Name))
		}
	} else {
		panic(err)
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

type GetAllGroupsOptions struct {
	Email string
}

func (s GroupService) GetAllGroups(opts GetAllGroupsOptions) (groups []*models.Group) {
	u := s.userService.GetUserByEmail(opts.Email)

	o := s.ormService.NewOrm()

	var groupMs []*models.GroupMember
	qs := o.QueryTable(new(models.GroupMember)).Filter("user_id", u.Id)

	// Grab the groups that the user has access to
	if _, err := qs.All(&groupMs); err != nil && err != orm.ErrNoRows {
		panic(err)
	}

	ids := make([]interface{}, len(groupMs))
	for i, _ := range groupMs {
		ids[i] = groupMs[i].Id
	}

	qs = o.
		QueryTable(new(models.Group)).
		Filter("id__in", ids...)

	if _, err := qs.All(&groups); err != nil && err != orm.ErrNoRows {
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
			s.loadMembers(&g)
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

	s.loadMembers(&g)

	return
}

func (s GroupService) IsAllowed(g models.Group, u models.User, ro bool) bool {
	if g.Id > 0 {
		g = s.GetGroup(g.Id, nil)
	} else {
		g = s.GetGroupByName(g.Name)
	}

	groupM := new(models.GroupMember)

	s.
		ormService.
		NewOrm().
		QueryTable(groupM).
		Filter("group_id", g.Id).
		Filter("user_id", u.Id).
		RelatedSel("role").
		One(groupM)

	if !ro && groupM.Role.Name == "Viewer" {
		return false
	}

	return true
}

func (s GroupService) loadMembers(g *models.Group) {
	s.
		ormService.
		NewOrm().
		QueryTable(new(models.GroupMember)).
		Filter("group_id", g.Id).
		RelatedSel("user").
		RelatedSel("role").
		All(&g.Members)
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
