package fluentsql

import (
	"strings"
	"testing"
)

func TestGenericQueryBuilder(t *testing.T) {
	// Test with MySQL dialect
	mysqlBuilder := NewQueryBuilder(MySQLDialect{})
	mysqlBuilder.Select("id", "name").
		From("users").
		Where("active", Eq, true)

	mysqlSQL, mysqlArgs, err := mysqlBuilder.Sql()
	if err != nil {
		t.Errorf("Error generating MySQL SQL: %v", err)
	}

	expectedMySQLSQL := "SELECT id, name FROM users WHERE active = ?"
	if mysqlSQL != expectedMySQLSQL {
		t.Errorf("MySQL SQL does not match expected.\nExpected: %s\nActual: %s", expectedMySQLSQL, mysqlSQL)
	}

	if len(mysqlArgs) != 1 || mysqlArgs[0] != true {
		t.Errorf("MySQL args do not match expected. Expected [true], got %v", mysqlArgs)
	}

	// Test with PostgreSQL dialect
	pgBuilder := NewQueryBuilder(PostgreSQLDialect{})
	pgBuilder.Select("id", "name").
		From("users").
		Where("active", Eq, true)

	pgSQL, pgArgs, err := pgBuilder.Sql()
	if err != nil {
		t.Errorf("Error generating PostgreSQL SQL: %v", err)
	}

	expectedPgSQL := "SELECT id, name FROM users WHERE active = $1"
	if pgSQL != expectedPgSQL {
		t.Errorf("PostgreSQL SQL does not match expected.\nExpected: %s\nActual: %s", expectedPgSQL, pgSQL)
	}

	if len(pgArgs) != 1 || pgArgs[0] != true {
		t.Errorf("PostgreSQL args do not match expected. Expected [true], got %v", pgArgs)
	}

	// Test with SQLite dialect
	sqliteBuilder := NewQueryBuilder(SQLiteDialect{})
	sqliteBuilder.Select("id", "name").
		From("users").
		Where("active", Eq, true)

	sqliteSQL, sqliteArgs, err := sqliteBuilder.Sql()
	if err != nil {
		t.Errorf("Error generating SQLite SQL: %v", err)
	}

	expectedSQLiteSQL := "SELECT id, name FROM users WHERE active = ?"
	if sqliteSQL != expectedSQLiteSQL {
		t.Errorf("SQLite SQL does not match expected.\nExpected: %s\nActual: %s", expectedSQLiteSQL, sqliteSQL)
	}

	if len(sqliteArgs) != 1 || sqliteArgs[0] != true {
		t.Errorf("SQLite args do not match expected. Expected [true], got %v", sqliteArgs)
	}
}

func TestGenericInsertBuilder(t *testing.T) {
	// Test with MySQL dialect
	mysqlBuilder := NewInsertBuilder(MySQLDialect{})
	mysqlBuilder.Insert("users", "id", "name").
		Row(1, "John")

	mysqlSQL, mysqlArgs, err := mysqlBuilder.Sql()
	if err != nil {
		t.Errorf("Error generating MySQL SQL: %v", err)
	}

	// The actual SQL might vary, but we're testing that the placeholders are correct
	if !strings.Contains(mysqlSQL, "INSERT INTO") || !strings.Contains(mysqlSQL, "users") || !strings.Contains(mysqlSQL, "id") || !strings.Contains(mysqlSQL, "name") || !strings.Contains(mysqlSQL, "VALUES") || !strings.Contains(mysqlSQL, "?") {
		t.Errorf("MySQL SQL does not match expected format. Got: %s", mysqlSQL)
	}

	if len(mysqlArgs) != 2 || mysqlArgs[0] != 1 || mysqlArgs[1] != "John" {
		t.Errorf("MySQL args do not match expected. Expected [1, \"John\"], got %v", mysqlArgs)
	}

	// Test with PostgreSQL dialect
	pgBuilder := NewInsertBuilder(PostgreSQLDialect{})
	pgBuilder.Insert("users", "id", "name").
		Row(1, "John")

	pgSQL, pgArgs, err := pgBuilder.Sql()
	if err != nil {
		t.Errorf("Error generating PostgreSQL SQL: %v", err)
	}

	// The actual SQL might vary, but we're testing that the placeholders are correct
	if !strings.Contains(pgSQL, "INSERT INTO") || !strings.Contains(pgSQL, "users") || !strings.Contains(pgSQL, "id") || !strings.Contains(pgSQL, "name") || !strings.Contains(pgSQL, "VALUES") || !strings.Contains(pgSQL, "$1") || !strings.Contains(pgSQL, "$2") {
		t.Errorf("PostgreSQL SQL does not match expected format. Got: %s", pgSQL)
	}

	if len(pgArgs) != 2 || pgArgs[0] != 1 || pgArgs[1] != "John" {
		t.Errorf("PostgreSQL args do not match expected. Expected [1, \"John\"], got %v", pgArgs)
	}
}

