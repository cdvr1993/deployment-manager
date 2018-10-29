package middleware

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/cdvr1993/deployment-manager/services"
)

const (
	EMAIL_HEADER = "email"
	USER_PARAM   = "user"
)

type UserParserMiddleware struct {
	UserService services.IUserService
}

func (m UserParserMiddleware) GetHandler() beego.FilterFunc {
	return func(c *context.Context) {
		defer func() {
			recover()
		}()

		c.Input.SetData(
			USER_PARAM,
			m.UserService.GetUserByEmail(c.Input.Header(EMAIL_HEADER)),
		)
	}
}

func (m UserParserMiddleware) GetPaths() []string {
	return []string{
		"/*",
	}
}

func (m UserParserMiddleware) GetPosition() int {
	return beego.BeforeRouter
}
