package models

import (
	"github.com/goctopus/silk"
	"github.com/goctopus/silk/dialect"
)

type Users struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`

	silk.Model
}

func NewUsers() *Users {
	return &Users{
		Model: silk.Model{
			DB:    silk.Table("users"),
			Table: "users",
		},
	}
}

func (user *Users) WhereId(value interface{}) *Users {
	user.DB = user.Where("id", "=", value)
	return user
}

func (user *Users) WhereName(value interface{}) *Users {
	user.DB = user.Where("name", "=", value)
	return user
}

func (user *Users) Save() {
	user.DB.Insert(dialect.H{
		"name": user.Name,
		"id":   user.Id,
	})
	user.Clean()
}

func (user *Users) All() []Users {
	return make([]Users, 0)
}

func (user *Users) First() Users {
	var u Users
	info, _ := user.DB.First()
	u.Id = info["id"].(int64)
	u.Name = string(info["name"].([]uint8))
	user.Clean()
	return u
}
