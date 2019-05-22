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
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Sex       int64  `json:"sex"`
	Country   string `json:"country"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Exist     bool
}

// construct methods

func Users() *UsersBuilder {
	return &UsersBuilder{
		db:    silk.Table("users"),
		Table: "users",
	}
}

// intermediate methods

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

func (builder *UsersBuilder) WhereSex(value interface{}) *UsersBuilder {
	builder.db = builder.Where("sex", "=", value)
	return builder
}

func (builder *UsersBuilder) WhereCountry(value interface{}) *UsersBuilder {
	builder.db = builder.Where("country", "=", value)
	return builder
}

func (builder *UsersBuilder) OrWhere(field string, value interface{}) *silk.Builder {
	panic("implement it")
}

// terminal methods

// TODO: 怎么知道模型字段被赋值了没有
func (builder *UsersBuilder) Save() {
	if builder.Id != 0 {
		_, _ = builder.db.Insert(dialect.H{
			"id":      builder.Id,
			"name":    builder.Name,
			"sex":     builder.Sex,
			"country": builder.Country,
		})
	} else {
		_, _ = builder.db.Insert(dialect.H{
			"name":    builder.Name,
			"sex":     builder.Sex,
			"country": builder.Country,
		})
	}
	builder.Clean()
}

func (builder *UsersBuilder) All() []UsersModel {
	users := make([]UsersModel, 0)

	info, _ := builder.db.All()

	for i := 0; i < len(info); i++ {
		var u UsersModel
		u.Id = info[i]["id"].(int64)
		if name, ok := info[i]["name"]; ok {
			u.Name = string(name.([]uint8))
		}
		if sex, ok := info[i]["sex"]; ok {
			u.Sex = sex.(int64)
		}
		if country, ok := info[i]["country"]; ok {
			u.Country = string(country.([]uint8))
		}
		if createdAt, ok := info[i]["created_at"]; ok {
			u.CreatedAt = createdAt.(string)
		}
		if updatedAt, ok := info[i]["updated_at"]; ok {
			u.UpdatedAt = updatedAt.(string)
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
	_ = builder.db.Delete()
	builder.Clean()
}

func (builder *UsersBuilder) Find(value interface{}) UsersModel {
	defer builder.Clean()
	return builder.WhereId(value).First()
}

func (builder *UsersBuilder) FindOrFail(value interface{}) UsersModel {
	defer builder.Clean()
	u := builder.WhereId(value).First()
	if !u.Exist {
		panic("not found model")
	}
	return u
}

func (builder *UsersBuilder) First() UsersModel {
	var u UsersModel
	info, _ := builder.db.First()
	u.Id = info["id"].(int64)
	if name, ok := info["name"]; ok {
		u.Name = string(name.([]uint8))
	}
	if sex, ok := info["sex"]; ok {
		u.Sex = sex.(int64)
	}
	if country, ok := info["country"]; ok {
		u.Country = string(country.([]uint8))
	}
	if createdAt, ok := info["created_at"]; ok {
		u.CreatedAt = createdAt.(string)
	}
	if updatedAt, ok := info["updated_at"]; ok {
		u.UpdatedAt = updatedAt.(string)
	}
	u.Exist = true
	builder.Clean()
	return u
}

func (builder *UsersBuilder) Update(values dialect.H) (int64, error) {
	defer builder.Clean()
	return builder.db.Update(values)
}

func (builder *UsersBuilder) Insert(values dialect.H) (int64, error) {
	defer builder.Clean()
	return builder.db.Insert(values)
}

// help methods

func (builder *UsersBuilder) Clean() {
	builder.db = silk.Table(builder.Table)
}
