package fluentsql

type Flavor int

const (
	MySQL Flavor = iota
	PostgreSQL
	SQLite
	MongoDB
)

var (
	// DBType is the default flavor for all builders.
	DBType = PostgreSQL
)

// String returns the name of f.
func (f Flavor) String() string {
	switch f {
	case MySQL:
		return "MySQL"
	case PostgreSQL:
		return "PostgreSQL"
	case SQLite:
		return "SQLite"
	case MongoDB:
		return "MongoDB"
	}

	return "<Unknown>"
}
