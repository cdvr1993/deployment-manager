package models

type Group struct {
	Id      int64
	Name    string
	Members []GroupMember `orm:"-"`
}

type GroupMember struct {
	Id    int64  `json:"-"`
	Group *Group `json:"-" orm:"rel(fk);on_delete(do_nothing)"`
	User  *User  `orm:"rel(fk);on_delete(do_nothing)"`
	Role  *Role  `orm:"rel(fk);on_delete(do_nothing)"`
}

func (u *GroupMember) TableName() string {
	return "user_group"
}
