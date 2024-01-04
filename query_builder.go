package fluentsql

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
