package fluentsql

import "strings"

// ====================================================================
//                   Update Builder :: Structure
// ====================================================================

// UpdateBuilder struct
//
// UPDATE [LOW_PRIORITY] [IGNORE] table_reference
//
//	SET assignment_list
//	[WHERE where_condition]
//	[ORDER BY ...]
//	[LIMIT row_count]
//
// value:
//
//	{expr | DEFAULT}
//
// assignment:
//
//	col_name = value
//
// assignment_list:
//
//	assignment [, assignment] ...
type UpdateBuilder struct {
	// updateStatement represents the UPDATE clause of the SQL statement.
	updateStatement Update
	// setStatement represents the SET clause of the SQL statement.
	setStatement UpdateSet
	// whereStatement represents the WHERE clause of the SQL statement.
	whereStatement Where
	// orderByStatement represents the ORDER BY clause of the SQL statement.
	orderByStatement OrderBy
	// limitStatement represents the LIMIT clause of the SQL statement.
	limitStatement Limit
}

// UpdateInstance Update Builder constructor
//
// UpdateInstance creates and returns a new instance of UpdateBuilder.
//
// Returns:
// - *UpdateBuilder: A pointer to a new UpdateBuilder instance.
func UpdateInstance() *UpdateBuilder {
	return &UpdateBuilder{}
}

// ====================================================================
//                   Update Builder :: Operators
// ====================================================================

// String generates the SQL query string by concatenating all parts of the query.
// Returns:
// - A string containing the complete SQL query.
func (ub *UpdateBuilder) String() string {
	var queryParts []string // Holds different parts of the SQL query.

	// Add UPDATE clause to the query parts.
	// Add SET clause to the query parts.
	queryParts = append(queryParts,
		ub.updateStatement.String(),
		ub.setStatement.String(),
	)

	// Add WHERE clause if available.
	whereSql := ub.whereStatement.String()
	if whereSql != "" {
		queryParts = append(queryParts, whereSql)
	}

	// Add ORDER BY clause if available.
	orderBySql := ub.orderByStatement.String()
	if orderBySql != "" {
		queryParts = append(queryParts, orderBySql)
	}

	// Add LIMIT clause if available.
	limitSql := ub.limitStatement.String()
	if limitSql != "" {
		queryParts = append(queryParts, limitSql)
	}

	// Join all SQL parts into a single SQL query string.
	sql := strings.Join(queryParts, " ")

	return sql
}

// Update sets the table and optional alias for the UPDATE clause.
// Parameters:
// - table (any): The table to be updated.
// - alias (...string): An optional alias for the table.
// Returns:
// - *UpdateBuilder: The current UpdateBuilder instance.
func (ub *UpdateBuilder) Update(table any, alias ...string) *UpdateBuilder {
	ub.updateStatement.Table = table

	// Table alias
	if len(alias) > 0 {
		ub.updateStatement.Alias = alias[0]
	}

	return ub
}

// Set adds a key-value pair to the SET clause.
// Parameters:
// - field (any): The column to be updated.
// - value (any): The value to set for the column.
// Returns:
// - *UpdateBuilder: The current UpdateBuilder instance.
func (ub *UpdateBuilder) Set(field, value any) *UpdateBuilder {
	ub.setStatement.Append(field, value)

	return ub
}

// Where adds a condition to the WHERE clause using an AND operator.
// Parameters:
// - field (any): The field or column to evaluate.
// - opt (WhereOpt): The conditional operator (e.g., "=", ">", "<").
// - value (any): The value to compare to.
// Returns:
// - *UpdateBuilder: The current UpdateBuilder instance.
func (ub *UpdateBuilder) Where(field any, opt WhereOpt, value any) *UpdateBuilder {
	ub.whereStatement.Append(Condition{
		Field: field,
		Opt:   opt,
		Value: value,
		AndOr: And,
	})

	return ub
}

// WhereOr adds a condition to the WHERE clause using an OR operator.
// Parameters:
// - field (any): The field or column to evaluate.
// - opt (WhereOpt): The conditional operator (e.g., "=", ">", "<").
// - value (any): The value to compare to.
// Returns:
// - *UpdateBuilder: The current UpdateBuilder instance.
func (ub *UpdateBuilder) WhereOr(field any, opt WhereOpt, value any) *UpdateBuilder {
	ub.whereStatement.Append(Condition{
		Field: field,
		Opt:   opt,
		Value: value,
		AndOr: Or,
	})

	return ub
}

// WhereGroup combines multiple WHERE conditions into a grouped condition.
// Parameters:
// - groupCondition (FnWhereBuilder): A function that defines grouped conditions.
// Returns:
// - *UpdateBuilder: The current UpdateBuilder instance.
func (ub *UpdateBuilder) WhereGroup(groupCondition FnWhereBuilder) *UpdateBuilder {
	// Create a new WhereBuilder using the groupCondition function.
	whereBuilder := groupCondition(*WhereInstance())

	cond := Condition{
		Group: whereBuilder.whereStatement.Conditions,
	}

	ub.whereStatement.Conditions = append(ub.whereStatement.Conditions, cond)

	return ub
}

// WhereCondition appends multiple conditions directly into the WHERE clause.
// Parameters:
// - conditions (...Condition): A variadic list of conditions to append.
// Returns:
// - *UpdateBuilder: The current UpdateBuilder instance.
func (ub *UpdateBuilder) WhereCondition(conditions ...Condition) *UpdateBuilder {
	ub.whereStatement.Conditions = append(ub.whereStatement.Conditions, conditions...)

	return ub
}
