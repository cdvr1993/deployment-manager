package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Projects_20181026_100202 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Projects_20181026_100202{}
	m.Created = "20181026_100202"

	migration.Register("Projects_20181026_100202", m)
}

// Run the migrations
func (m *Projects_20181026_100202) Up() {
	m.SQL(`
		create table project (
			id int primary key auto_increment,
			name varchar(128) not null unique
		)
	`)

	m.SQL(`
		create table environment (
			id int primary key auto_increment,
			name varchar(128) not null,
			project_id int not null,
			foreign key (project_id) references project(id),
			unique key environment_id (project_id, name)
		)
	`)

	m.SQL(`
		create table project_user (
			id int primary key auto_increment,
			project_id int not null,
			user_id int not null,
			foreign key (project_id) references project(id),
			foreign key (user_id) references user(id)
		)
	`)

	m.SQL(`
		create table project_group (
			id int primary key auto_increment,
			project_id int not null,
			group_id int not null,
			foreign key (project_id) references project(id),
			foreign key (group_id) references ` + "`group`" + `(id)
		)
	`)

	m.SQL(`
		create table environment_user (
			id int primary key auto_increment,
			environment_id int not null,
			user_id int not null,
			foreign key (environment_id) references environment(id),
			foreign key (user_id) references user(id)
		)
	`)

	m.SQL(`
		create table environment_group (
			id int primary key auto_increment,
			environment_id int not null,
			group_id int not null,
			foreign key (environment_id) references environment(id),
			foreign key (group_id) references ` + "`group`" + `(id)
		)
	`)
}

// Reverse the migrations
func (m *Projects_20181026_100202) Down() {
	m.SQL("drop table environment_group")
	m.SQL("drop table environment_user")
	m.SQL("drop table project_group")
	m.SQL("drop table project_user")
	m.SQL("drop table environment")
	m.SQL("drop table project")
}
