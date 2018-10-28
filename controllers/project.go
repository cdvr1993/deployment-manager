package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cdvr1993/deployment-manager/models"
	"github.com/cdvr1993/deployment-manager/services"
)

type ProjectController struct {
	beego.Controller
	ProjectService services.IProjectService
}

type ResponseGetAllProjects struct {
	Data []*models.Project
}

// @Title Get
// @Description Get all projects
// @Success 200 {object} controllers.ResponseGetAllProjects
// @router / [get]
func (c *ProjectController) GetAll() {
	defer services.ServeJson(&c.Controller)

	projects := c.ProjectService.ListProjects()
	c.Data["json"] = ResponseGetAllProjects{projects}
}
