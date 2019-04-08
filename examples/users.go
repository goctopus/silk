package examples

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
	})
	user.Clean()
}

func (user *Users) All() []Users {
	return make([]Users, 0)
}

func (user *Users) First() Users {
	var u Users
	user.DB.FormFirst(&u)
	user.Clean()
	return u
}
