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
	}

	return "<Unknown>"
}

// p Get place holder format
func p(args []any) string {
	switch DBType {
	case MySQL:
		return Question
	case PostgreSQL:
		return fmt.Sprintf("%s%d", Dollar, len(args))
	case SQLite:
		return Question
	}

	return "#"
}
