package fluentsql

import (
	"fmt"
	"strings"
)

// From clause
type From struct {
	// Table represents the table name or a nested query. It can be of type string or *QueryBuilder.
	Table any
	// Alias defines an alias for the table or query in the SQL statement.
	Alias string
}

// String generates the SQL representation of the "FROM" clause.
// It returns the constructed SQL string for the "FROM" clause.
//
// The Table field supports either a simple table name (string)
// or a nested query (using a *QueryBuilder). An optional Alias
// can also be appended to the clause.
func (f *From) String() string {
	var sb strings.Builder

	// Check if Table is of type string
	if _, ok := f.Table.(string); ok {
		sb.WriteString(fmt.Sprintf("FROM %s", f.Table))
	} else if _, ok := f.Table.(*QueryBuilder); ok { // Check if Table is of type *QueryBuilder
		selectQuery := f.Table.(*QueryBuilder).String()

		// If the QueryBuilder has no alias, wrap the query in parentheses
		if f.Table.(*QueryBuilder).alias == "" {
			sb.WriteString(fmt.Sprintf("FROM (%s)", selectQuery))
		} else {
			sb.WriteString(fmt.Sprintf("FROM %s", selectQuery))
		}
	}

	// Append the alias if it is not empty
	if f.Alias != "" {
		sb.WriteString(" " + f.Alias)
	}

	return sb.String()
}
