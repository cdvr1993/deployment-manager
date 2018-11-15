package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cdvr1993/deployment-manager/models"
	"github.com/cdvr1993/deployment-manager/services"
)

type EnvironmentController struct {
	beego.Controller
	EnvironmentService services.IEnvironmentService
}

type ResponseGetAllEnvironments struct {
	Data []*models.Environment
}

// @Title Get
// @Description Get all environments
// @Success 200 {object} controllers.ResponseGetAllEnvironments
// @router / [get]
func (c *EnvironmentController) GetAll() {
	defer c.ServeJSON()
	defer services.RecoverUnexpectedError(&c.Controller)

	environments, err := c.EnvironmentService.ListEnvironments()

	if err != nil {
		c.Data["json"] = services.TransformError(err)
		return
	}

	c.Data["json"] = ResponseGetAllEnvironments{environments}
}
