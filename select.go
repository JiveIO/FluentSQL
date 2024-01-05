package fluentsql

import (
	"fmt"
	"strings"
)

// Select type struct
type Select struct {
	// Columns type string or a QueryBuilder
	Columns []any
}

func (s *Select) String() string {
	selectOf := "*"

	if len(s.Columns) > 0 {
		var columns []string

		for _, col := range s.Columns {
			if _, ok := col.(string); ok { // Column type string
				columns = append(columns, col.(string))
			} else if _, ok := col.(FieldYear); ok { // Column type FieldYear
				columns = append(columns, col.(FieldYear).String())
			} else if _, ok := col.(*QueryBuilder); ok { // Column type is QueryBuilder.
				selectQuery := col.(*QueryBuilder).String()

				if col.(*QueryBuilder).Query.Alias == "" {
					selectQuery = fmt.Sprintf("(%s)", selectQuery)
				}

				columns = append(columns, selectQuery)
			}
		}

		selectOf = strings.Join(columns, ", ")
	}

	return fmt.Sprintf("SELECT %s", selectOf)
}
