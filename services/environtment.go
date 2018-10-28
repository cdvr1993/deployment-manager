package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/cdvr1993/deployment-manager/models"
)

type IEnvironmentService interface {
	ListEnvironments() []*models.Environment
}

type EnvironmentService struct {
	groupService   IGroupService
	ormService     IOrmService
	projectService IProjectService
	userService    IUserService
}

var (
	environmentService = EnvironmentService{
		groupService:   NewGroupService(),
		ormService:     NewOrmService(),
		projectService: NewProjectService(),
		userService:    NewUserService(),
	}
)

func NewEnvironmentService() *EnvironmentService {
	return &environmentService
}

func (s EnvironmentService) ListEnvironments() (result []*models.Environment) {
	qs := s.ormService.NewOrm().QueryTable(new(models.Environment))

	if _, err := qs.All(&result); err != nil && err != orm.ErrNoRows {
		panic(err)
	}

	return
}
