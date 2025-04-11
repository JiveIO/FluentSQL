package fluentsql

import "fmt"

// Limit clause
type Limit struct {
	Limit  int // Limit specifies the maximum number of rows to return.
	Offset int // Offset specifies the starting point for rows to return.
}

// String generates the SQL LIMIT and OFFSET clause string.
// It returns an empty string if both Limit and Offset are zero.
//
// Returns:
// - string: The SQL LIMIT and OFFSET clause string.
func (l *Limit) String() string {
	// Check if Limit or Offset is greater than zero.
	if l.Limit > 0 || l.Offset > 0 {
		// Return formatted LIMIT and OFFSET clause.
		return fmt.Sprintf("LIMIT %d OFFSET %d", l.Limit, l.Offset)
	}

	// Return an empty string if no limit or offset is set.
	return ""
}
