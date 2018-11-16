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
		resp, status := services.TransformError(err)
		c.Data["json"] = resp
		c.Ctx.ResponseWriter.WriteHeader(status)
		return
	}

	c.Data["json"] = ResponseGetAllProjects{projects}
}
