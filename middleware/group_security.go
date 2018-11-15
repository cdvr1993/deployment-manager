package middleware

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/cdvr1993/deployment-manager/models"
	"github.com/cdvr1993/deployment-manager/services"
)

type GroupSecurityMiddleware struct {
	GroupService services.IGroupService
}

func (m GroupSecurityMiddleware) GetHandler() beego.FilterFunc {
	return func(c *context.Context) {
		defer recoverFromPanic(c)

		group := models.Group{Name: c.Input.Param(":group")}

		if groupId, err := strconv.ParseInt(group.Name, 10, 64); err == nil {
			group.Id = groupId
		}

		user := c.Input.GetData(USER_PARAM).(models.User)
		if !m.GroupService.IsAllowed(&group, &user, c.Input.IsGet()) {
			panic("User not anuthorized")
		}
	}
}

func (m GroupSecurityMiddleware) GetPaths() []string {
	return []string{
		"/:version([a-z]+[0-9]+)/group/:group([a-zA-Z0-9_]+)",
		"/:version([a-z]+[0-9]+)/group/:group([a-zA-Z0-9_]+)/:any(.*)",
	}
}

func (m GroupSecurityMiddleware) GetPosition() int {
	return beego.BeforeRouter
}
