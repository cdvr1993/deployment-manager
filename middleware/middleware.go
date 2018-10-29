package middleware

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/cdvr1993/deployment-manager/services"
)

type IMiddleware interface {
	GetHandler() beego.FilterFunc
	GetPaths() []string
	GetPosition() int
}

func recoverFromPanic(c *context.Context) {
	// If middleware fails mark it as 'Access denied'
	if r := recover(); r != nil {
		beego.Error(r)
		c.ResponseWriter.WriteHeader(403)
		c.Output.JSON(map[string][]string{"errors": []string{"Access denied"}}, false, false)
	}
}

func InsertMiddleware(m []IMiddleware) {
	for _, item := range m {
		for _, p := range item.GetPaths() {
			beego.InsertFilter(p, item.GetPosition(), item.GetHandler(), false)
		}
	}

}

func InitMiddleware() {
	// Remove automatic template
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {})

	InsertMiddleware([]IMiddleware{
		UserParserMiddleware{
			UserService: services.NewUserService(),
		},
		GroupSecurityMiddleware{
			GroupService: services.NewGroupService(),
		},
	})
}
