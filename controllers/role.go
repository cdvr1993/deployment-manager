package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cdvr1993/deployment-manager/models"
	"github.com/cdvr1993/deployment-manager/services"
)

type RoleController struct {
	beego.Controller
	RoleService services.IRoleService
}

type ResponseGetAllRoles struct {
	Data []*models.Role
}

// @Title Get
// @Description Get all roles
// @Success 200 {object} controllers.ResponseGetAllRoles
// @router / [get]
func (c *RoleController) GetAll() {
	defer c.ServeJSON()
	defer services.RecoverUnexpectedError(&c.Controller)

	roles, err := c.RoleService.ListRoles()

	if err != nil {
		c.Data["json"] = services.TransformError(err)
		return
	}

	c.Data["json"] = ResponseGetAllRoles{roles}
}
