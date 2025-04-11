package fluentsql

import (
	"fmt"
	"strings"
)

// Select clause
type Select struct {
	// Columns type string or a QueryBuilder
	Columns []any
}

// String generates the SQL SELECT statement based on the columns provided in the Select struct.
// If no columns are specified, it defaults to "SELECT *".
//
// Returns:
// - A string representing the constructed SQL SELECT statement.
func (s *Select) String() string {
	// Default SELECT clause to "*"
	selectOf := "*"

	// Check if Columns is not empty
	if len(s.Columns) > 0 {
		var columns []string

		// Process each column in Columns
		for _, col := range s.Columns {
			if valueCase, ok := col.(*Case); ok { // Column is of type Case
				columns = append(columns, valueCase.String())
			} else if valueString, ok := col.(string); ok { // Column is a plain string
				columns = append(columns, valueString)
			} else if valueFieldYear, ok := col.(FieldYear); ok { // Column is of type FieldYear
				columns = append(columns, valueFieldYear.String())
			} else if valueQueryBuilder, ok := col.(*QueryBuilder); ok { // Column is a QueryBuilder
				selectQuery := valueQueryBuilder.String()

				// Add parentheses if alias is not provided
				if valueQueryBuilder.alias == "" {
					selectQuery = fmt.Sprintf("(%s)", valueQueryBuilder)
				}

				columns = append(columns, selectQuery)
			}
		}

		// Join all processed column representations with a comma
		selectOf = strings.Join(columns, ", ")
	}

	// Return the constructed SQL SELECT statement
	return fmt.Sprintf("SELECT %s", selectOf)
}
