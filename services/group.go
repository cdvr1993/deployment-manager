package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/cdvr1993/deployment-manager/models"
)

type IGroupService interface {
	AddMember(int64, int64)
	CreateGroup(*models.Group)
	GetGroup(string) models.Group
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
	group := models.Group{Id: gid}

	o := s.ormService.NewOrm()

	if err := o.Read(&group); err == orm.ErrNoRows {
		panic(ErroGroupIdNotFound(gid))
	}

	memb := models.GroupMember{User: &user, Group: &group}

	o.ReadOrCreate(&memb, "user_id", "group_id")
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

func (s GroupService) GetGroup(n string) (g models.Group) {
	o := s.ormService.NewOrm()

	g.Name = n

	if err := o.Read(&g, "name"); err == orm.ErrNoRows {
		panic(ErroGroupNotFound(n))
	}

	o.LoadRelated(&g, "Members")

	return
}
