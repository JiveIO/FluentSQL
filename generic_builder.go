package fluentsql

// ====================================================================
// =========================== Interfaces =============================
// ====================================================================

// Builder is a generic interface for SQL query builders.
// It defines methods that all builder types should implement.
type Builder[T any] interface {
	// WithDialect sets the dialect for the builder.
	WithDialect(dialect Dialect) T

	// GetDialect returns the current dialect of the builder.
	GetDialect() Dialect
}

// ====================================================================
// ====================== Generic Query Builder =======================
// ====================================================================

// GenericQueryBuilder is a generic version of QueryBuilder that supports different dialects.
type GenericQueryBuilder struct {
	*QueryBuilder
	dialect Dialect
}

// NewQueryBuilder creates a new GenericQueryBuilder with the specified dialect.
// If no dialect is provided, it uses the default dialect.
func NewQueryBuilder(dialect ...Dialect) *GenericQueryBuilder {
	var d Dialect
	if len(dialect) > 0 {
		d = dialect[0]
	} else {
		d = DefaultDialect()
	}

	return &GenericQueryBuilder{
		QueryBuilder: QueryInstance(),
		dialect:      d,
	}
}

// WithDialect sets the dialect for the GenericQueryBuilder.
func (qb *GenericQueryBuilder) WithDialect(dialect Dialect) *GenericQueryBuilder {
	qb.dialect = dialect
	return qb
}

// GetDialect returns the current dialect of the GenericQueryBuilder.
func (qb *GenericQueryBuilder) GetDialect() Dialect {
	return qb.dialect
}

// Sql constructs the SQL query string and associated arguments using the builder's dialect.
func (qb *GenericQueryBuilder) Sql() (string, []any, error) {
	var args []any
	return qb.StringArgs(args)
}

// StringArgs constructs the SQL query string with placeholders and its associated arguments
// using the builder's dialect.
func (qb *GenericQueryBuilder) StringArgs(args []any) (string, []any, error) {
	// Save the original global dialect
	originalDialect := GetDialect()

	// Set the global dialect to the builder's dialect
	SetDialect(qb.dialect)

	// Use the original QueryBuilder's StringArgs method
	sql, args, err := qb.QueryBuilder.StringArgs(args)

	// Restore the original global dialect
	SetDialect(originalDialect)

	return sql, args, err
}

// ====================================================================
// ====================== Generic Insert Builder ======================
// ====================================================================

// GenericInsertBuilder is a generic version of InsertBuilder that supports different dialects.
type GenericInsertBuilder struct {
	*InsertBuilder
	dialect Dialect
}

// NewInsertBuilder creates a new GenericInsertBuilder with the specified dialect.
// If no dialect is provided, it uses the default dialect.
func NewInsertBuilder(dialect ...Dialect) *GenericInsertBuilder {
	var d Dialect
	if len(dialect) > 0 {
		d = dialect[0]
	} else {
		d = DefaultDialect()
	}

	return &GenericInsertBuilder{
		InsertBuilder: InsertInstance(),
		dialect:       d,
	}
}

// WithDialect sets the dialect for the GenericInsertBuilder.
func (ib *GenericInsertBuilder) WithDialect(dialect Dialect) *GenericInsertBuilder {
	ib.dialect = dialect
	return ib
}

// GetDialect returns the current dialect of the GenericInsertBuilder.
func (ib *GenericInsertBuilder) GetDialect() Dialect {
	return ib.dialect
}

// Sql constructs the SQL query string and associated arguments using the builder's dialect.
func (ib *GenericInsertBuilder) Sql() (string, []any, error) {
	var args []any
	return ib.StringArgs(args)
}

