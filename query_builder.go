package fluentsql

import (
	"fmt"
	"strings"
)

// ====================================================================
//                   Query Builder :: Structure
// ====================================================================

// QueryBuilder struct
// Syntax:
//
// SELECT
//
//	[ALL | DISTINCT | DISTINCTROW ]
//	[HIGH_PRIORITY]
//	[STRAIGHT_JOIN]
//	[SQL_SMALL_RESULT] [SQL_BIG_RESULT] [SQL_BUFFER_RESULT]
//	[SQL_NO_CACHE] [SQL_CALC_FOUND_ROWS]
//	select_expr [, select_expr] ...
//	[into_option]
//	[FROM table_references
//	  [PARTITION partition_list]]
//	[WHERE where_condition]
//	[GROUP BY {col_name | expr | position}, ... [WITH ROLLUP]]
//	[HAVING where_condition]
//	[WINDOW window_name AS (window_spec)
//	    [, window_name AS (window_spec)] ...]
//	[ORDER BY {col_name | expr | position}
//	  [ASC | DESC], ... [WITH ROLLUP]]
//	[LIMIT {[offset,] row_count | row_count OFFSET offset}]
//	[into_option]
//	[FOR {UPDATE | SHARE}
//	    [OF tbl_name [, tbl_name] ...]
//	    [NOWAIT | SKIP LOCKED]
//	  | LOCK IN SHARE MODE]
//	[into_option]
//
//	into_option: {
//	   INTO OUTFILE 'file_name'
//	       [CHARACTER SET charset_name]
//	       export_options
//	 | INTO DUMPFILE 'file_name'
//	 | INTO var_name [, var_name] ...
//	}
type QueryBuilder struct {
	// alias defines an optional alias for the query.
	alias string

	// selectStatement represents the SELECT clause of the query.
	selectStatement Select

	// fromStatement represents the FROM clause of the query.
	fromStatement From

	// joinStatement represents the JOIN clauses of the query.
	joinStatement Join

	// whereStatement represents the WHERE clause of the query.
	whereStatement Where

	// groupByStatement represents the GROUP BY clause of the query.
	groupByStatement GroupBy

	// havingStatement represents the HAVING clause of the query.
	havingStatement Having

	// orderByStatement represents the ORDER BY clause of the query.
	orderByStatement OrderBy

	// limitStatement represents the LIMIT clause of the query.
	limitStatement Limit

	// fetchStatement represents a FETCH clause, an alternative to LIMIT.
	fetchStatement Fetch
}

// QueryInstance creates and returns a new instance of QueryBuilder.
//
// Returns:
// - *QueryBuilder: A pointer to a new QueryBuilder instance.
func QueryInstance() *QueryBuilder {
	return &QueryBuilder{}
}

// ====================================================================
//                   Query Builder :: Operators
// ====================================================================

// String converts the QueryBuilder instance to a SQL query string.
//
// Returns:
// - string: The SQL query string representation of the QueryBuilder.
func (qb *QueryBuilder) String() string {
	var queryParts []string

	// Append SELECT clause
	// Append FROM clause
	queryParts = append(queryParts,
		qb.selectStatement.String(),
		qb.fromStatement.String(),
	)

	// Append JOIN clauses
	joinSql := qb.joinStatement.String()
	if joinSql != "" {
		queryParts = append(queryParts, joinSql)
	}

	// Append WHERE clause
	whereSql := qb.whereStatement.String()
	if whereSql != "" {
		queryParts = append(queryParts, whereSql)
	}

	// Append GROUP BY clause
	groupSql := qb.groupByStatement.String()
	if groupSql != "" {
		queryParts = append(queryParts, groupSql)
	}

	// Append HAVING clause
	havingSql := qb.havingStatement.String()
	if havingSql != "" {
		queryParts = append(queryParts, havingSql)
	}

	// Append ORDER BY clause
	orderBySql := qb.orderByStatement.String()
	if orderBySql != "" {
		queryParts = append(queryParts, orderBySql)
	}

	// Append LIMIT clause
	limitSql := qb.limitStatement.String()
	if limitSql != "" {
		queryParts = append(queryParts, limitSql)
	}

	// Append FETCH clause
	fetchSql := qb.fetchStatement.String()
	if fetchSql != "" {
		queryParts = append(queryParts, fetchSql)
	}

	// Join all parts with a space
	sql := strings.Join(queryParts, " ")

	// Add alias if provided
	if qb.alias != "" {
		sql = fmt.Sprintf("(%s) AS %s", sql, qb.alias)
	}

	return sql
}

