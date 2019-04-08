package dialect

type Mssql struct {
	CommonDialect
}

func (Mssql) GetName() string {
	return "mssql"
}