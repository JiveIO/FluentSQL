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
			} else if _, ok := col.(*QueryBuilder); ok { // Column type is a complex query.
				// Example: SELECT s.name, (SELECT COUNT(*) FROM product AS p WHERE p.store_id=s.id) AS count FROM store AS s
				selectQuery := col.(*QueryBuilder).String()

				columns = append(columns, selectQuery)
			}
		}

		selectOf = strings.Join(columns, ", ")
	}

	return fmt.Sprintf("SELECT %s", selectOf)
}
