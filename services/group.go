package services

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/cdvr1993/deployment-manager/models"
)

type IGroupService interface {
	AddMember(int64, int64, string) error
	CreateGroup(*models.Group, string) error
	DeleteGroup(int64) error
	GetAllGroups(GetAllGroupsOptions) ([]*models.Group, error)
	GetGroup(int64, *GetGroupOptions) (*models.Group, error)
	GetGroupByName(string) (*models.Group, error)
	IsAllowed(*models.Group, *models.User, bool) bool
	RemoveMember(int64, int64) error
	UpdateGroup(*models.Group) error
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

func (s GroupService) AddMember(gid, uid int64, rName string) error {
	role, err := s.roleService.GetRole(rName)

	if err != nil {
		return err
	}

	user, err := s.userService.GetUser(uid)

	if err != nil {
		return err
	}

	group, err := s.GetGroup(gid, nil)

	if err != nil {
		return err
	}

	memb := models.GroupMember{User: user, Group: group, Role: role}

	if _, _, err := s.ormService.NewOrm().ReadOrCreate(&memb, "user_id", "group_id"); err != nil {
		return err
	}

	return nil
}

func (s GroupService) CreateGroup(g *models.Group, e string) error {
	o := s.ormService.NewOrm()

	user, err := s.userService.GetUserByEmail(e)

	if err != nil {
		return err
	}

	created, id, err := o.ReadOrCreate(g, "name")

	if err != nil {
		return err
	}

	if !created {
		return ErrorGroupNameExists(g.Name)
	}

	g.Id = id
	if err := s.AddMember(g.Id, user.Id, "Owner"); err != nil {
		return err
	}

	return nil
}

func (s GroupService) DeleteGroup(gid int64) error {
	group, err := s.GetGroup(gid, nil)

	if err != nil {
		return err
	}

	o := s.ormService.NewOrm()

	memb := models.GroupMember{Group: group}

	// Remove all the user group relationships
	if _, err := o.Delete(&memb, "group_id"); err != nil {
		return err
	}

	if _, err := o.Delete(&group); err != nil {
		return err
	}

	return nil
}

type GetAllGroupsOptions struct {
	User *models.User
}

func (s GroupService) GetAllGroups(opts GetAllGroupsOptions) ([]*models.Group, error) {
	groups := make([]*models.Group, 0)

	if opts.User == nil {
		return groups, nil
	}

	o := s.ormService.NewOrm()

	var groupMs []*models.GroupMember
	qs := o.QueryTable(new(models.GroupMember)).Filter("user_id", opts.User.Id)

	// Grab the groups that the user has access to
	if _, err := qs.All(&groupMs); err != nil && err != orm.ErrNoRows {
		return nil, err
	}

	if len(groupMs) == 0 {
		return groups, nil
	}

	ids := make([]interface{}, len(groupMs))
	for i, _ := range groupMs {
		ids[i] = groupMs[i].Id
	}

	qs = o.
		QueryTable(new(models.Group)).
		Filter("id__in", ids...)

	if _, err := qs.All(&groups); err != nil && err != orm.ErrNoRows {
		return nil, err
	}

	return groups, nil
}

type GetGroupOptions struct {
	LoadRelated bool
}

func (s GroupService) GetGroup(gid int64, opts *GetGroupOptions) (*models.Group, error) {
	o := s.ormService.NewOrm()
	g := models.Group{Id: gid}

	if err := o.Read(&g); err == orm.ErrNoRows {
		return nil, ErroGroupIdNotFound(gid)
	}

	if opts != nil {
		if opts.LoadRelated {
			s.loadMembers(&g)
		}
	}

	return &g, nil
}

func (s GroupService) GetGroupByName(n string) (*models.Group, error) {
	o := s.ormService.NewOrm()

	g := models.Group{Name: n}
	if err := o.Read(&g, "name"); err == orm.ErrNoRows {
		return nil, ErroGroupNotFound(n)
	}

	s.loadMembers(&g)

	return &g, nil
}

func (s GroupService) IsAllowed(g *models.Group, u *models.User, ro bool) bool {
	var err error

	// Check if they exist
	if g.Id > 0 {
		_, err = s.GetGroup(g.Id, nil)
	} else {
		_, err = s.GetGroupByName(g.Name)
	}

	if err != nil {
		beego.Error(err)
		return false
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

func (s GroupService) RemoveMember(gid int64, uid int64) error {
	group, err := s.GetGroup(gid, nil)

	if err != nil {
		return err
	}

	user := models.User{Id: uid}
	memb := models.GroupMember{User: &user, Group: group}

	if _, err := s.ormService.NewOrm().Delete(&memb, "user_id"); err != nil {
		return err
	}

	return nil
}

func (s GroupService) UpdateGroup(g *models.Group) error {
	if g.Name == "" {
		return ErrorNothingToUpdate(g)
	}

	// Check if group actually exists
	group, err := s.GetGroup(g.Id, nil)

	if err != nil {
		return err
	}

	group.Name = g.Name

	if _, err := s.ormService.NewOrm().Update(group); err != nil {
		return err
	}

	return nil
}
