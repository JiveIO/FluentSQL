package fluentsql

import (
	"fmt"
	"strings"
)

// ===========================================================================================================
//										Query structure
// ===========================================================================================================

// Query statement
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

// ===========================================================================================================
//										Query Builder :: Structure
// ===========================================================================================================

type QueryBuilder struct {
	Query Query
}

// NewQueryBuilder Query builder constructor
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		Query: Query{},
	}
}

// ===========================================================================================================
//										Query Builder :: Operators
// ===========================================================================================================

// String convert query builder to string
func (qb *QueryBuilder) String() string {
	return qb.Query.String()
}

// Select builder
func (qb *QueryBuilder) Select(columns ...any) *QueryBuilder {
	qb.Query.Select.Columns = columns

	return qb
}

// From builder
func (qb *QueryBuilder) From(table any, alias ...string) *QueryBuilder {
	qb.Query.From.Table = table

	// Table alias
	if len(alias) > 0 {
		qb.Query.From.Alias = alias[0]
	}

	return qb
}

// Join builder
func (qb *QueryBuilder) Join(join JoinType, table string, condition Condition) *QueryBuilder {
	qb.Query.Join.Append(JoinItem{
		Join:      join,
		Table:     table,
		Condition: condition,
	})

	return qb
}

// Where builder
func (qb *QueryBuilder) Where(Field any, Opt WhereOpt, Value any) *QueryBuilder {
	qb.Query.Where.Append(Condition{
		Field: Field,
		Opt:   Opt,
		Value: Value,
		AndOr: And,
	})

	return qb
}

// Having builder
func (qb *QueryBuilder) Having(Field any, Opt WhereOpt, Value any) *QueryBuilder {
	qb.Query.Having.Append(Condition{
		Field: Field,
		Opt:   Opt,
		Value: Value,
		AndOr: And,
	})

	return qb
}

// WhereOr builder
func (qb *QueryBuilder) WhereOr(Field string, Opt WhereOpt, Value any) *QueryBuilder {
	qb.Query.Where.Append(Condition{
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
	whereBuilder := groupCondition(*NewQueryBuilder())

	cond := Condition{
		Group: whereBuilder.Query.Where.Conditions,
	}

	qb.Query.Where.Conditions = append(qb.Query.Where.Conditions, cond)

	return qb
}

// GroupBy fields in a query
func (qb *QueryBuilder) GroupBy(fields ...string) *QueryBuilder {
	qb.Query.GroupBy.Append(fields...)

	return qb
}

// OrderBy builder
func (qb *QueryBuilder) OrderBy(field string, dir OrderByDir) *QueryBuilder {
	qb.Query.OrderBy.Append(field, dir)

	return qb
}

// Limit builder
func (qb *QueryBuilder) Limit(Limit, Offset int) *QueryBuilder {
	qb.Query.Limit.Limit = Limit
	qb.Query.Limit.Offset = Offset

	return qb
}

// Fetch builder
func (qb *QueryBuilder) Fetch(Offset, Fetch int) *QueryBuilder {
	qb.Query.Fetch.Offset = Offset
	qb.Query.Fetch.Fetch = Fetch

	return qb
}

// AS to create an alias of query builder,
//
// Examples:
// SELECT s.name, (SELECT COUNT(*) FROM product AS p WHERE p.store_id=s.id) AS counter FROM store AS s
// SELECT p.* FROM (SELECT first_name, last_name FROM Customers) AS p;
func (qb *QueryBuilder) AS(alias string) *QueryBuilder {
	qb.Query.Alias = alias

	return qb
}
