package fluentsql

import (
	"fmt"
	"strings"
)

// Having clause
type Having struct {
	Where
}

// String generates the SQL HAVING clause string based on the conditions provided.
// If there are no conditions, it returns an empty string.
//
// Returns:
//
//	string - The generated HAVING clause as a string.
func (w *Having) String() string {
	var conditions []string // Tracks individual condition strings.

	// Iterate through the provided conditions and format them into a single SQL string.
	if len(w.Conditions) > 0 {
		for _, cond := range w.Conditions {
			var _condition = cond.String() // Convert the condition to a string.

			// Check if the condition is combined with OR and there are existing conditions.
			if cond.AndOr == Or && len(conditions) > 0 {
				_orCondition := fmt.Sprint(" OR ", _condition) // Format the OR condition.

				last := len(conditions) - 1 // Index of the last condition.

				// Append the OR condition to the last condition.
				conditions[last] = conditions[last] + _orCondition
			} else {
				// Add the current condition to the condition list.
				conditions = append(conditions, _condition)
			}
		}
	}

	// If no conditions, return an empty string.
	if len(conditions) == 0 {
		return ""
	}

	// Combine the conditions with "AND" and wrap them in a "HAVING" clause.
	return fmt.Sprintf("HAVING %s", strings.Join(conditions, " AND "))
}
