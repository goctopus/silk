package models

import (
	"github.com/goctopus/silk"
	"github.com/goctopus/silk/dialect"
)

type UsersBuilder struct {
	UsersModel

	db *silk.Builder

	Table string
}

type UsersModel struct {
	Name  string `json:"name"`
	Id    int64  `json:"id"`
	Exist bool
}

func Users() *UsersBuilder {
	return &UsersBuilder{
		db:    silk.Table("users"),
		Table: "users",
	}
}

func (builder *UsersBuilder) Clean() {
	builder.db = silk.Table(builder.Table)
}

func (builder *UsersBuilder) Where(field string, op string, value interface{}) *silk.Builder {
	return builder.db.Where(field, op, value)
}

func (builder *UsersBuilder) WhereId(value interface{}) *UsersBuilder {
	builder.db = builder.Where("id", "=", value)
	return builder
}

func (builder *UsersBuilder) WhereName(value interface{}) *UsersBuilder {
	builder.db = builder.Where("name", "=", value)
	return builder
}

func (builder *UsersBuilder) Save() {
	builder.db.Insert(dialect.H{
		"name": builder.Name,
		"id":   builder.Id,
	})
	builder.Clean()
}

func (builder *UsersBuilder) All() []UsersModel {
	users := make([]UsersModel, 0)

	info, _ := builder.db.All()

	for i := 0; i < len(info); i++ {
		var u UsersModel
		if name, ok := info[i]["name"]; ok {
			u.Name = name.(string)
		}
		if id, ok := info[i]["id"]; ok {
			u.Id = id.(int64)
		}
		users = append(users, u)
	}

	return users
}

func (builder *UsersBuilder) Collection() silk.Collection {
	info, _ := builder.db.All()
	return silk.Collect(info)
}

func (builder *UsersBuilder) Delete() {
	builder.db.Delete()
	builder.Clean()
}

func (builder *UsersBuilder) First() UsersModel {
	var u UsersModel
	info, _ := builder.db.First()
	if name, ok := info["name"]; ok {
		u.Name = name.(string)
	}
	if id, ok := info["id"]; ok {
		u.Id = id.(int64)
	}
	u.Exist = true
	builder.Clean()
	return u
}
