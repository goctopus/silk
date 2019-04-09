package models

import (
	"github.com/goctopus/silk"
	"github.com/goctopus/silk/dialect"
)

type Users struct {
	UsersModel
	silk.Model
}

type UsersModel struct {
	Name  string `json:"name"`
	Id    int64  `json:"id"`
	Exist bool
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

func (user *Users) All() []UsersModel {
	return make([]UsersModel, 0)
}

func (user *Users) Collection() silk.Collection {
	info, _ := user.DB.First()
	return silk.Collection(info)
}

func (user *Users) Collections() silk.Collections {
	info, _ := user.DB.All()
	return silk.GetCollections(info)
}

func (user *Users) Delete() {
	user.DB.Delete()
	user.Clean()
}

func (user *Users) First() UsersModel {
	var u UsersModel
	info, _ := user.DB.First()

	if info != nil {
		u.Id = info["id"].(int64)
		u.Name = info["name"].(string)
		u.Exist = true
	}

	user.Clean()
	return u
}
