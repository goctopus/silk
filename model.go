package silk

type Model struct {
	// db conn
	DB *Sql

	// table
	Table string
}

func (m *Model) Clean() {
	m.DB = Table(m.Table)
}

func (m *Model) Where(field string, op string, value interface{}) *Sql {
	return m.DB.Where(field, op, value)
}