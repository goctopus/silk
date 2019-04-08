package examples

import (
	"github.com/chenhg5/silk"
	"github.com/chenhg5/silk/dialect"
)

type Users struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
	db   *silk.Sql
}

func NewUsers() *Users {
	return &Users{
		Name: "",
		Id:   0,
		db:   silk.Table("users"),
	}
}

func (user *Users) Where(field string, op string, value interface{}) *Users {
	user.db = user.db.Where(field, op, value)
	return user
}

func (user *Users) WhereId(value interface{}) *Users {
	return user.Where("id", "=", value)
}

func (user *Users) WhereName(value interface{}) *Users {
	return user.Where("name", "=", value)
}

func (user *Users) Save() {
	user.db.Insert(dialect.H{
		"name": user.Name,
	})
}

func (user *Users) All() []Users {
	return make([]Users, 0)
}

func (user *Users) First() Users {
	var u Users
	user.db.FormFirst(&u)
	return u
}