// Select defines the SELECT clause of the query.
//
// Parameters:
// - columns ...any: The columns to be selected in the query.
//
// Returns:
// - *QueryBuilder: The QueryBuilder instance with updated SELECT clause.
func (qb *QueryBuilder) Select(columns ...any) *QueryBuilder {
	qb.selectStatement.Columns = columns
	return qb
}

// From defines the FROM clause in the query.
//
// Parameters:
// - table any: The table from which to select data.
// - alias ...string: Optional alias for the table.
//
// Returns:
// - *QueryBuilder: The QueryBuilder instance with updated FROM clause.
func (qb *QueryBuilder) From(table any, alias ...string) *QueryBuilder {
	qb.fromStatement.Table = table

	// Table alias
	if len(alias) > 0 {
		qb.fromStatement.Alias = alias[0]
	}

	return qb
}

// Join adds a join clause to the query.
//
// Parameters:
// - join JoinType: The type of join (e.g., INNER JOIN, LEFT JOIN).
// - table string: The table to join.
// - condition Condition: The ON condition for the join.
//
// Returns:
// - *QueryBuilder: The QueryBuilder instance with the added JOIN clause.
func (qb *QueryBuilder) Join(join JoinType, table string, condition Condition) *QueryBuilder {
	qb.joinStatement.Append(JoinItem{
		Join:      join,
		Table:     table,
		Condition: condition,
	})
	return qb
}

// Having defines the HAVING clause of the query.
//
// Parameters:
// - field any: The column or expression to apply the condition to.
// - opt WhereOpt: The operator for the condition (e.g., =, >, <).
// - value any: The value to compare against.
//
// Returns:
// - *QueryBuilder: The QueryBuilder instance with updated HAVING clause.
func (qb *QueryBuilder) Having(field any, opt WhereOpt, value any) *QueryBuilder {
	qb.havingStatement.Append(Condition{
		Field: field,
		Opt:   opt,
		Value: value,
		AndOr: And,
	})
	return qb
}

// Where adds a condition to the WHERE clause of the query.
//
// Parameters:
// - field any: The column or expression to apply the condition to.
// - opt WhereOpt: The operator for the condition (e.g., =, >, <).
// - value any: The value to compare against.
//
// Returns:
// - *QueryBuilder: The QueryBuilder instance with updated WHERE clause.
func (qb *QueryBuilder) Where(field any, opt WhereOpt, value any) *QueryBuilder {
	qb.whereStatement.Append(Condition{
		Field: field,
		Opt:   opt,
		Value: value,
		AndOr: And,
	})
	return qb
}

// WhereOr adds a condition to the WHERE clause with OR logic.
//
// Parameters:
// - field any: The column or expression to apply the condition to.
// - opt WhereOpt: The operator for the condition (e.g., =, >, <).
// - value any: The value to compare against.
//
// Returns:
// - *QueryBuilder: The QueryBuilder instance with updated WHERE clause.
func (qb *QueryBuilder) WhereOr(field any, opt WhereOpt, value any) *QueryBuilder {
	qb.whereStatement.Append(Condition{
		Field: field,
		Opt:   opt,
		Value: value,
		AndOr: Or,
	})
	return qb
}

// WhereGroup combines multiple conditions into a grouped WHERE clause.
//
// Parameters:
// - groupCondition FnWhereBuilder: A function that constructs a group of conditions.
//
// Returns:
// - *QueryBuilder: The QueryBuilder instance with the grouped WHERE clause.
func (qb *QueryBuilder) WhereGroup(groupCondition FnWhereBuilder) *QueryBuilder {
	// Create new WhereBuilder
	whereBuilder := groupCondition(*WhereInstance())

	cond := Condition{
		Group: whereBuilder.whereStatement.Conditions,
	}

	qb.whereStatement.Conditions = append(qb.whereStatement.Conditions, cond)

	return qb
}

