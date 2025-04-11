package fluentsql

import "fmt"

type Flavor int

const (
	MySQL Flavor = iota
	PostgreSQL
	SQLite
)

var (
	// Question is a PlaceholderFormat instance that leaves placeholders as
	// question marks.
	// Use for MySQL, SQLite
	Question = "?"

	// Dollar is a PlaceholderFormat instance that replaces placeholders with
	// dollar-prefixed positional placeholders (e.g. $1, $2, $3).
	// Use for PostgreSQL, SQLite
	Dollar = "$"

	// Colon is a PlaceholderFormat instance that replaces placeholders with
	// colon-prefixed positional placeholders (e.g. :1, :2, :3).
	// Use for Oracle
	Colon = ":"

	// AtP is a PlaceholderFormat instance that replaces placeholders with
	// "@p"-prefixed positional placeholders (e.g. @p1, @p2, @p3).
	AtP = "@p"
)

var (
	// dbType is the default flavor for all builders. It determines which SQL flavor to use for placeholder formatting.
	dbType = PostgreSQL
)

// String returns the name of the Flavor.
// It provides a human-readable name corresponding to the Flavor value.
func (f Flavor) String() string {
	switch f {
	case MySQL:
		return "MySQL"
	case PostgreSQL:
		return "PostgreSQL"
	case SQLite:
		return "SQLite"
	}
	return "<Unknown>"
}

// DBType returns the current database flavor being used.
// Output:
//   - (Flavor): The current database flavor.
func DBType() Flavor {
	return dbType
}

// SetDBType sets the current database flavor for placeholder formatting.
// Parameters:
//   - flavor (Flavor): The database flavor to set as the current one.
func SetDBType(flavor Flavor) {
	dbType = flavor
}

// p generates the correct placeholder format based on the current database flavor.
// Parameters:
//   - args ([]any): A slice of arguments used to calculate the placeholder number for PostgreSQL.
//
// Output:
//   - (string): The placeholder formatted string.
//
// Notes:
//   - MySQL and SQLite use question marks (?) for placeholders.
//   - PostgreSQL uses dollar-prefixed positional placeholders (e.g., $1, $2).
func p(args []any) string {
	switch dbType {
	case MySQL:
		return Question
	case PostgreSQL:
		return fmt.Sprintf("%s%d", Dollar, len(args))
	case SQLite:
		return Question
	}
	return "#"
}
