package services

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
)

func ServeJson(c *beego.Controller) {
	if r := recover(); r != nil {
		if svc, ok := r.(IServiceError); ok {
			c.Ctx.ResponseWriter.WriteHeader(svc.Status())
			c.Data["json"] = map[string][]string{"errors": []string{svc.Error()}}
		} else {
			// Create an unique id to make a relationship between
			// API error and log error
			var u1String string

			if u1, err := uuid.NewV4(); err == nil {
				u1String = fmt.Sprintf("%s", u1)
			} else {
				u1String = "Unknown"
			}

			beego.Error(fmt.Sprintf("%s - %s", u1String, r))
			c.Ctx.ResponseWriter.WriteHeader(500)
			c.Data["json"] = map[string][]string{
				"errors": []string{fmt.Sprintf("Transaction: '%s'", u1String)},
			}
		}
	}
	c.ServeJSON()
}
