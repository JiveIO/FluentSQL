package fluentsql

import "strings"

// ===========================================================================================================
//										Update Builder :: Structure
// ===========================================================================================================

// UpdateBuilder struct
/*
UPDATE [LOW_PRIORITY] [IGNORE] table_reference
    SET assignment_list
    [WHERE where_condition]
    [ORDER BY ...]
    [LIMIT row_count]

value:
    {expr | DEFAULT}

assignment:
    col_name = value

assignment_list:
    assignment [, assignment] ...
*/
type UpdateBuilder struct {
	updateStatement  Update
	setStatement     UpdateSet
	whereStatement   Where
	orderByStatement OrderBy
	limitStatement   Limit
}

// UpdateInstance Update Builder constructor
func UpdateInstance() *UpdateBuilder {
	return &UpdateBuilder{}
}

// ===========================================================================================================
//										Update Builder :: Operators
// ===========================================================================================================

func (ub *UpdateBuilder) String() string {
	var queryParts []string

	queryParts = append(queryParts, ub.updateStatement.String())
	queryParts = append(queryParts, ub.setStatement.String())

	whereSql := ub.whereStatement.String()
	if whereSql != "" {
		queryParts = append(queryParts, whereSql)
	}

	orderBySql := ub.orderByStatement.String()
	if orderBySql != "" {
		queryParts = append(queryParts, orderBySql)
	}

	limitSql := ub.limitStatement.String()
	if limitSql != "" {
		queryParts = append(queryParts, limitSql)
	}

	sql := strings.Join(queryParts, " ")

	return sql
}

// Update builder
func (ub *UpdateBuilder) Update(table any, alias ...string) *UpdateBuilder {
	ub.updateStatement.Table = table

	// Table alias
	if len(alias) > 0 {
		ub.updateStatement.Alias = alias[0]
	}

	return ub
}

func (ub *UpdateBuilder) Set(field any, value any) *UpdateBuilder {
	ub.setStatement.Append(field, value)

	return ub
}

// Where builder
func (ub *UpdateBuilder) Where(Field any, Opt WhereOpt, Value any) *UpdateBuilder {
	ub.whereStatement.Append(Condition{
		Field: Field,
		Opt:   Opt,
		Value: Value,
		AndOr: And,
	})

	return ub
}
