package main

import (
	"fmt"

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
	m.SQL(`
		create table user (
			id int primary key auto_increment,
			name varchar(128) not null,
			email varchar(128) not null unique
		)
	`)

	m.SQL(`
		create table .group (
			id int primary key auto_increment,
			name varchar(128) not null unique
		)
	`)

	m.SQL(`
		create table .role (
			id int primary key auto_increment,
			name varchar(32) not null unique
		)
	`)

	roles := []string{"SuperAdministrator", "Owner", "Administrator", "Viewer"}
	for _, role := range roles {
		m.SQL(fmt.Sprintf("insert into .role values (NULL, '%s')", role))
	}

	m.SQL(`
		create table user_group (
			id int primary key auto_increment,
			user_id int not null,
			group_id int not null,
			role_id int not null,
			foreign key (user_id) references user(id),
			foreign key (group_id) references ` + "`group`" + `(id),
			foreign key (role_id) references role(id)
		)
	`)
}

// Reverse the migrations
func (m *UserGroups_20181016_145311) Down() {
	m.SQL("drop table user_group")
	m.SQL("drop table user")
	m.SQL("drop table group")
}
