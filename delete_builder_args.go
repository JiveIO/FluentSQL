package fluentsql

import (
	"fmt"
	"strings"
)

// Sql generates the DELETE SQL query string and returns it along with its arguments.
//
// Returns:
//   - string: The DELETE SQL query string.
//   - []any: A slice of arguments used in the query.
//   - error: Any error that occurs during query generation.
func (db *DeleteBuilder) Sql() (string, []any, error) {
	var args []any // A slice to hold query arguments.

	return db.StringArgs(args)
}

// StringArgs constructs and returns the DELETE SQL query string along with its arguments.
//
// Parameters:
//   - args ([]any): A slice of arguments passed to be used in the query.
//
// Returns:
//   - string: The DELETE SQL query string built from various components.
//   - []any: A slice of any type containing the arguments used in the query.
//   - error: Any error that may occur during the query construction.
func (db *DeleteBuilder) StringArgs(args []any) (string, []any, error) {
	var queryParts []string // A slice to gather all query parts (e.g., DELETE, WHERE, etc.).
	var sqlStr string       // Holds the current query string component.

	// Add the DELETE statement and arguments.
	sqlStr, args = db.deleteStatement.StringArgs(args)
	queryParts = append(queryParts, sqlStr)

	// Add the WHERE clause if present.
	sqlStr, args = db.whereStatement.StringArgs(args)
	if sqlStr != "" {
		queryParts = append(queryParts, sqlStr)
	}

	// Add the ORDER BY clause if present.
	sqlStr, args = db.orderByStatement.StringArgs(args)
	if sqlStr != "" {
		queryParts = append(queryParts, sqlStr)
	}

	// Add the LIMIT clause if present.
	sqlStr, args = db.limitStatement.StringArgs(args)
	if sqlStr != "" {
		queryParts = append(queryParts, sqlStr)
	}

	// Combine all parts into a single SQL string.
	sql := strings.Join(queryParts, " ")

	return sql, args, nil
}

// StringArgs generates the DELETE SQL statement as a string and updates the provided arguments.
//
// Parameters:
//   - args ([]any): A slice of arguments passed to be used in the query.
//
// Returns:
//   - string: The DELETE SQL statement including the table and alias (if present).
//   - []any: The updated slice of query arguments.
func (u *Delete) StringArgs(args []any) (string, []any) {
	var sb strings.Builder                                 // A strings.Builder to construct the query string efficiently.
	sb.WriteString(fmt.Sprintf("DELETE FROM %s", u.Table)) // Add the table to the DELETE statement.

	// Add alias to the DELETE statement if it's not empty.
	if u.Alias != "" {
		sb.WriteString(" " + u.Alias)
	}

	return sb.String(), args
}
