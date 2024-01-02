package fluentsql

import (
	"fmt"
	"strings"
)

// From type struct
type From struct {
	Table any
	Alias string
}

func (f *From) String() string {
	var sb strings.Builder

	if _, ok := f.Table.(string); ok { // Table type string
		sb.WriteString(fmt.Sprintf("FROM %s", f.Table))

		if f.Alias != "" {
			sb.WriteString(" " + f.Alias)
		}
	} else if _, ok := f.Table.(*QueryBuilder); ok { // Table type is a complex query.
		// Example: SELECT Count(*) AS DistinctCountries FROM (SELECT DISTINCT Country FROM Customers);
		selectQuery := f.Table.(*QueryBuilder).String()

		sb.WriteString(fmt.Sprintf("FROM (%s)", selectQuery))
	}

	return sb.String()
}
