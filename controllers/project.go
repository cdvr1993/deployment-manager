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
	defer c.ServeJSON()
	defer services.RecoverUnexpectedError(&c.Controller)

	projects, err := c.ProjectService.ListProjects()

	if err != nil {
		c.Data["json"] = services.TransformError(err)
		return
	}

	c.Data["json"] = ResponseGetAllProjects{projects}
}
