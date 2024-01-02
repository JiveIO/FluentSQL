package fluentsql

import (
	"fmt"
	"strings"
)

// Select type struct
type Select struct {
	// Columns type string or Query
	Columns []any
}

func (s Select) String() string {
	selectOf := "*"

	if len(s.Columns) > 0 {
		var columns []string

		for _, col := range s.Columns {
			// Column type string
			if _, ok := col.(string); ok {
				columns = append(columns, col.(string))
				// Column type is complex query.
			} else if _, ok := col.(Query); ok {
				// Build Query struct to string
				// SELECT s.name, (SELECT COUNT(*) FROM product AS p WHERE p.store_id=s.id) FROM store AS s
				selectQuery := col.(Query).String()

				columns = append(columns, selectQuery)
			}
		}

		selectOf = strings.Join(columns, ", ")
	}

	return fmt.Sprintf("SELECT %s", selectOf)
}
