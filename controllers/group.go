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

type ResponseCreateGroup struct {
	Data models.Group
}

// @Title Post
// @Description Create new group
// @Param	body	body 	models.Group	true	"Body for group content"
// @Success 200 {object} controllers.ResponseCreateGroup
// @router / [post]
func (c *GroupController) Post() {
	defer services.ServeJson(&c.Controller)

	var group models.Group
	json.Unmarshal(c.Ctx.Input.RequestBody, &group)

	c.GroupService.CreateGroup(&group)
	c.Data["json"] = ResponseCreateGroup{group}
}

type ResponseGetGroupByName struct {
	Data models.Group
}

// @Title Get
// @Description get group by name
// @Param	name	path 	string	true	"The name of the group"
// @Success 200 {object} controllers.ResponseGetGroupByName
// @router /:name [get]
func (c *GroupController) Get() {
	defer services.ServeJson(&c.Controller)

	name := c.GetString(":name")
	group := c.GroupService.GetGroup(name)
	c.Data["json"] = ResponseGetGroupByName{group}
}

type RequestAddUser struct {
	Id int64
}

type ResponseAddMember struct {
	Data string
}

// @Title Put
// @Description Add member
// @Param	group_id	path 	string	true	"Group in which the user will be added"
// @Param	body	body controllers.RequestAddUser	true	"Body with id of the user to be added"
// @Success 200 {object} controllers.ResponseAddMember
// @router /:group_id/member [put]
func (c *GroupController) AddMember() {
	defer services.ServeJson(&c.Controller)

	var req RequestAddUser
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	groupId, _ := c.GetInt64(":group_id")

	c.GroupService.AddMember(groupId, req.Id)
	c.Data["json"] = ResponseAddMember{"Member added"}
}

type RequestRemoveUser struct {
	Id int64
}

type ResponseRemoveMember struct {
	Data string
}

// @Title Delete
// @Description Add member
// @Param	group_id	path 	string	true	"Group in which the user will be removed"
// @Param	body	body controllers.RequestRemoveUser	true	"Body with id of the user to be removed"
// @Success 200 {object} controllers.ResponseAddMember
// @router /:group_id/member [delete]
func (c *GroupController) RemoveMember() {
	defer services.ServeJson(&c.Controller)

	var req RequestRemoveUser
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	groupId, _ := c.GetInt64(":group_id")

	c.GroupService.RemoveMember(groupId, req.Id)
	c.Data["json"] = ResponseRemoveMember{"Member removed"}
}
