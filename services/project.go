package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/cdvr1993/deployment-manager/models"
)

type IProjectService interface {
	ListProjects() ([]*models.Project, error)
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

func (s ProjectService) ListProjects() ([]*models.Project, error) {
	qs := s.ormService.NewOrm().QueryTable(new(models.Project))

	var result []*models.Project
	if _, err := qs.All(&result); err != nil && err != orm.ErrNoRows {
		return nil, err
	}

	return result, nil
}
