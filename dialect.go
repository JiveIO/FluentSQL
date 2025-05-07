package fluentsql

import "fmt"

// ====================================================================
// =========================== Interfaces =============================
// ====================================================================

// Dialect defines the interface for database-specific SQL generation.
// Each supported database should implement this interface.
type Dialect interface {
	// Name returns the name of the dialect.
	Name() string

	// Placeholder generates a placeholder for a parameter at the given position.
	// For example, MySQL uses "?", PostgreSQL uses "$1", "$2", etc.
	Placeholder(position int) string

	// YearFunction returns the SQL function to extract the year from a date.
	// For example, MySQL uses "YEAR(?)", PostgreSQL uses "DATE_PART('year', ?)"
	YearFunction(field string) string
}

// DefaultDialect returns the default dialect based on the global dbType.
// This is for backward compatibility.
func DefaultDialect() Dialect {
	switch dbType {
	case MySQL:
		return MySQLDialect{}
	case PostgreSQL:
		return PostgreSQLDialect{}
	case SQLite:
		return SQLiteDialect{}
	default:
		return PostgreSQLDialect{} // Default to PostgreSQL
	}
}

// ====================================================================
// ========================== MySQLDialect ============================
// ====================================================================

// MySQLDialect implements the Dialect interface for MySQL.
type MySQLDialect struct{}

// Name returns the name of the MySQL dialect.
func (d MySQLDialect) Name() string {
	return "MySQL"
}

// Placeholder returns the placeholder for MySQL, which is always "?".
//
// Parameter:
//   - position: Parameter position (not used in MySQL)
//
// Returns a string containing the question mark placeholder.
func (d MySQLDialect) Placeholder(_ int) string {
	return Question
}

// YearFunction returns the MySQL-specific function to extract the year from a date.
//
// Parameter:
//   - field: The date field or expression to extract the year from
//
// Returns a string containing the MySQL YEAR function call.
func (d MySQLDialect) YearFunction(field string) string {
	return "YEAR(" + field + ")"
}

// ====================================================================
// ======================== PostgreSQLDialect =========================
// ====================================================================

// PostgreSQLDialect implements the Dialect interface for PostgreSQL.
type PostgreSQLDialect struct{}

// Name returns the name of the PostgreSQL dialect.
func (d PostgreSQLDialect) Name() string {
	return "PostgreSQL"
}

// Placeholder returns the placeholder for PostgreSQL, which is "$n" where n is the position.
//
// Parameter:
//   - position: The position of the placeholder (1-based)
//
// Returns a string containing the dollar-prefixed position (e.g. "$1", "$2", etc).
func (d PostgreSQLDialect) Placeholder(position int) string {
	return Dollar + fmt.Sprintf("%d", position)
}

// YearFunction returns the PostgreSQL-specific function to extract the year from a date.
//
// Parameter:
//   - field: The date field or expression to extract the year from
//
// Returns a string containing the PostgreSQL DATE_PART function call.
func (d PostgreSQLDialect) YearFunction(field string) string {
	return "DATE_PART('year', " + field + ")"
}

// ====================================================================
// ========================== SQLiteDialect ===========================
// ====================================================================

// SQLiteDialect implements the Dialect interface for SQLite.
type SQLiteDialect struct{}

// Name returns the name of the SQLite dialect.
func (d SQLiteDialect) Name() string {
	return "SQLite"
}

// Placeholder returns the placeholder for SQLite, which is "?".
//
// Parameter:
//   - position: Parameter position (not used in SQLite)
//
// Returns a string containing the question mark placeholder.
func (d SQLiteDialect) Placeholder(_ int) string {
	return Question
}

// YearFunction returns the SQLite-specific function to extract the year from a date.
//
// Parameter:
//   - field: The date field or expression to extract the year from
//
// Returns a string containing the SQLite strftime function call for year extraction.
func (d SQLiteDialect) YearFunction(field string) string {
	return "strftime('%Y', " + field + ")"
}
