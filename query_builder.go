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
func (qb *QueryBuilder) Where(Field string, Opt WhereOpt, Value any) *QueryBuilder {
	qb.Query.Where.Conditions = append(qb.Query.Where.Conditions, Condition{
		Field: Field,
		Opt:   Opt,
		Value: Value,
		AndOr: And,
	})

	return qb
}

// WhereNull builder
func (qb *QueryBuilder) WhereNull(Field string) *QueryBuilder {
	qb.Query.Where.Conditions = append(qb.Query.Where.Conditions, Condition{
		Field: Field,
		Opt:   Null,
		AndOr: And,
	})

	return qb
}

// WhereNotNull builder
func (qb *QueryBuilder) WhereNotNull(Field string) *QueryBuilder {
	qb.Query.Where.Conditions = append(qb.Query.Where.Conditions, Condition{
		Field: Field,
		Opt:   NotNull,
		AndOr: And,
	})

	return qb
}

// WhereOr builder
func (qb *QueryBuilder) WhereOr(Field string, Opt WhereOpt, Value any) *QueryBuilder {
	qb.Query.Where.Conditions = append(qb.Query.Where.Conditions, Condition{
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

// AS of query builder
// Need for case build SELECT query which the column from another SELECT
// Example: Count number product for each category.
// SQL> SELECT s.name, (SELECT COUNT(*) FROM product AS p WHERE p.store_id=s.id) FROM store AS s
func (qb *QueryBuilder) AS(alias string) *QueryBuilder {
	qb.Query.Alias = alias

	return qb
}
