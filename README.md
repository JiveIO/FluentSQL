# Fluent SQL

Test
```bash
go test -v
``` 

## SQL Builder

Reference 

    https://www.sqltutorial.org/
    https://learnsql.com/
    https://www.w3schools.com/sql/sql_syntax.asp

## Group `Builder`

    QueryBuilder: SELECT - extracts data from a database
        Don't support
            `GROUPING SETS` https://www.sqltutorial.org/sql-grouping-sets/
            `ROLLUP` https://www.sqltutorial.org/sql-rollup/
            `CUBE` https://www.sqltutorial.org/sql-cube/
            `UNION` Operator https://www.sqltutorial.org/sql-union/
            `INTERSECT` Operator https://www.sqltutorial.org/sql-intersect/
            `MINUS` Operator https://www.sqltutorial.org/sql-minus/
            `CASE` https://www.sqltutorial.org/sql-case/ (Working)

    UpdateBuilder: UPDATE - updates data in a database
    DeleteBuilder: DELETE - deletes data from a database
    InsertBuilder: INSERT INTO - inserts new data into a database

## Group `Seeder and Migration`

    CREATE DATABASE - creates a new database
    ALTER DATABASE - modifies a database
    CREATE TABLE - creates a new table
    ALTER TABLE - modifies a table
    DROP TABLE - deletes a table
    CREATE INDEX - creates an index (search key)
    DROP INDEX - deletes an index
