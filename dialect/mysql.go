package dialect

type Mysql struct {
	CommonDialect
}

func (Mysql) GetName() string {
	return "mysql"
}

func (Mysql) ShowColumns(table string) string {
	return "show columns in " + table
}

func (Mysql) ShowTables() string {
	return "show tables"
}
