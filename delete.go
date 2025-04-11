package fluentsql

import (
	"fmt"
	"strings"
)

// Delete clause
type Delete struct {
	Table any    // Table specifies the name of the table to delete data from
	Alias string // Alias represents an optional alias for the table
}

// String generates the DELETE SQL query as a string.
// It constructs the query by including the table name and an optional table alias.
//
// Returns:
//   - A string representing the DELETE SQL query.
func (u *Delete) String() string {
	var sb strings.Builder // sb is a string builder used to efficiently construct the query string
	sb.WriteString(fmt.Sprintf("DELETE FROM %s", u.Table))

	if u.Alias != "" {
		sb.WriteString(" " + u.Alias)
	}

	return sb.String()
}
