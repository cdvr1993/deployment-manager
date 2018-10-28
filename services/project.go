package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/cdvr1993/deployment-manager/models"
)

type IProjectService interface {
	ListProjects() []*models.Project
}

type ProjectService struct {
	groupService IGroupService
	ormService   IOrmService
	userService  IUserService
}

var (
	projectService = ProjectService{
		groupService: NewGroupService(),
		ormService:   NewOrmService(),
		userService:  NewUserService(),
	}
)

func NewProjectService() *ProjectService {
	return &projectService
}

func (s ProjectService) ListProjects() (result []*models.Project) {
	qs := s.ormService.NewOrm().QueryTable(new(models.Project))

	if _, err := qs.All(&result); err != nil && err != orm.ErrNoRows {
		panic(err)
	}

	return
}
