package models

type Group struct {
	Id      int64
	Name    string
	Members []*User `orm:"rel(m2m);rel_through(github.com/cdvr1993/deployment-manager/models.GroupMember)"`
}

type GroupMember struct {
	Id    int64
	Group *Group `orm:"rel(fk);on_delete(do_nothing)"`
	User  *User  `orm:"rel(fk);on_delete(do_nothing)"`
}

func (u *GroupMember) TableName() string {
	return "user_group"
}
