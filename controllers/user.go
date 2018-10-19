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
func (u *UserController) GetAll() {
	// users := models.GetAllUsers()
	// u.Data["json"] = users
	// u.ServeJSON()
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
