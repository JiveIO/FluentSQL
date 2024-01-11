package fluentsql

import (
	"fmt"
	"strings"
)

// ===========================================================================================================
//										Query Builder :: Structure
// ===========================================================================================================

// QueryBuilder struct
/*
SELECT
    [ALL | DISTINCT | DISTINCTROW ]
    [HIGH_PRIORITY]
    [STRAIGHT_JOIN]
    [SQL_SMALL_RESULT] [SQL_BIG_RESULT] [SQL_BUFFER_RESULT]
    [SQL_NO_CACHE] [SQL_CALC_FOUND_ROWS]
    select_expr [, select_expr] ...
    [into_option]
    [FROM table_references
      [PARTITION partition_list]]
    [WHERE where_condition]
    [GROUP BY {col_name | expr | position}, ... [WITH ROLLUP]]
    [HAVING where_condition]
    [WINDOW window_name AS (window_spec)
        [, window_name AS (window_spec)] ...]
    [ORDER BY {col_name | expr | position}
      [ASC | DESC], ... [WITH ROLLUP]]
    [LIMIT {[offset,] row_count | row_count OFFSET offset}]
    [into_option]
    [FOR {UPDATE | SHARE}
        [OF tbl_name [, tbl_name] ...]
        [NOWAIT | SKIP LOCKED]
      | LOCK IN SHARE MODE]
    [into_option]

into_option: {
    INTO OUTFILE 'file_name'
        [CHARACTER SET charset_name]
        export_options
  | INTO DUMPFILE 'file_name'
  | INTO var_name [, var_name] ...
}
*/
type QueryBuilder struct {
	alias            string // Query alias `AS <alias>
	selectStatement  Select
	fromStatement    From
	joinStatement    Join
	whereStatement   Where
	groupByStatement GroupBy
	havingStatement  Having // A version of Where
	orderByStatement OrderBy
	limitStatement   Limit
	fetchStatement   Fetch // A version of Limit
}

// QueryInstance Query builder constructor
func QueryInstance() *QueryBuilder {
	return &QueryBuilder{}
}

// ===========================================================================================================
//										Query Builder :: Operators
// ===========================================================================================================

// String convert query builder to string
func (qb *QueryBuilder) String() string {
	var queryParts []string

	queryParts = append(queryParts, qb.selectStatement.String())
	queryParts = append(queryParts, qb.fromStatement.String())

	joinSql := qb.joinStatement.String()
	if joinSql != "" {
		queryParts = append(queryParts, joinSql)
	}

	whereSql := qb.whereStatement.String()
	if whereSql != "" {
		queryParts = append(queryParts, whereSql)
	}

	groupSql := qb.groupByStatement.String()
	if groupSql != "" {
		queryParts = append(queryParts, groupSql)
	}

	havingSql := qb.havingStatement.String()
	if havingSql != "" {
		queryParts = append(queryParts, havingSql)
	}

	orderBySql := qb.orderByStatement.String()
	if orderBySql != "" {
		queryParts = append(queryParts, orderBySql)
	}

	limitSql := qb.limitStatement.String()
	if limitSql != "" {
		queryParts = append(queryParts, limitSql)
	}

	fetchSql := qb.fetchStatement.String()
	if fetchSql != "" {
		queryParts = append(queryParts, fetchSql)
	}

	sql := strings.Join(queryParts, " ")

	if qb.alias != "" {
		sql = fmt.Sprintf("(%s) AS %s",
			sql,
			qb.alias)
	}

	return sql
}

// Select builder
func (qb *QueryBuilder) Select(columns ...any) *QueryBuilder {
	qb.selectStatement.Columns = columns

	return qb
}

// From builder
func (qb *QueryBuilder) From(table any, alias ...string) *QueryBuilder {
	qb.fromStatement.Table = table

	// Table alias
	if len(alias) > 0 {
		qb.fromStatement.Alias = alias[0]
	}

	return qb
}

