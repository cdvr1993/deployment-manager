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
	defer services.ServeJson(&c.Controller)

	environments := c.EnvironmentService.ListEnvironments()
	c.Data["json"] = ResponseGetAllEnvironments{environments}
}