// StringArgs constructs the SQL query string with placeholders and its associated arguments
// using the builder's dialect.
func (ib *GenericInsertBuilder) StringArgs(args []any) (string, []any, error) {
	// Save the original global dialect
	originalDialect := GetDialect()

	// Set the global dialect to the builder's dialect
	SetDialect(ib.dialect)

	// Use the original InsertBuilder's StringArgs method
	sql, args, err := ib.InsertBuilder.StringArgs(args)

	// Restore the original global dialect
	SetDialect(originalDialect)

	return sql, args, err
}

// ====================================================================
// ====================== Generic Update Builder ======================
// ====================================================================

// GenericUpdateBuilder is a generic version of UpdateBuilder that supports different dialects.
type GenericUpdateBuilder struct {
	*UpdateBuilder
	dialect Dialect
}

// NewUpdateBuilder creates a new GenericUpdateBuilder with the specified dialect.
// If no dialect is provided, it uses the default dialect.
func NewUpdateBuilder(dialect ...Dialect) *GenericUpdateBuilder {
	var d Dialect
	if len(dialect) > 0 {
		d = dialect[0]
	} else {
		d = DefaultDialect()
	}

	return &GenericUpdateBuilder{
		UpdateBuilder: UpdateInstance(),
		dialect:       d,
	}
}

// WithDialect sets the dialect for the GenericUpdateBuilder.
func (ub *GenericUpdateBuilder) WithDialect(dialect Dialect) *GenericUpdateBuilder {
	ub.dialect = dialect
	return ub
}

// GetDialect returns the current dialect of the GenericUpdateBuilder.
func (ub *GenericUpdateBuilder) GetDialect() Dialect {
	return ub.dialect
}

// Sql constructs the SQL query string and associated arguments using the builder's dialect.
func (ub *GenericUpdateBuilder) Sql() (string, []any, interface{}) {
	// Save the original global dialect
	originalDialect := GetDialect()

	// Set the global dialect to the builder's dialect
	SetDialect(ub.dialect)

	// Use the original UpdateBuilder's Sql method
	sql, args, err := ub.UpdateBuilder.Sql()

	// Restore the original global dialect
	SetDialect(originalDialect)

	return sql, args, err
}

// ====================================================================
// ====================== Generic Delete Builder ======================
// ====================================================================

// GenericDeleteBuilder is a generic version of DeleteBuilder that supports different dialects.
type GenericDeleteBuilder struct {
	*DeleteBuilder
	dialect Dialect
}

// NewDeleteBuilder creates a new GenericDeleteBuilder with the specified dialect.
// If no dialect is provided, it uses the default dialect.
func NewDeleteBuilder(dialect ...Dialect) *GenericDeleteBuilder {
	var d Dialect
	if len(dialect) > 0 {
		d = dialect[0]
	} else {
		d = DefaultDialect()
	}

	return &GenericDeleteBuilder{
		DeleteBuilder: DeleteInstance(),
		dialect:       d,
	}
}

// WithDialect sets the dialect for the GenericDeleteBuilder.
func (db *GenericDeleteBuilder) WithDialect(dialect Dialect) *GenericDeleteBuilder {
	db.dialect = dialect
	return db
}

// GetDialect returns the current dialect of the GenericDeleteBuilder.
func (db *GenericDeleteBuilder) GetDialect() Dialect {
	return db.dialect
}

// Sql constructs the SQL query string and associated arguments using the builder's dialect.
func (db *GenericDeleteBuilder) Sql() (string, []any, error) {
	var args []any
	return db.StringArgs(args)
}

// StringArgs constructs the SQL query string with placeholders and its associated arguments
// using the builder's dialect.
func (db *GenericDeleteBuilder) StringArgs(args []any) (string, []any, error) {
	// Save the original global dialect
	originalDialect := GetDialect()

	// Set the global dialect to the builder's dialect
	SetDialect(db.dialect)

	// Use the original DeleteBuilder's StringArgs method
	sql, args, err := db.DeleteBuilder.StringArgs(args)

	// Restore the original global dialect
	SetDialect(originalDialect)

	return sql, args, err
}
