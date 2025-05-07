package fluentsql

import (
	"fmt"
	"strings"
)

type JoinType int

const (
	InnerJoin JoinType = iota
	LeftJoin
	RightJoin
	FullOuterJoin
	CrossJoin
)

// JoinItem represents a single join entry in a SQL statement.
// Fields:
//   - Join: The type of join (e.g., InnerJoin, LeftJoin).
//   - Table: The table name to join.
//   - Condition: The ON clause condition for the join.
type JoinItem struct {
	Join      JoinType
	Table     string
	Condition Condition
}

// opt returns the SQL join type as a string based on the JoinType.
// It is a method on JoinItem struct.
//
// Returns:
//   - string: The SQL join type ("INNER JOIN", "LEFT JOIN", etc.).
func (j *JoinItem) opt() string {
	var sign string

	switch j.Join {
	case InnerJoin:
		sign = "INNER JOIN"
	case LeftJoin:
		sign = "LEFT JOIN"
	case RightJoin:
		sign = "RIGHT JOIN"
	case FullOuterJoin:
		sign = "FULL OUTER JOIN"
	case CrossJoin:
		sign = "CROSS JOIN"
	}

	return sign
}

// Join represents a collection of join statements used in a SQL query.
// Fields:
//   - Items: A slice of JoinItem representing all join statements.
type Join struct {
	Items []JoinItem
}

// Append adds a new join item to the list of joins.
//
// Parameters:
//   - item (JoinItem): The join item to be appended.
func (j *Join) Append(item JoinItem) {
	j.Items = append(j.Items, item)
}

// String converts the Join object into a SQL-compatible join string.
//
// Returns:
//   - string: A SQL string representing the join clauses.
//     Returns an empty string if there are no join items.
func (j *Join) String() string {
	if len(j.Items) == 0 {
		return ""
	}

	var joinItems []string
	for _, item := range j.Items {
		joinStr := fmt.Sprintf("%s %s ON %s", item.opt(), item.Table, item.Condition.String())

		if item.Join == CrossJoin {
			joinStr = fmt.Sprintf("%s %s", item.opt(), item.Table)
		}

		joinItems = append(joinItems, joinStr)
	}

	return strings.Join(joinItems, " ")
}
