package fluentsql

import (
	"strings"
)

// ====================================================================
//                   Insert Builder :: Structure
// ====================================================================

// InsertBuilder struct

/*
INSERT [LOW_PRIORITY | DELAYED | HIGH_PRIORITY] [IGNORE]
    [INTO] tbl_name
    [PARTITION (partition_name [, partition_name] ...)]
    [(col_name [, col_name] ...)]
    { {VALUES | VALUE} (value_list) [, (value_list)] ... }
    [AS row_alias[(col_alias [, col_alias] ...)]]
    [ON DUPLICATE KEY UPDATE assignment_list]

INSERT [LOW_PRIORITY | DELAYED | HIGH_PRIORITY] [IGNORE]
    [INTO] tbl_name
    [PARTITION (partition_name [, partition_name] ...)]
    SET assignment_list
    [AS row_alias[(col_alias [, col_alias] ...)]]
    [ON DUPLICATE KEY UPDATE assignment_list]

INSERT [LOW_PRIORITY | HIGH_PRIORITY] [IGNORE]
    [INTO] tbl_name
    [PARTITION (partition_name [, partition_name] ...)]
    [(col_name [, col_name] ...)]
    { SELECT ...
      | TABLE table_name
      | VALUES row_constructor_list
    }
    [ON DUPLICATE KEY UPDATE assignment_list]

value:
    {expr | DEFAULT}

value_list:
    value [, value] ...

row_constructor_list:
    ROW(value_list)[, ROW(value_list)][, ...]

assignment:
    col_name =
          value
        | [row_alias.]col_name
        | [tbl_name.]col_name
        | [row_alias.]col_alias

assignment_list:
    assignment [, assignment] ...
*/

// InsertBuilder struct represents a builder for constructing SQL INSERT statements.
// It contains components for managing the INSERT clause, rows, and query statements.
type InsertBuilder struct {
	// insertStatement represents the INSERT clause, including the table name and columns.
	insertStatement Insert
	// rowStatement represents the rows to be inserted into the specified table.
	rowStatement InsertRows
	// queryStatement represents a subquery for the INSERT statement.
	queryStatement InsertQuery
}

// InsertInstance creates and returns a new instance of InsertBuilder.
// It initializes an empty InsertBuilder structure.
//
// Returns:
//
//	*InsertBuilder - A new instance of the InsertBuilder structure.
func InsertInstance() *InsertBuilder {
	return &InsertBuilder{}
}

// ====================================================================
//                   Insert Builder :: Operators
// ====================================================================

// String constructs and returns the complete SQL INSERT statement as a string.
// It combines the insertStatement, rowStatement, and queryStatement components.
//
// Returns:
//
//	string - A string representation of the SQL INSERT statement.
func (ib *InsertBuilder) String() string {
	var queryParts []string
	var sqlStr string

	// Append the INSERT clause.
	queryParts = append(queryParts, ib.insertStatement.String())

	// Append the ROWS clause if present.
	sqlStr = ib.rowStatement.String()
	if sqlStr != "" {
		queryParts = append(queryParts, sqlStr)
	}

	// Append the QUERY clause if present.
	sqlStr = ib.queryStatement.String()
	if sqlStr != "" {
		queryParts = append(queryParts, sqlStr)
	}

	// Combine all parts into a single SQL string.
	sql := strings.Join(queryParts, " ")

	return sql
}

// Insert sets the table name and column names for the INSERT statement.
//
// Parameters:
//   - table string: The name of the table into which the data will be inserted.
//   - columns ...string: The column names for the INSERT statement.
//
// Returns:
//
//	*InsertBuilder - The updated InsertBuilder instance.
func (ib *InsertBuilder) Insert(table string, columns ...string) *InsertBuilder {
	ib.insertStatement.Table = table
	ib.insertStatement.Columns = columns

	return ib
}

// Row appends a new row of values to the INSERT statement.
//
// Parameters:
//   - values ...any: The values to be inserted into a new row.
//
// Returns:
//
//	*InsertBuilder - The updated InsertBuilder instance.
func (ib *InsertBuilder) Row(values ...any) *InsertBuilder {
	ib.rowStatement.Append(values...)

	return ib
}

// Query sets a subquery for the INSERT statement.
//
// Parameters:
//   - query *QueryBuilder: The subquery to be used in the INSERT statement.
//
// Returns:
//
//	*InsertBuilder - The updated InsertBuilder instance.
func (ib *InsertBuilder) Query(query *QueryBuilder) *InsertBuilder {
	ib.queryStatement.Query = query

	return ib
}
