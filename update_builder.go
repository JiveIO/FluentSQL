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

// NewUpdateBuilder Query builder constructor
func NewUpdateBuilder() *UpdateBuilder {
	return &UpdateBuilder{}
}

// ===========================================================================================================
//										Update Builder :: Operators
// ===========================================================================================================

func (u *UpdateBuilder) String() string {
	var queryParts []string

	queryParts = append(queryParts, u.updateStatement.String())
	queryParts = append(queryParts, u.setStatement.String())

	whereSql := u.whereStatement.String()
	if whereSql != "" {
		queryParts = append(queryParts, whereSql)
	}

	orderBySql := u.orderByStatement.String()
	if orderBySql != "" {
		queryParts = append(queryParts, orderBySql)
	}

	limitSql := u.limitStatement.String()
	if limitSql != "" {
		queryParts = append(queryParts, limitSql)
	}

	sql := strings.Join(queryParts, " ")

	return sql
}