func TestGenericUpdateBuilder(t *testing.T) {
	// Test with MySQL dialect
	mysqlBuilder := NewUpdateBuilder(MySQLDialect{})
	mysqlBuilder.Update("users").
		Set("name", "John").
		Where("id", Eq, 1)

	mysqlSQL, mysqlArgs, _ := mysqlBuilder.Sql()

	// The actual SQL might vary, but we're testing that the placeholders are correct
	if !strings.Contains(mysqlSQL, "UPDATE") || !strings.Contains(mysqlSQL, "users") || !strings.Contains(mysqlSQL, "SET") || !strings.Contains(mysqlSQL, "name") || !strings.Contains(mysqlSQL, "?") || !strings.Contains(mysqlSQL, "WHERE") || !strings.Contains(mysqlSQL, "id") {
		t.Errorf("MySQL SQL does not match expected format. Got: %s", mysqlSQL)
	}

	if len(mysqlArgs) != 2 || mysqlArgs[0] != "John" || mysqlArgs[1] != 1 {
		t.Errorf("MySQL args do not match expected. Expected [\"John\", 1], got %v", mysqlArgs)
	}

	// Test with PostgreSQL dialect
	pgBuilder := NewUpdateBuilder(PostgreSQLDialect{})
	pgBuilder.Update("users").
		Set("name", "John").
		Where("id", Eq, 1)

	pgSQL, pgArgs, _ := pgBuilder.Sql()

	expectedPgSQL := "UPDATE users SET name = $1 WHERE id = $2"
	if pgSQL != expectedPgSQL {
		t.Errorf("PostgreSQL SQL does not match expected.\nExpected: %s\nActual: %s", expectedPgSQL, pgSQL)
	}

	if len(pgArgs) != 2 || pgArgs[0] != "John" || pgArgs[1] != 1 {
		t.Errorf("PostgreSQL args do not match expected. Expected [\"John\", 1], got %v", pgArgs)
	}
}

func TestGenericDeleteBuilder(t *testing.T) {
	// Test with MySQL dialect
	mysqlBuilder := NewDeleteBuilder(MySQLDialect{})
	mysqlBuilder.Delete("users").
		Where("id", Eq, 1)

	mysqlSQL, mysqlArgs, err := mysqlBuilder.Sql()
	if err != nil {
		t.Errorf("Error generating MySQL SQL: %v", err)
	}

	expectedMySQLSQL := "DELETE FROM users WHERE id = ?"
	if mysqlSQL != expectedMySQLSQL {
		t.Errorf("MySQL SQL does not match expected.\nExpected: %s\nActual: %s", expectedMySQLSQL, mysqlSQL)
	}

	if len(mysqlArgs) != 1 || mysqlArgs[0] != 1 {
		t.Errorf("MySQL args do not match expected. Expected [1], got %v", mysqlArgs)
	}

	// Test with PostgreSQL dialect
	pgBuilder := NewDeleteBuilder(PostgreSQLDialect{})
	pgBuilder.Delete("users").
		Where("id", Eq, 1)

	pgSQL, pgArgs, err := pgBuilder.Sql()
	if err != nil {
		t.Errorf("Error generating PostgreSQL SQL: %v", err)
	}

	expectedPgSQL := "DELETE FROM users WHERE id = $1"
	if pgSQL != expectedPgSQL {
		t.Errorf("PostgreSQL SQL does not match expected.\nExpected: %s\nActual: %s", expectedPgSQL, pgSQL)
	}

	if len(pgArgs) != 1 || pgArgs[0] != 1 {
		t.Errorf("PostgreSQL args do not match expected. Expected [1], got %v", pgArgs)
	}
}

func TestMultipleDialects(t *testing.T) {
	// Create builders with different dialects
	mysqlBuilder := NewQueryBuilder(MySQLDialect{})
	pgBuilder := NewQueryBuilder(PostgreSQLDialect{})
	sqliteBuilder := NewQueryBuilder(SQLiteDialect{})

	// Build the same query with all dialects
	mysqlBuilder.Select("id", "name").From("users").Where("active", Eq, true)
	pgBuilder.Select("id", "name").From("users").Where("active", Eq, true)
	sqliteBuilder.Select("id", "name").From("users").Where("active", Eq, true)

	// Generate SQL for all dialects
	mysqlSQL, _, _ := mysqlBuilder.Sql()
	pgSQL, _, _ := pgBuilder.Sql()
	sqliteSQL, _, _ := sqliteBuilder.Sql()

	// Verify that the SQL is different for each dialect
	if mysqlSQL == pgSQL {
		t.Errorf("MySQL and PostgreSQL SQL should be different, but both are: %s", mysqlSQL)
	}

	if pgSQL == sqliteSQL {
		t.Errorf("PostgreSQL and SQLite SQL should be different, but both are: %s", pgSQL)
	}

	// Verify that the global dialect hasn't been affected
	originalDialect := GetDialect()
	if originalDialect.Name() != "PostgreSQL" { // Default is PostgreSQL
		t.Errorf("Global dialect has been changed. Expected PostgreSQL, got %v", originalDialect.Name())
	}
}
