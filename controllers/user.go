package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/cdvr1993/deployment-manager/models"
	"github.com/cdvr1993/deployment-manager/services"
)

// Operations about Users
type UserController struct {
	beego.Controller
	UserService services.IUserService
}

type ResponseCreateUser struct {
	Data models.User
}

// @Title CreateUser
// @Description Create users
// @Param	body	body 	models.User	true	"Body for new user"
// @Success 200 {object} controllers.ResponseCreateUser
// @router / [post]
func (c *UserController) Post() {
	defer services.ServeJson(&c.Controller)

	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	c.UserService.AddUser(&user)
	c.Data["json"] = ResponseCreateUser{user}
}

type ResponseGetAllUsers struct {
	Data []*models.User
}

// @Title GetAll
// @Description Get all Users
// @Success 200 {object} controllers.ResponseGetAllUsers
// @router / [get]
func (c *UserController) GetAll() {
	defer services.ServeJson(&c.Controller)

	c.Data["json"] = ResponseGetAllUsers{c.UserService.GetAll()}
}

type ResponseGetUserByEmail struct {
	Data models.User
}

// @Title Get
// @Description Get user by email
// @Param	email	path 	string	true	"The user's email"
// @Success 200 {object} controllers.ResponseGetUserByEmail
// @router /:email [get]
func (c *UserController) Get() {
	defer services.ServeJson(&c.Controller)

	email := c.GetString(":email")
	user := c.UserService.GetUserByEmail(email)
	c.Data["json"] = ResponseGetUserByEmail{user}
}

type RequestUpdateUser struct {
	Name string
}

type ResponseUpdateUser struct {
	Date string
}

// @Title Put
// @Description Update user by id
// @Param	user_id	path 	number	true	"The user's id"
// @Param	body	body 	controllers.RequestUpdateUser	true	"Body to update the user"
// @Success 200 {object} controllers.ResponseUpdateUser
// @router /:user_id [put]
func (c *UserController) UpdateUser() {
	defer services.ServeJson(&c.Controller)

	var request RequestUpdateUser
	json.Unmarshal(c.Ctx.Input.RequestBody, &request)

	user_id, _ := c.GetInt64(":user_id")
	c.UserService.UpdateUser(models.User{
		Id:   user_id,
		Name: request.Name,
	})
	c.Data["json"] = ResponseUpdateUser{"User updated successfully"}
}

type ResponseDeleteUser struct {
	Data string
}

// @Title Delete
// @Description Delete user by id
// @Param	user_id	path 	number	true	"The user's id"
// @Success 200 {object} controllers.ResponseDeleteUser
// @router /:user_id [delete]
func (c *UserController) DeleteUser() {
	defer services.ServeJson(&c.Controller)

	user_id, _ := c.GetInt64(":user_id")
	c.UserService.DeleteUser(user_id)
	c.Data["json"] = ResponseDeleteUser{"User deleted successfully"}
}