// Join builder
func (qb *QueryBuilder) Join(join JoinType, table string, condition Condition) *QueryBuilder {
	qb.joinStatement.Append(JoinItem{
		Join:      join,
		Table:     table,
		Condition: condition,
	})

	return qb
}

// Where builder
func (qb *QueryBuilder) Where(Field any, Opt WhereOpt, Value any) *QueryBuilder {
	qb.whereStatement.Append(Condition{
		Field: Field,
		Opt:   Opt,
		Value: Value,
		AndOr: And,
	})

	return qb
}

// Having builder
func (qb *QueryBuilder) Having(Field any, Opt WhereOpt, Value any) *QueryBuilder {
	qb.havingStatement.Append(Condition{
		Field: Field,
		Opt:   Opt,
		Value: Value,
		AndOr: And,
	})

	return qb
}

// WhereOr builder
func (qb *QueryBuilder) WhereOr(Field string, Opt WhereOpt, Value any) *QueryBuilder {
	qb.whereStatement.Append(Condition{
		Field: Field,
		Opt:   Opt,
		Value: Value,
		AndOr: Or,
	})

	return qb
}

type FnWhereGroupBuilder func(query QueryBuilder) *QueryBuilder

// WhereGroup combine multi where conditions into a group.
// Example: Group 2 condition created_at and update_at.
// SQL> SELECT * FROM users WHERE first_name LIKE '%john%' AND (created_at > '2024-01-12' OR update_at >= '2024-01-12') LIMIT 10 OFFSET 0
func (qb *QueryBuilder) WhereGroup(groupCondition FnWhereGroupBuilder) *QueryBuilder {
	// Create new WhereBuilder
	whereBuilder := groupCondition(*QueryInstance())

	cond := Condition{
		Group: whereBuilder.whereStatement.Conditions,
	}

	qb.whereStatement.Conditions = append(qb.whereStatement.Conditions, cond)

	return qb
}

// GroupBy fields in a query
func (qb *QueryBuilder) GroupBy(fields ...string) *QueryBuilder {
	qb.groupByStatement.Append(fields...)

	return qb
}

// OrderBy builder
func (qb *QueryBuilder) OrderBy(field string, dir OrderByDir) *QueryBuilder {
	qb.orderByStatement.Append(field, dir)

	return qb
}

// Limit builder
func (qb *QueryBuilder) Limit(Limit, Offset int) *QueryBuilder {
	qb.limitStatement.Limit = Limit
	qb.limitStatement.Offset = Offset

	return qb
}

// RemoveLimit builder
func (qb *QueryBuilder) RemoveLimit() Limit {
	var _limitStatement Limit

	_limitStatement.Limit = qb.limitStatement.Limit
	_limitStatement.Offset = qb.limitStatement.Offset

	qb.limitStatement.Limit = 0
	qb.limitStatement.Offset = 0

	return _limitStatement
}

// Fetch builder
func (qb *QueryBuilder) Fetch(Offset, Fetch int) *QueryBuilder {
	qb.fetchStatement.Offset = Offset
	qb.fetchStatement.Fetch = Fetch

	return qb
}

// RemoveFetch builder
func (qb *QueryBuilder) RemoveFetch() Fetch {
	var _fetchStatement Fetch

	_fetchStatement.Offset = qb.fetchStatement.Offset
	_fetchStatement.Fetch = qb.fetchStatement.Fetch

	qb.fetchStatement.Offset = 0
	qb.fetchStatement.Fetch = 0

	return _fetchStatement
}

// AS to create an alias of query builder,
//
// Examples:
// SELECT s.name, (SELECT COUNT(*) FROM product AS p WHERE p.store_id=s.id) AS counter FROM store AS s
// SELECT p.* FROM (SELECT first_name, last_name FROM Customers) AS p;
func (qb *QueryBuilder) AS(alias string) *QueryBuilder {
	qb.alias = alias

	return qb
}
