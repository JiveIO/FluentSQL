package fluentsql

import (
	"fmt"
	"strings"
)

// GroupBy clause
type GroupBy struct {
	// Items stores the list of fields that will be grouped by in the query.
	Items []string
}

// Append adds one or more fields to the GroupBy clause.
//
// Parameters:
//   - field: One or more strings representing the fields to group by.
func (g *GroupBy) Append(field ...string) {
	g.Items = append(g.Items, field...)
}

// String converts the GroupBy clause to its SQL string representation.
//
// Returns:
//   - string: The SQL representation of the GroupBy clause. Returns an empty string if no fields are added.
func (g *GroupBy) String() string {
	if len(g.Items) == 0 {
		return ""
	}

	return fmt.Sprintf("GROUP BY %s", strings.Join(g.Items, ", "))
}
