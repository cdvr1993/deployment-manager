package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Group))
	orm.RegisterModel(new(GroupMember))
	orm.RegisterModel(new(Role))
	orm.RegisterModel(new(User))
}
