package fluentsql

// InsertQuery represents a query that can be used as a subquery in an INSERT statement.
type InsertQuery struct {
	// Query stores any query information, typically expected to be a pointer to QueryBuilder.
	Query any
}

// String converts the query to a string representation.
//
// Returns:
//   - string: The string representation of the query. If the Query is a QueryBuilder,
//     it calls the QueryBuilder's String method; otherwise, it returns an empty string.
func (q *InsertQuery) String() string {
	if queryBuilder, ok := q.Query.(*QueryBuilder); ok {
		// Call the String method of QueryBuilder if Query is of type *QueryBuilder.
		return queryBuilder.String()
	}

	// Return an empty string if Query is not of type *QueryBuilder.
	return ""
}
