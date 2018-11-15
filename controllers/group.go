package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/cdvr1993/deployment-manager/middleware"
	"github.com/cdvr1993/deployment-manager/models"
	"github.com/cdvr1993/deployment-manager/services"
)

type GroupController struct {
	beego.Controller
	GroupService services.IGroupService
}

type ResponseGetAllGroups struct {
	Data []*models.Group
}

// @Title Get
// @Description Get all groups
// @Success 200 {object} controllers.ResponseGetAllGroups
// @router / [get]
func (c *GroupController) GetAll() {
	defer c.ServeJSON()
	defer services.RecoverUnexpectedError(&c.Controller)

	groups, err := c.GroupService.GetAllGroups(services.GetAllGroupsOptions{
		Email: c.Ctx.Input.Header(middleware.EMAIL_HEADER),
	})

	if err != nil {
		c.Data["json"] = services.TransformError(err)
		return
	}

	c.Data["json"] = ResponseGetAllGroups{groups}
}

type ResponseCreateGroup struct {
	Data *models.Group
}

// @Title Post
// @Description Create new group
// @Param	body	body 	models.Group	true	"Body for group content"
// @Success 200 {object} controllers.ResponseCreateGroup
// @router / [post]
func (c *GroupController) Post() {
	defer c.ServeJSON()
	defer services.RecoverUnexpectedError(&c.Controller)

	var group models.Group
	json.Unmarshal(c.Ctx.Input.RequestBody, &group)

	err := c.GroupService.CreateGroup(
		&group,
		c.Ctx.Input.Header(middleware.EMAIL_HEADER),
	)

	if err != nil {
		c.Data["json"] = services.TransformError(err)
		return
	}

	c.Data["json"] = ResponseCreateGroup{&group}
}

type ResponseGetGroupByName struct {
	Data *models.Group
}

// @Title Get
// @Description Get group by name
// @Param	name	path 	string	true	"The name of the group"
// @Param	body	body 	models.Group	true	"Body for group content"
// @Success 200 {object} controllers.ResponseGetGroupByName
// @router /:name [get]
func (c *GroupController) Get() {
	defer c.ServeJSON()
	defer services.RecoverUnexpectedError(&c.Controller)

	name := c.GetString(":name")
	group, err := c.GroupService.GetGroupByName(name)

	if err != nil {
		c.Data["json"] = services.TransformError(err)
		return
	}

	c.Data["json"] = ResponseGetGroupByName{group}
}

type ResponseUpdateGroup struct {
	Data string
}

// @Title Put
// @Description Update group
// @Param	group_id	path 	number	true	"The id of the group"
// @Param	body	body 	models.Group	true	"Body for group content"
// @Success 200 {object} controllers.ResponseUpdateGroup
// @router /:group_id [put]
// It helps on changing the name of the group without loosing user membership
func (c *GroupController) UpdateGroup() {
	defer c.ServeJSON()
	defer services.RecoverUnexpectedError(&c.Controller)

	var group models.Group
	json.Unmarshal(c.Ctx.Input.RequestBody, &group)

	group.Id, _ = c.GetInt64(":group_id")
	err := c.GroupService.UpdateGroup(&group)

	if err != nil {
		c.Data["json"] = services.TransformError(err)
		return
	}

	c.Data["json"] = ResponseUpdateGroup{
		fmt.Sprintf("Group '%s' updated successfully", group.Name),
	}
}

type ResponseDeleteGroup struct {
	Data string
}

// @Title Delete
// @Description Delete group
// @Param	group_id	path 	number	true	"The id of the group"
// @Success 200 {object} controllers.ResponseDeleteGroup
// @router /:group_id [delete]
func (c *GroupController) DeleteGroup() {
	defer c.ServeJSON()
	defer services.RecoverUnexpectedError(&c.Controller)

	groupId, _ := c.GetInt64(":group_id")
	err := c.GroupService.DeleteGroup(groupId)

	if err != nil {
		c.Data["json"] = services.TransformError(err)
		return
	}

	c.Data["json"] = ResponseUpdateGroup{"Group deleted successfully"}
}

type RequestAddUser struct {
	Id   int64
	Role string
}

type ResponseAddMember struct {
	Data string
}

// @Title Post
// @Description Add member
// @Param	group_id	path 	number	true	"Group in which the user will be added"
// @Param	body	body controllers.RequestAddUser	true	"Body with id of the user to be added"
// @Success 200 {object} controllers.ResponseAddMember
// @router /:group_id/member [post]
func (c *GroupController) AddMember() {
	defer c.ServeJSON()
	defer services.RecoverUnexpectedError(&c.Controller)

	var req RequestAddUser
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	groupId, _ := c.GetInt64(":group_id")

	err := c.GroupService.AddMember(groupId, req.Id, req.Role)

	if err != nil {
		c.Data["json"] = services.TransformError(err)
		return
	}

	c.Data["json"] = ResponseAddMember{"Member added successfully"}
}

type ResponseRemoveMember struct {
	Data string
}

// @Title Delete
// @Description Add member
// @Param	group_id	path 	number	true	"Group in which the user will be removed"
// @Param	member_id	path	number	true	"Id of the user to be removed"
// @Success 200 {object} controllers.ResponseAddMember
// @router /:group_id/member/:member_id [delete]
func (c *GroupController) RemoveMember() {
	defer c.ServeJSON()
	defer services.RecoverUnexpectedError(&c.Controller)

	groupId, _ := c.GetInt64(":group_id")
	member_id, _ := c.GetInt64(":member_id")

	err := c.GroupService.RemoveMember(groupId, member_id)

	if err != nil {
		c.Data["json"] = services.TransformError(err)
		return
	}

	c.Data["json"] = ResponseRemoveMember{"Member removed successfully"}
}
