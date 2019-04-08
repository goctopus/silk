package silk

import (
	"github.com/chenhg5/silk/dialect"
)

type Dialect interface {
	// GetName get dialect's name
	GetName() string

	// ShowColumns show columns of specified table
	ShowColumns(table string) string

	// ShowTables show tables of database
	ShowTables() string

	// Insert
	Insert(comp *dialect.SqlComponent) string

	// Delete
	Delete(comp *dialect.SqlComponent) string

	// Update
	Update(comp *dialect.SqlComponent) string

	// Select
	Select(comp *dialect.SqlComponent) string
}

func GetDialectByDriver(driver string) Dialect {
	switch driver {
	case "mysql":
		return dialect.Mysql{}
	case "mssql":
		return dialect.Mssql{}
	case "postgresql":
		return dialect.Postgresql{}
	case "sqlite":
		return dialect.Sqlite{}
	default:
		return dialect.CommonDialect{}
	}
}
