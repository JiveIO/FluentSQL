package fluentsql

// ===========================================================================================================
//										Query Builder :: Structure
// ===========================================================================================================

type QueryBuilder struct {
	Query Query
}

type SelectBuilder struct {
	QueryBuilder
}

type FromBuilder struct {
	QueryBuilder
}

type WhereBuilder struct {
	QueryBuilder
}

type LimitBuilder struct {
	QueryBuilder
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
func (qb *QueryBuilder) Select(columns ...any) *SelectBuilder {
	_qb := &SelectBuilder{*qb}

	_qb.Query.Select.Columns = columns

	return _qb
}

// From builder
func (qb *QueryBuilder) From(table string, alias ...string) *FromBuilder {
	_qb := &FromBuilder{*qb}

	_qb.Query.From.Table = table

	// Table alias
	if len(alias) > 0 {
		_qb.Query.From.Alias = alias[0]
	}

	return _qb
}

// Where builder
func (qb *QueryBuilder) Where(Field string, Opt WhereOpt, Value any) *WhereBuilder {
	_qb := &WhereBuilder{*qb}

	_qb.Query.Where.Conditions = append(_qb.Query.Where.Conditions, Condition{
		Field: Field,
		Opt:   Opt,
		Value: Value,
		AndOr: And,
	})

	return _qb
}

// WhereOr builder
func (qb *QueryBuilder) WhereOr(Field string, Opt WhereOpt, Value any) *WhereBuilder {
	_qb := &WhereBuilder{*qb}

	_qb.Query.Where.Conditions = append(_qb.Query.Where.Conditions, Condition{
		Field: Field,
		Opt:   Opt,
		Value: Value,
		AndOr: Or,
	})

	return _qb
}

type FnWhereGroupBuilder func(query WhereBuilder) *WhereBuilder

// WhereGroup combine multi where conditions into a group.
// Example: Group 2 condition created_at and update_at.
// SQL> SELECT * FROM users WHERE first_name LIKE '%john%' AND (created_at > '2024-01-12' OR update_at >= '2024-01-12') LIMIT 10 OFFSET 0
func (qb *QueryBuilder) WhereGroup(groupCondition FnWhereGroupBuilder) *WhereBuilder {
	_qb := &WhereBuilder{*qb}

	// Create new WhereBuilder
	whereBuilder := groupCondition(WhereBuilder{*NewQueryBuilder()})

	cond := Condition{
		Group: whereBuilder.Query.Where.Conditions,
	}

	_qb.Query.Where.Conditions = append(_qb.Query.Where.Conditions, cond)

	return _qb
}

// Limit builder
func (qb *QueryBuilder) Limit(Limit, Offset int) *LimitBuilder {
	_qb := &LimitBuilder{*qb}

	_qb.Query.Limit.Limit = Limit
	_qb.Query.Limit.Offset = Offset

	return _qb
}

// AS of query builder
// Need for case build SELECT query which the column from another SELECT
// Example: Count number product for each category.
// SQL> SELECT s.name, (SELECT COUNT(*) FROM product AS p WHERE p.store_id=s.id) FROM store AS s
func (qb *QueryBuilder) AS(alias string) *QueryBuilder {
	qb.Query.Alias = alias

	return qb
}
