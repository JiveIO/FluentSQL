package fluentsql

import "strings"

// ====================================================================
//                   Delete Builder :: Structure
// ====================================================================

// DeleteBuilder struct represents a builder for constructing DELETE SQL queries.
//
// Syntax:
//
//	DELETE [LOW_PRIORITY] [QUICK] [IGNORE] FROM tbl_name [[AS] tbl_alias]
//	[PARTITION (partition_name [, partition_name] ...)]
//	[WHERE where_condition]
//	[ORDER BY ...]
//	[LIMIT row_count]
//
// It defines the components of the DELETE query.
type DeleteBuilder struct {
	deleteStatement  Delete  // Defines the DELETE clause for specifying the table and optional alias
	whereStatement   Where   // Stores conditions for the WHERE clause
	orderByStatement OrderBy // Represents sorting conditions for the ORDER BY clause
	limitStatement   Limit   // Specifies the LIMIT and OFFSET for the query
}

// DeleteInstance creates a new instance of DeleteBuilder.
//
// Returns:
//   - *DeleteBuilder: A pointer to the newly created DeleteBuilder instance.
func DeleteInstance() *DeleteBuilder {
	return &DeleteBuilder{}
}

// ====================================================================
//                   Delete Builder :: Operators
// ====================================================================

// String generates the DELETE SQL query as a string.
//
// It constructs the query by combining various parts of the query,
// including the DELETE clause, WHERE clause, ORDER BY clause, and LIMIT clause.
//
// Returns:
//   - A string representing the complete DELETE SQL query.
func (db *DeleteBuilder) String() string {
	var queryParts []string

	// Add the DELETE statement
	queryParts = append(queryParts, db.deleteStatement.String())

	// Add the WHERE clause if present
	whereSql := db.whereStatement.String()
	if whereSql != "" {
		queryParts = append(queryParts, whereSql)
	}

	// Add the ORDER BY clause if present
	orderBySql := db.orderByStatement.String()
	if orderBySql != "" {
		queryParts = append(queryParts, orderBySql)
	}

	// Add the LIMIT clause if present
	limitSql := db.limitStatement.String()
	if limitSql != "" {
		queryParts = append(queryParts, limitSql)
	}

	// Combine all parts into a single SQL string
	sql := strings.Join(queryParts, " ")

	return sql
}

// Delete specifies the table and an optional alias for the DELETE query.
//
// Parameters:
//   - table (string): The name of the table from which rows will be deleted.
//   - alias (...string): An optional alias for the table.
//
// Returns:
//   - *DeleteBuilder: A pointer to the current instance of DeleteBuilder.
func (db *DeleteBuilder) Delete(table string, alias ...string) *DeleteBuilder {
	db.deleteStatement.Table = table

	if len(alias) > 0 {
		db.deleteStatement.Alias = alias[0]
	}

	return db
}

// Where adds a condition to the WHERE clause using the AND operator.
//
// Parameters:
//   - field (any): The field or column involved in the condition.
//   - opt (WhereOpt): The comparison operator (e.g., =, !=, >, <).
//   - value (any): The value to compare against.
//
// Returns:
//   - *DeleteBuilder: A pointer to the current instance of DeleteBuilder.
func (db *DeleteBuilder) Where(field any, opt WhereOpt, value any) *DeleteBuilder {
	db.whereStatement.Append(Condition{
		Field: field,
		Opt:   opt,
		Value: value,
		AndOr: And,
	})

	return db
}

// WhereOr adds a condition to the WHERE clause using the OR operator.
//
// Parameters:
//   - field (any): The field or column involved in the condition.
//   - opt (WhereOpt): The comparison operator (e.g., =, !=, >, <).
//   - value (any): The value to compare against.
//
// Returns:
//   - *DeleteBuilder: A pointer to the current instance of DeleteBuilder.
func (db *DeleteBuilder) WhereOr(field any, opt WhereOpt, value any) *DeleteBuilder {
	db.whereStatement.Append(Condition{
		Field: field,
		Opt:   opt,
		Value: value,
		AndOr: Or,
	})

	return db
}

// WhereGroup combines multiple conditions into a grouped WHERE clause.
//
// Parameters:
//   - groupCondition (FnWhereBuilder): A function that defines the grouped conditions.
//
// Returns:
//   - *DeleteBuilder: A pointer to the current instance of DeleteBuilder.
func (db *DeleteBuilder) WhereGroup(groupCondition FnWhereBuilder) *DeleteBuilder {
	// Create new WhereBuilder
	whereBuilder := groupCondition(*WhereInstance())

	cond := Condition{
		Group: whereBuilder.whereStatement.Conditions,
	}

	db.whereStatement.Conditions = append(db.whereStatement.Conditions, cond)

	return db
}

// WhereCondition appends multiple conditions to the WHERE clause.
//
// Parameters:
//   - conditions (...Condition): A variadic parameter of conditions to be added.
//
// Returns:
//   - *DeleteBuilder: A pointer to the current instance of DeleteBuilder.
func (db *DeleteBuilder) WhereCondition(conditions ...Condition) *DeleteBuilder {
	db.whereStatement.Conditions = append(db.whereStatement.Conditions, conditions...)

	return db
}
