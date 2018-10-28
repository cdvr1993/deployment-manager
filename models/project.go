package models

type Project struct {
	Id    int64
	Name  string
	Owner *User `orm:"rel(fk);on_delete(do_nothing)"`
}

type ProjectMember struct {
	Id      int64
	Project *Project `orm:"rel(fk);on_delete(do_nothing)"`
	User    *User    `orm:"rel(fk);on_delete(do_nothing)"`
}

type ProjectGroupMember struct {
	Id      int64
	Project *Project `orm:"rel(fk);on_delete(do_nothing)"`
	Group   *Group   `orm:"rel(fk);on_delete(do_nothing)"`
}

func (u *ProjectMember) TableName() string {
	return "project_user"
}

func (u *ProjectGroupMember) TableName() string {
	return "project_group"
}
