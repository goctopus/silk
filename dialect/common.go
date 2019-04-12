package dialect

import "strings"

type CommonDialect struct {
}

func (CommonDialect) Insert(comp *SqlComponent) string {
	comp.prepareInsert()
	return comp.Statement
}

func (CommonDialect) Delete(comp *SqlComponent) string {
	comp.Statement = "delete from " + comp.TableName + comp.getWheres()
	return comp.Statement
}

func (CommonDialect) Update(comp *SqlComponent) string {
	comp.prepareUpdate()
	return comp.Statement
}

func (CommonDialect) Count(comp *SqlComponent) string {
	comp.prepareUpdate()
	return comp.Statement
}

func (CommonDialect) Select(comp *SqlComponent) string {
	comp.Statement = "select " + comp.getFields() + " from " + comp.TableName + comp.getJoins() + comp.getWheres() +
		comp.getOrderBy() + comp.getLimit() + comp.getOffset()
	return comp.Statement
}

func (CommonDialect) ShowColumns(table string) string {
	return "select column_name, udt_name from information_schema.columns where table_name = " + table
}

func (CommonDialect) GetName() string {
	return "common"
}

func (CommonDialect) ShowTables() string {
	return "show tables"
}

type H map[string]interface{}

type SqlComponent struct {
	Fields     []string
	TableName  string
	Wheres     []Where
	Leftjoins  []Join
	Args       []interface{}
	Order      string
	Offset     string
	Limit      string
	WhereRaws  string
	UpdateRaws []RawUpdate
	Statement  string
	Values     H
}

type Where struct {
	Operation string
	Field     string
	Qmark     string
}

type Join struct {
	Table     string
	FieldA    string
	Operation string
	FieldB    string
}

type RawUpdate struct {
	Expression string
	Args       []interface{}
}

// *******************************
// internal help function
// *******************************

func (sql *SqlComponent) getLimit() string {
	if sql.Limit == "" {
		return ""
	}
	return " limit " + sql.Limit + " "
}

func (sql *SqlComponent) getOffset() string {
	if sql.Offset == "" {
		return ""
	}
	return " offset " + sql.Offset + " "
}

func (sql *SqlComponent) getOrderBy() string {
	if sql.Order == "" {
		return ""
	}
	return " order by " + sql.Order + " "
}

func (sql *SqlComponent) getJoins() string {
	if len(sql.Leftjoins) == 0 {
		return ""
	}
	joins := ""
	for _, join := range sql.Leftjoins {
		joins += " left join " + join.Table + " on " + join.FieldA + " " + join.Operation + " " + join.FieldB + " "
	}
	return joins
}

func (sql *SqlComponent) getFields() string {
	if len(sql.Fields) == 0 {
		return "*"
	}
	if sql.Fields[0] == "count(*)" {
		return "count(*)"
	}
	fields := ""
	if len(sql.Leftjoins) == 0 {
		for _, field := range sql.Fields {
			fields += "`" + field + "`,"
		}
	} else {
		for _, field := range sql.Fields {
			arr := strings.Split(field, ".")
			if len(arr) > 1 {
				fields += arr[0] + ".`" + arr[1] + "`,"
			} else {
				fields += "`" + field + "`,"
			}
		}
	}
	return fields[:len(fields)-1]
}

func (sql *SqlComponent) getWheres() string {
	if len(sql.Wheres) == 0 {
		if sql.WhereRaws != "" {
			return " where " + sql.WhereRaws
		}
		return ""
	}
	wheres := " where "
	var arr []string
	for _, where := range sql.Wheres {
		arr = strings.Split(where.Field, ".")
		if len(arr) > 1 {
			wheres += arr[0] + ".`" + arr[1] + "` " + where.Operation + " " + where.Qmark + " and "
		} else {
			wheres += "`" + where.Field + "` " + where.Operation + " " + where.Qmark + " and "
		}
	}

	if sql.WhereRaws != "" {
		return wheres + sql.WhereRaws
	} else {
		return wheres[:len(wheres)-5]
	}
}

func (sql *SqlComponent) prepareUpdate() {
	fields := ""
	args := make([]interface{}, 0)

	if len(sql.Values) != 0 {

		for key, value := range sql.Values {
			fields += "`" + key + "` = ?, "
			args = append(args, value)
		}

		if len(sql.UpdateRaws) == 0 {
			fields = fields[:len(fields)-2]
		} else {
			for i := 0; i < len(sql.UpdateRaws); i++ {
				if i == len(sql.UpdateRaws)-1 {
					fields += sql.UpdateRaws[i].Expression + " "
				} else {
					fields += sql.UpdateRaws[i].Expression + ","
				}
				args = append(args, sql.UpdateRaws[i].Args...)
			}
		}

		sql.Args = append(args, sql.Args...)
	} else {
		if len(sql.UpdateRaws) == 0 {
			panic("prepareUpdate: wrong parameter")
		} else {
			for i := 0; i < len(sql.UpdateRaws); i++ {
				if i == len(sql.UpdateRaws)-1 {
					fields += sql.UpdateRaws[i].Expression + " "
				} else {
					fields += sql.UpdateRaws[i].Expression + ","
				}
				args = append(args, sql.UpdateRaws[i].Args...)
			}
		}
		sql.Args = append(args, sql.Args...)
	}

	sql.Statement = "update " + sql.TableName + " set " + fields + sql.getWheres()
}

func (sql *SqlComponent) prepareInsert() {
	fields := " ("
	quesMark := "("

	for key, value := range sql.Values {
		fields += "`" + key + "`,"
		quesMark += "?,"
		sql.Args = append(sql.Args, value)
	}
	fields = fields[:len(fields)-1] + ")"
	quesMark = quesMark[:len(quesMark)-1] + ")"

	sql.Statement = "insert into " + sql.TableName + fields + " values " + quesMark
}
