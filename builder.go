// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package silk

import (
	"errors"
	"github.com/goctopus/silk/dialect"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

type Builder struct {
	dialect.SqlComponent
	connection *DB
}

var BuilderPool = sync.Pool{
	New: func() interface{} {
		return &Builder{
			SqlComponent: dialect.SqlComponent{
				Fields:     make([]string, 0),
				TableName:  "",
				Args:       make([]interface{}, 0),
				Wheres:     make([]dialect.Where, 0),
				Leftjoins:  make([]dialect.Join, 0),
				UpdateRaws: make([]dialect.RawUpdate, 0),
				WhereRaws:  "",
			},
			connection: DBInstance.clone(),
		}
	},
}

type H map[string]interface{}

func newBuilder() *Builder {
	return BuilderPool.Get().(*Builder)
}

// *******************************
// process method
// *******************************

func Table(table string) *Builder {
	sql := newBuilder()
	sql.TableName = table
	sql.connection = DBInstance.clone()
	return sql
}

func (sql *Builder) Table(table string) *Builder {
	sql.TableName = table
	return sql
}

func (sql *Builder) Select(fields ...string) *Builder {
	sql.Fields = fields
	return sql
}

func (sql *Builder) OrderBy(filed string, order string) *Builder {
	sql.Order = "`" + filed + "` " + order
	return sql
}

func (sql *Builder) Skip(offset int) *Builder {
	sql.Offset = strconv.Itoa(offset)
	return sql
}

func (sql *Builder) Take(take int) *Builder {
	sql.Limit = strconv.Itoa(take)
	return sql
}

func (sql *Builder) Where(field string, operation string, arg interface{}) *Builder {
	sql.Wheres = append(sql.Wheres, dialect.Where{
		Field:     field,
		Operation: operation,
		Qmark:     "?",
	})
	sql.Args = append(sql.Args, arg)
	return sql
}

func (sql *Builder) WhereIn(field string, arg []interface{}) *Builder {
	if len(arg) == 0 {
		return sql
	}
	sql.Wheres = append(sql.Wheres, dialect.Where{
		Field:     field,
		Operation: "in",
		Qmark:     "(" + strings.Repeat("?,", len(arg)-1) + "?)",
	})
	sql.Args = append(sql.Args, arg...)
	return sql
}

func (sql *Builder) WhereNotIn(field string, arg []interface{}) *Builder {
	if len(arg) == 0 {
		return sql
	}
	sql.Wheres = append(sql.Wheres, dialect.Where{
		Field:     field,
		Operation: "not in",
		Qmark:     "(" + strings.Repeat("?,", len(arg)-1) + "?)",
	})
	sql.Args = append(sql.Args, arg...)
	return sql
}

func (sql *Builder) Find(arg interface{}) (map[string]interface{}, error) {
	return sql.Where("id", "=", arg).First()
}

func (sql *Builder) Count() (int64, error) {
	var (
		res map[string]interface{}
		err error
	)
	if res, err = sql.Select("count(*)").First(); err != nil {
		return 0, err
	}
	return res["count(*)"].(int64), nil
}

func (sql *Builder) WhereRaw(raw string, args ...interface{}) *Builder {
	sql.WhereRaws = raw
	sql.Args = append(sql.Args, args...)
	return sql
}

func (sql *Builder) UpdateRaw(raw string, args ...interface{}) *Builder {
	sql.UpdateRaws = append(sql.UpdateRaws, dialect.RawUpdate{
		Expression: raw,
		Args:       args,
	})
	return sql
}

func (sql *Builder) LeftJoin(table string, fieldA string, operation string, fieldB string) *Builder {
	sql.Leftjoins = append(sql.Leftjoins, dialect.Join{
		FieldA:    fieldA,
		FieldB:    fieldB,
		Table:     table,
		Operation: operation,
	})
	return sql
}

// *******************************
// terminal method
// -------------------------------
// sql args order:
// update ... => where ...
// *******************************

func (sql *Builder) First() (map[string]interface{}, error) {
	defer RecycleSql(sql)

	sql.connection.dialect.Select(&sql.SqlComponent)

	res := sql.connection.Query(sql.Statement, sql.Args...)

	if len(res) < 1 {
		return nil, errors.New("out of index")
	}
	return res[0], nil
}

func (sql *Builder) All() ([]map[string]interface{}, error) {
	defer RecycleSql(sql)

	sql.connection.dialect.Select(&sql.SqlComponent)

	res := sql.connection.Query(sql.Statement, sql.Args...)

	return res, nil
}

func (sql *Builder) ShowColumns() ([]map[string]interface{}, error) {
	defer RecycleSql(sql)

	res := sql.connection.Query(sql.connection.dialect.ShowColumns(sql.TableName))

	return res, nil
}

func (sql *Builder) ShowTables() ([]map[string]interface{}, error) {
	defer RecycleSql(sql)

	res := sql.connection.Query(sql.connection.dialect.ShowTables())

	return res, nil
}

func (sql *Builder) Update(values dialect.H) (int64, error) {
	defer RecycleSql(sql)

	sql.Values = values

	sql.connection.dialect.Update(&sql.SqlComponent)

	res := sql.connection.Exec(sql.Statement, sql.Args...)

	if affectRow, _ := res.RowsAffected(); affectRow < 1 {
		return 0, errors.New("no affect row")
	}

	return res.LastInsertId()
}

func (sql *Builder) Delete() error {
	defer RecycleSql(sql)

	sql.connection.dialect.Delete(&sql.SqlComponent)

	res := sql.connection.Exec(sql.Statement, sql.Args...)

	if affectRow, _ := res.RowsAffected(); affectRow < 1 {
		return errors.New("no affect row")
	}

	return nil
}

func (sql *Builder) Exec() (int64, error) {
	defer RecycleSql(sql)

	sql.connection.dialect.Update(&sql.SqlComponent)

	res := sql.connection.Exec(sql.Statement, sql.Args...)

	if affectRow, _ := res.RowsAffected(); affectRow < 1 {
		return 0, errors.New("no affect row")
	}

	return res.LastInsertId()
}

func (sql *Builder) Insert(values dialect.H) (int64, error) {
	defer RecycleSql(sql)

	sql.Values = values

	sql.connection.dialect.Insert(&sql.SqlComponent)

	res := sql.connection.Exec(sql.Statement, sql.Args...)

	if affectRow, _ := res.RowsAffected(); affectRow < 1 {
		return 0, errors.New("no affect row")
	}

	return res.LastInsertId()
}

// *******************************
// model method
// *******************************

func (sql *Builder) FormFirst(v interface{}) error {
	data, err := sql.First()

	if err != nil {
		return err
	}

	Transfer(data, v)

	return nil
}

func (sql *Builder) empty() *Builder {
	sql.Fields = make([]string, 0)
	sql.Args = make([]interface{}, 0)
	sql.TableName = ""
	sql.Wheres = make([]dialect.Where, 0)
	sql.Leftjoins = make([]dialect.Join, 0)
	return sql
}

func RecycleSql(sql *Builder) {
	sql.Fields = make([]string, 0)
	sql.TableName = ""
	sql.Wheres = make([]dialect.Where, 0)
	sql.Leftjoins = make([]dialect.Join, 0)
	sql.Args = make([]interface{}, 0)
	sql.Order = ""
	sql.Offset = ""
	sql.Limit = ""
	sql.WhereRaws = ""
	sql.UpdateRaws = make([]dialect.RawUpdate, 0)
	sql.Statement = ""

	BuilderPool.Put(sql)
}

func Transfer(sourceMap map[string]interface{}, targetStruct interface{}) error {
	p := reflect.ValueOf(targetStruct)

	if p.Kind() != reflect.Ptr || p.IsNil() {
		return errors.New("wrong type")
	}

	v := reflect.Indirect(p)
	t := v.Type()

	var (
		value interface{}
		ok    bool
	)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { // 判断是否为可导出字段
			value, ok = sourceMap[t.Field(i).Name]
			if !ok {
				value, ok = sourceMap[Lcfirst(t.Field(i).Name)]
				if ok {
					switch v.Field(i).Type().String() {
					case "string":
						v.Field(i).SetString(InterfaceToString(value))
					case "bool":
						v.Field(i).SetBool(InterfaceToBool(value))
					case "int":
						v.Field(i).SetInt(InterfaceToInt64(value))
					case "int64":
						v.Field(i).SetInt(InterfaceToInt64(value))
					case "int32":
						v.Field(i).SetInt(InterfaceToInt64(value))
					case "float32":
						v.Field(i).SetFloat(InterfaceToFloat64(value))
					case "float64":
						v.Field(i).SetFloat(InterfaceToFloat64(value))
					}
				}
			} else {
				return errors.New("wrong map key")
			}
		} else {
			return errors.New("wrong struct field")
		}
	}
	return nil
}

func InterfaceToInt64(value interface{}) int64 {

	if val, ok := value.(int64); ok {
		return val
	}

	if val, ok := value.(float64); ok {
		return int64(val)
	}

	if val, ok := value.(string); ok {
		if valInt, err := strconv.ParseInt(val, 10, 64); err == nil {
			return valInt
		}

		return 0
	}

	return 0
}

func InterfaceToFloat64(value interface{}) float64 {
	if val, ok := value.(float64); ok {
		return val
	}

	if val, ok := value.(string); ok {
		if valFloat64, err := strconv.ParseFloat(val, 64); err == nil {
			return valFloat64
		}
		return 0
	}

	return 0
}

func InterfaceToString(value interface{}) string {
	if val, ok := value.(string); ok {
		return val
	}

	return ""
}

func InterfaceToBool(value interface{}) bool {
	if val, ok := value.(bool); ok {
		return val
	}

	if val, ok := value.(string); ok {
		if val == "true" {
			return true
		} else {
			return false
		}
	}

	if val, ok := value.(int64); ok {
		if val > 0 {
			return true
		} else {
			return false
		}
	}

	return false
}

func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
