package dialect

type Postgresql struct {
	CommonDialect
}

func (Postgresql) GetName() string {
	return "postgresql"
}

func (Postgresql) ShowTables() string {
	return "select tablename from pg_catalog.pg_tables"
}
