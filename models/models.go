package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Environment))
	orm.RegisterModel(new(EnvironmentMember))
	orm.RegisterModel(new(EnvironmentGroupMember))
	orm.RegisterModel(new(Group))
	orm.RegisterModel(new(GroupMember))
	orm.RegisterModel(new(Project))
	orm.RegisterModel(new(ProjectMember))
	orm.RegisterModel(new(ProjectGroupMember))
	orm.RegisterModel(new(Role))
	orm.RegisterModel(new(User))
}
