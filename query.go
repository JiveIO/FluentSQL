package fluentsql

import (
	"fmt"
	"strings"
)

// Query type struct
type Query struct {
	Alias   string // Query alias `AS <alias>
	Select  Select
	From    From
	Join    Join
	Where   Where
	GroupBy GroupBy
	Having  Having // A version of Where
	OrderBy OrderBy
	Limit   Limit
	Fetch   Fetch // A version of Limit
}

func (q *Query) String() string {
	var query []string

	query = append(query, q.Select.String())
	query = append(query, q.From.String())

	joinSql := q.Join.String()
	if joinSql != "" {
		query = append(query, joinSql)
	}

	whereSql := q.Where.String()
	if whereSql != "" {
		query = append(query, whereSql)
	}

	groupSql := q.GroupBy.String()
	if groupSql != "" {
		query = append(query, groupSql)
	}

	havingSql := q.Having.String()
	if havingSql != "" {
		query = append(query, havingSql)
	}

	orderBySql := q.OrderBy.String()
	if orderBySql != "" {
		query = append(query, orderBySql)
	}

	limitSql := q.Limit.String()
	if limitSql != "" {
		query = append(query, limitSql)
	}

	fetchSql := q.Fetch.String()
	if fetchSql != "" {
		query = append(query, fetchSql)
	}

	sql := strings.Join(query, " ")

	if q.Alias != "" {
		sql = fmt.Sprintf("(%s) AS %s",
			sql,
			q.Alias)
	}

	return sql
}
