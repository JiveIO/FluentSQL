package fluentsql

// WhereBuilder struct
type WhereBuilder struct {
	whereStatement Where // whereStatement holds the WHERE conditions of the query.
}

// WhereInstance Query builder constructor
// Returns:
//   - *WhereBuilder: A new instance of WhereBuilder.
func WhereInstance() *WhereBuilder {
	return &WhereBuilder{}
}

// Where builder
// Adds a new condition to the WHERE clause with an AND operator.
//
// Parameters:
//   - field (any): The field name or expression to be checked.
//   - opt (WhereOpt): The operator for the condition (e.g., Eq, Greater).
//   - value (any): The value to compare the field against.
//
// Returns:
//   - *WhereBuilder: The current instance of WhereBuilder for chaining.
func (wb *WhereBuilder) Where(field any, opt WhereOpt, value any) *WhereBuilder {
	wb.whereStatement.Append(Condition{
		Field: field,
		Opt:   opt,
		Value: value,
		AndOr: And,
	})

	return wb
}

// WhereOr builder
// Adds a new condition to the WHERE clause with an OR operator.
//
// Parameters:
//   - field (any): The field name or expression to be checked.
//   - opt (WhereOpt): The operator for the condition (e.g., Eq, Greater).
//   - value (any): The value to compare the field against.
//
// Returns:
//   - *WhereBuilder: The current instance of WhereBuilder for chaining.
func (wb *WhereBuilder) WhereOr(field any, opt WhereOpt, value any) *WhereBuilder {
	wb.whereStatement.Append(Condition{
		Field: field,
		Opt:   opt,
		Value: value,
		AndOr: Or,
	})

	return wb
}

// WhereGroup combines multiple WHERE conditions into a grouped condition.
//
// Parameters:
//   - groupCondition (FnWhereBuilder): Function defining the grouped conditions.
//
// Returns:
//   - *WhereBuilder: The current instance of WhereBuilder for chaining.
func (wb *WhereBuilder) WhereGroup(groupCondition FnWhereBuilder) *WhereBuilder {
	// Create new WhereBuilder
	whereBuilder := groupCondition(*WhereInstance())

	cond := Condition{
		Group: whereBuilder.whereStatement.Conditions,
	}

	wb.whereStatement.Conditions = append(wb.whereStatement.Conditions, cond)

	return wb
}

// WhereCondition appends multiple conditions to the WHERE clause.
//
// Parameters:
//   - conditions (...Condition): Multiple Condition objects to append.
//
// Returns:
//   - *WhereBuilder: The current instance of WhereBuilder for chaining.
func (wb *WhereBuilder) WhereCondition(conditions ...Condition) *WhereBuilder {
	wb.whereStatement.Conditions = append(wb.whereStatement.Conditions, conditions...)

	return wb
}

// FnWhereBuilder function type
// Used to group multiple conditions into a WhereBuilder.
//
// Parameters:
//   - whereBuilder (WhereBuilder): A WhereBuilder instance to modify.
//
// Returns:
//   - *WhereBuilder: The modified instance of WhereBuilder.
type FnWhereBuilder func(whereBuilder WhereBuilder) *WhereBuilder

// String constructs the WHERE clause as a string.
//
// Returns:
//   - string: The string representation of the WHERE clause.
func (wb *WhereBuilder) String() string {
	return wb.whereStatement.String()
}

// StringArgs constructs the WHERE clause as a string with argument placeholders.
//
// Parameters:
//   - args ([]any): A slice of arguments to use in the WHERE clause.
//
// Returns:
//   - string: The string representation of the WHERE clause with placeholders.
//   - []any: The updated slice of arguments.
func (wb *WhereBuilder) StringArgs(args []any) (string, []any) {
	return wb.whereStatement.StringArgs(args)
}

// Conditions retrieves all conditions of the WHERE clause.
//
// Returns:
//   - []Condition: A slice containing all conditions in the WHERE clause.
func (wb *WhereBuilder) Conditions() []Condition {
	return wb.whereStatement.Conditions
}
