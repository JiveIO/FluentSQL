package fluentsql

import (
	"fmt"
	"strings"
)

// Query type struct
type Query struct {
	Alias   string
	Select  Select
	From    From
	Where   Where
	OrderBy OrderBy
	Limit   Limit
}

func (q *Query) String() string {
	var query []string

	query = append(query, q.Select.String())
	query = append(query, q.From.String())

	whereSql := q.Where.String()
	if whereSql != "" {
		query = append(query, whereSql)
	}

	orderBySql := q.OrderBy.String()
	if orderBySql != "" {
		query = append(query, orderBySql)
	}

	limitSql := q.Limit.String()
	if limitSql != "" {
		query = append(query, limitSql)
	}

	sql := strings.Join(query, " ")

	if q.Alias != "" {
		sql = fmt.Sprintf("(%s) AS %s",
			sql,
			q.Alias)
	}

	return sql
}
