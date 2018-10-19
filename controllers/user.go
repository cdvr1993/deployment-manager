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

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	defer services.ServeJson(&c.Controller)

	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	c.UserService.AddUser(&user)
	c.Data["json"] = map[string]models.User{"data": user}
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	// users := models.GetAllUsers()
	// u.Data["json"] = users
	// u.ServeJSON()
}

// @Title Get
// @Description get user by email
// @Param	email		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :email is empty
// @router /:email [get]
func (c *UserController) Get() {
	defer services.ServeJson(&c.Controller)

	email := c.GetString(":email")
	user := c.UserService.GetUserByEmail(email)
	c.Data["json"] = map[string]models.User{"data": user}
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	// uid := u.GetString(":uid")
	// if uid != "" {
	// 	var user models.User
	// 	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	// 	uu, err := models.UpdateUser(uid, &user)
	// 	if err != nil {
	// 		u.Data["json"] = err.Error()
	// 	} else {
	// 		u.Data["json"] = uu
	// 	}
	// }
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	// uid := u.GetString(":uid")
	// models.DeleteUser(uid)
	// u.Data["json"] = "delete success!"
	u.ServeJSON()
}