// WhereCondition appends multiple conditions to the WHERE clause.
//
// Parameters:
// - conditions ...Condition: Conditions to be added.
//
// Returns:
// - *QueryBuilder: The QueryBuilder instance with updated WHERE clause.
func (qb *QueryBuilder) WhereCondition(conditions ...Condition) *QueryBuilder {
	qb.whereStatement.Conditions = append(qb.whereStatement.Conditions, conditions...)
	return qb
}

// GroupBy defines the GROUP BY clause of the query.
//
// Parameters:
// - fields ...string: The fields to group by.
//
// Returns:
// - *QueryBuilder: The QueryBuilder instance with updated GROUP BY clause.
func (qb *QueryBuilder) GroupBy(fields ...string) *QueryBuilder {
	qb.groupByStatement.Append(fields...)
	return qb
}

// OrderBy defines the ORDER BY clause of the query.
//
// Parameters:
// - field string: The field to sort by.
// - dir OrderByDir: The direction of sorting (ASC or DESC).
//
// Returns:
// - *QueryBuilder: The QueryBuilder instance with updated ORDER BY clause.
func (qb *QueryBuilder) OrderBy(field string, dir OrderByDir) *QueryBuilder {
	qb.orderByStatement.Append(field, dir)
	return qb
}

// Limit sets the LIMIT clause of the query.
//
// Parameters:
// - limit int: The maximum number of rows to return.
// - offset int: The number of rows to skip.
//
// Returns:
// - *QueryBuilder: The QueryBuilder instance with updated LIMIT clause.
func (qb *QueryBuilder) Limit(limit, offset int) *QueryBuilder {
	qb.limitStatement.Limit = limit
	qb.limitStatement.Offset = offset
	return qb
}

// RemoveLimit removes the LIMIT clause from the query.
//
// Returns:
// - Limit: The removed LIMIT clause.
func (qb *QueryBuilder) RemoveLimit() Limit {
	var _limitStatement Limit

	_limitStatement.Limit = qb.limitStatement.Limit
	_limitStatement.Offset = qb.limitStatement.Offset

	qb.limitStatement.Limit = 0
	qb.limitStatement.Offset = 0

	return _limitStatement
}

// Fetch sets the FETCH clause of the query.
//
// Parameters:
// - offset int: The number of rows to skip.
// - fetch int: The number of rows to fetch.
//
// Returns:
// - *QueryBuilder: The QueryBuilder instance with updated FETCH clause.
func (qb *QueryBuilder) Fetch(offset, fetch int) *QueryBuilder {
	qb.fetchStatement.Offset = offset
	qb.fetchStatement.Fetch = fetch
	return qb
}

// RemoveFetch removes the FETCH clause from the query.
//
// Returns:
// - Fetch: The removed FETCH clause.
func (qb *QueryBuilder) RemoveFetch() Fetch {
	var _fetchStatement Fetch

	_fetchStatement.Offset = qb.fetchStatement.Offset
	_fetchStatement.Fetch = qb.fetchStatement.Fetch

	qb.fetchStatement.Offset = 0
	qb.fetchStatement.Fetch = 0

	return _fetchStatement
}

// AS sets an alias for the entire QueryBuilder instance.
//
// Parameters:
// - alias string: The alias to be used.
//
// Returns:
// - *QueryBuilder: The QueryBuilder instance with an alias.
//
// Examples:
//
//	SELECT s.name, (SELECT COUNT(*) FROM product AS p WHERE p.store_id=s.id) AS counter FROM store AS s
//	SELECT p.* FROM (SELECT first_name, last_name FROM Customers) AS p;
func (qb *QueryBuilder) AS(alias string) *QueryBuilder {
	qb.alias = alias
	return qb
}
