package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UserGroups_20181016_145311 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserGroups_20181016_145311{}
	m.Created = "20181016_145311"

	migration.Register("UserGroups_20181016_145311", m)
}

// Run the migrations
func (m *UserGroups_20181016_145311) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`
		create table user (
			id int primary key auto_increment,
			name varchar(128) not null,
			email varchar(128) not null
		)
	`)

	m.SQL(`
		create table .group (
			id int primary key auto_increment,
			name varchar(128) not null
		)
	`)

	m.SQL(`
		create table user_group (
			user_id int not null,
			group_id int not null,
			foreign key (user_id) references user(id),
			foreign key (group_id) references ` + "`group`" + `(id)
		)
	`)
}

// Reverse the migrations
func (m *UserGroups_20181016_145311) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("drop table user_group")
	m.SQL("drop table user")
	m.SQL("drop table group")
}
