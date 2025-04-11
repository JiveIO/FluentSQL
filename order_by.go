package fluentsql

import (
	"fmt"
	"strings"
)

// OrderByDir represents the sorting direction.
//
// Values:
// - Asc: Ascending order.
// - Desc: Descending order.
type OrderByDir int

// Constants representing sorting directions.
const (
	Asc  OrderByDir = iota // Ascending order.
	Desc                   // Descending order.
)

// SortItem defines a single field and its sorting direction for the ORDER BY clause.
//
// Fields:
// - Field (string): The name of the field to sort by.
// - Direction (OrderByDir): The direction of sorting (Asc or Desc).
type SortItem struct {
	Field     string     // The field to sort by.
	Direction OrderByDir // The direction of the sort (Asc or Desc).
}

// OrderBy represents the ORDER BY clause of a SQL query.
//
// Fields:
// - Items ([]SortItem): A slice of sorting items specifying fields and their respective sorting directions.
type OrderBy struct {
	Items []SortItem // List of sort items for constructing the ORDER BY clause.
}

// Dir returns the string representation of the sorting direction.
//
// Returns:
// - string: "ASC" if direction is ascending, "DESC" if direction is descending.
func (o *SortItem) Dir() string {
	var sign string // Holds the sorting direction string.

	switch o.Direction {
	case Asc:
		sign = "ASC"
	case Desc:
		sign = "DESC"
	}

	return sign
}

// Append adds a new field and its sorting direction to the ORDER BY clause.
//
// Parameters:
// - field string: The name of the field to add.
// - dir OrderByDir: The direction of sorting (Asc or Desc).
func (o *OrderBy) Append(field string, dir OrderByDir) {
	// Add new SortItem to the Items slice.
	o.Items = append(o.Items, SortItem{
		Field:     field,
		Direction: dir,
	})
}

// String generates the SQL ORDER BY clause.
//
// Returns:
// - string: The constructed ORDER BY clause. Returns an empty string if no fields are specified.
func (o *OrderBy) String() string {
	// Return empty string if no items are present.
	if len(o.Items) == 0 {
		return ""
	}

	var orderItems []string // Holds individual order by items in string format.
	for _, item := range o.Items {
		// Construct field-direction pair and append to the list.
		orderItems = append(orderItems, fmt.Sprintf("%s %s", item.Field, item.Dir()))
	}

	// Join all items and prefix with "ORDER BY".
	return fmt.Sprintf("ORDER BY %s", strings.Join(orderItems, ", "))
}
