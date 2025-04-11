package fluentsql

import "fmt"

// Fetch clause represents a SQL FETCH clause with offset and limit.
type Fetch struct {
	// Fetch specifies the number of rows to fetch.
	Fetch int
	// Offset specifies the number of rows to skip before starting to fetch rows.
	Offset int
}

// String generates the SQL FETCH clause as a string.
//
// If either Fetch or Offset is greater than 0, it returns the string in the format:
// "OFFSET <Offset> ROWS FETCH NEXT <Fetch> ROWS ONLY". Otherwise, it returns an empty string.
//
// Returns:
//   - A string representing the SQL FETCH clause.
func (f *Fetch) String() string {
	if f.Fetch > 0 || f.Offset > 0 {
		return fmt.Sprintf("OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", f.Offset, f.Fetch)
	}
	return ""
}
