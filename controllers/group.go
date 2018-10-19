package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/cdvr1993/deployment-manager/models"
	"github.com/cdvr1993/deployment-manager/services"
)

type GroupController struct {
	beego.Controller
	GroupService services.IGroupService
}

// @Title Post
// @Description Create new group
// @Param	body		body 	models.Group	true		"body for group content"
// @Success 200 {object} models.Group
// @Failure 403 body is empty
// @router / [post]
func (c *GroupController) Post() {
	defer services.ServeJson(&c.Controller)

	var group models.Group
	json.Unmarshal(c.Ctx.Input.RequestBody, &group)

	c.GroupService.CreateGroup(&group)
	c.Data["json"] = map[string]models.Group{"data": group}
}

// @Title Get
// @Description get group by name
// @Param	name		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Group
// @Failure 403 :name is empty
// @router /:name [get]
func (c *GroupController) Get() {
	defer services.ServeJson(&c.Controller)

	name := c.GetString(":name")
	group := c.GroupService.GetGroup(name)
	c.Data["json"] = map[string]models.Group{"data": group}
}

type AddUserRequest struct {
	Id int64
}

// @Title Post
// @Description Add member
// @Param	group_id		path 	string	true	"Group in which the user will be added"
// @Param	body		body 	AddUser	true		"body for user id"
// @Success 200
// @Failure 403 body is empty
// @router /:group_id [post]
func (c *GroupController) AddMember() {
	defer services.ServeJson(&c.Controller)

	var req AddUserRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	groupId, _ := c.GetInt64(":group_id")

	c.GroupService.AddMember(groupId, req.Id)
	c.Data["json"] = map[string]string{"data": "Member added"}
}
