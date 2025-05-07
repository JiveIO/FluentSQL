package fluentsql

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
	// defaultDialect is the default dialect for all builders. It determines which SQL dialect to use for placeholder formatting.
	defaultDialect Dialect = new(PostgreSQLDialect)
)

// GetDialect returns the current database dialect being used.
// Output:
//   - (Dialect): The current database dialect.
func GetDialect() Dialect {
	return defaultDialect
}

// SetDialect sets the current database dialect for placeholder formatting.
// Parameters:
//   - dialect (Dialect): The database dialect to set as the current one.
func SetDialect(dialect Dialect) {
	defaultDialect = dialect
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
