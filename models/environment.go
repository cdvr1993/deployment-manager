package models

type Environment struct {
	Id      int64
	Name    string
	Project *Project `orm:"rel(fk);on_delete(do_nothing)"`
}

type EnvironmentMember struct {
	Id          int64
	Environment *Environment `orm:"rel(fk);on_delete(do_nothing)"`
	User        *User        `orm:"rel(fk);on_delete(do_nothing)"`
}

type EnvironmentGroupMember struct {
	Id          int64
	Environment *Environment `orm:"rel(fk);on_delete(do_nothing)"`
	Group       *Group       `orm:"rel(fk);on_delete(do_nothing)"`
}

func (u *EnvironmentMember) TableName() string {
	return "environment_user"
}

func (u *EnvironmentGroupMember) TableName() string {
	return "environment_group"
}
