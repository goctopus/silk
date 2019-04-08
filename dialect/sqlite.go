package dialect

type Sqlite struct {
	CommonDialect
}

func (Sqlite) GetName() string {
	return "sqlite"
}

func (Sqlite) ShowColumns(table string) string {
	return "PRAGMA table_info(" + table + ");"
}
