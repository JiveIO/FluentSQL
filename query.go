package fluentsql

import "fmt"

// Query type struct
type Query struct {
	Alias  string
	Select Select
	From   From
	Where  Where
	Limit  Limit
}

func (q Query) String() string {
	sql := fmt.Sprintf("%s %s %s %s", q.Select.String(), q.From.String(), q.Where.String(), q.Limit.String())

	if q.Alias != "" {
		sql = fmt.Sprintf("(%s %s %s %s) AS %s",
			q.Select.String(),
			q.From.String(),
			q.Where.String(),
			q.Limit.String(),
			q.Alias)
	}

	return sql
}
