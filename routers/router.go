// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/cdvr1993/deployment-manager/controllers"
	"github.com/cdvr1993/deployment-manager/services"
)

type ServiceManager struct {
	EnvironmentService services.IEnvironmentService
	GroupService       services.IGroupService
	ProjectService     services.IProjectService
	RoleService        services.IRoleService
	UserService        services.IUserService
}

// @APIVersion 1.0.0
// @Title Deployment manager API
// @Description Management of your projects and deployments
// @Contact cdvr1993@gmail.com
// @License MIT
// @LicenseUrl https://opensource.org/licenses/MIT
func InitRouter(m ServiceManager) {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/environment",
			beego.NSInclude(
				&controllers.EnvironmentController{
					EnvironmentService: m.EnvironmentService,
				},
			),
		),
		beego.NSNamespace("/group",
			beego.NSInclude(
				&controllers.GroupController{
					GroupService: m.GroupService,
				},
			),
		),
		beego.NSNamespace("/project",
			beego.NSInclude(
				&controllers.ProjectController{
					ProjectService: m.ProjectService,
				},
			),
		),
		beego.NSNamespace("/role",
			beego.NSInclude(
				&controllers.RoleController{
					RoleService: m.RoleService,
				},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{
					UserService: m.UserService,
				},
			),
		),
	)
	beego.AddNamespace(ns)
}

func InjectServices() {
	// Used as a dependency injection
	InitRouter(ServiceManager{
		EnvironmentService: services.NewEnvironmentService(),
		GroupService:       services.NewGroupService(),
		RoleService:        services.NewRoleService(),
		UserService:        services.NewUserService(),
		ProjectService:     services.NewProjectService(),
	})
}
