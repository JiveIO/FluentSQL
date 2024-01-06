package fluentsql

import "strings"

// ===========================================================================================================
//										Update structure
// ===========================================================================================================

// Edit Update statement
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
type Edit struct {
	Update  Update
	Set     UpdateSet
	Where   Where
	OrderBy OrderBy
	Limit   Limit
}

func (m *Edit) String() string {
	var query []string

	query = append(query, m.Update.String())
	query = append(query, m.Set.String())

	whereSql := m.Where.String()
	if whereSql != "" {
		query = append(query, whereSql)
	}

	orderBySql := m.OrderBy.String()
	if orderBySql != "" {
		query = append(query, orderBySql)
	}

	limitSql := m.Limit.String()
	if limitSql != "" {
		query = append(query, limitSql)
	}

	sql := strings.Join(query, " ")

	return sql
}

// ===========================================================================================================
//										Update Builder :: Structure
// ===========================================================================================================

type UpdateBuilder struct {
	Edit Edit
}

// NewUpdateBuilder Query builder constructor
func NewUpdateBuilder() *UpdateBuilder {
	return &UpdateBuilder{
		Edit: Edit{},
	}
}

// ===========================================================================================================
//										Update Builder :: Operators
// ===========================================================================================================
