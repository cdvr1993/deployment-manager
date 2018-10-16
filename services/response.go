package services

import (
	"github.com/astaxie/beego"
)

func ServeJson(c *beego.Controller) {
	if r := recover(); r != nil {
		if svc, ok := r.(IServiceError); ok {
			c.Ctx.ResponseWriter.WriteHeader(svc.Status())
			c.Data["json"] = map[string][]string{"errors": []string{svc.Error()}}
		}
	}
	c.ServeJSON()
}
