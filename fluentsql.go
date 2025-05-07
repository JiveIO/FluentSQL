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

// ====================================================================
// ========================== Declarations ============================
// ====================================================================

var (
	// Question is a PlaceholderFormat instance that leaves placeholders as
	// question marks.
	// Use for MySQL, SQLite
	question = "?"

	// Dollar is a PlaceholderFormat instance that replaces placeholders with
	// dollar-prefixed positional placeholders (e.g. $1, $2, $3).
	// Use for PostgreSQL, SQLite
	dollar = "$"

	// ------------------------- Dialects -------------------------

	// MySQL is a constant representing the MySQL database type.
	MySQL = "MySQL"
	// PostgreSQL is a constant representing the PostgreSQL database type.
	PostgreSQL = "PostgreSQL"
	// SQLite is a constant representing the SQLite database type.
	SQLite = "SQLite"

	// defaultDialect is the default dialect. It determines which SQL dialect to use for placeholder formatting.
	defaultDialect Dialect = new(PostgreSQLDialect)
)

// DefaultDialect returns the default dialect.
// This is for backward compatibility.
func DefaultDialect() Dialect {
	return defaultDialect
}

// SetDialect sets the current database dialect for placeholder formatting.
// Parameters:
//   - dialect (Dialect): The database dialect to set as the current one.
func SetDialect(dialect Dialect) {
	defaultDialect = dialect
}

// IsDialect checks if the current dialect matches the specified dialect name.
// Parameters:
//   - dialectName (string): The name of the dialect to check (e.g., "MySQL", "PostgreSQL", "SQLite")
//
// Returns:
//   - bool: true if the current dialect matches the specified name, false otherwise
func IsDialect(dialectName string) bool {
	return defaultDialect.Name() == dialectName
}

// ====================================================================
// ========================== MySQLDialect ============================
// ====================================================================

// MySQLDialect implements the Dialect interface for MySQL.
type MySQLDialect struct{}

// Name returns the name of the MySQL dialect.
func (d MySQLDialect) Name() string {
	return MySQL
}

// Placeholder returns the placeholder for MySQL, which is always "?".
//
// Parameter:
//   - position: Parameter position (not used in MySQL)
//
// Returns a string containing the question mark placeholder.
func (d MySQLDialect) Placeholder(_ int) string {
	return question
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
	return PostgreSQL
}

// Placeholder returns the placeholder for PostgreSQL, which is "$n" where n is the position.
//
// Parameter:
//   - position: The position of the placeholder (1-based)
//
// Returns a string containing the dollar-prefixed position (e.g. "$1", "$2", etc).
func (d PostgreSQLDialect) Placeholder(position int) string {
	return dollar + fmt.Sprintf("%d", position)
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
	return SQLite
}

// Placeholder returns the placeholder for SQLite, which is "?".
//
// Parameter:
//   - position: Parameter position (not used in SQLite)
//
// Returns a string containing the question mark placeholder.
func (d SQLiteDialect) Placeholder(_ int) string {
	return question
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

// ====================================================================
// ============================ Utilities =============================
// ====================================================================

// p generates the correct placeholder format based on the current database dialect.
// Parameters:
//   - args ([]any): A slice of arguments used to calculate the placeholder number for PostgreSQL.
//
// Output:
//   - (string): The placeholder-formatted string.
//
// Notes:
//   - MySQL and SQLite use question marks (?) for placeholders.
//   - PostgreSQL uses dollar-prefixed positional placeholders (e.g., $1, $2).
func p(args []any) string {
	return defaultDialect.Placeholder(len(args))
}
