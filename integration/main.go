package main

import (
	"fmt"
	qb "github.com/jiveio/fluentsql"
)

func main() {
	sql := qb.NewQueryBuilder().
		Select().
		From("employees").
		String()

	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "hire_date").
		From("employees").
		String()

	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("first_name", "last_name", "salary", "salary * 1.05").
		From("employees").
		String()

	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("first_name", "last_name", "salary", "salary * 1.05 AS new_salary").
		From("employees").
		String()

	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("first_name", "last_name", "hire_date", "salary").
		From("employees").
		OrderBy("first_name", qb.Asc).
		String()

	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("first_name", "last_name", "hire_date", "salary").
		From("employees").
		OrderBy("first_name", qb.Asc).
		OrderBy("last_name", qb.Desc).
		String()

	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "hire_date", "salary").
		From("employees").
		OrderBy("salary", qb.Desc).
		String()

	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "hire_date", "salary").
		From("employees").
		OrderBy("hire_date", qb.Asc).
		String()

	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("DISTINCT salary").
		From("employees").
		OrderBy("salary", qb.Desc).
		String()

	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("DISTINCT job_id", "salary").
		From("employees").
		OrderBy("job_id", qb.Asc).
		OrderBy("salary", qb.Desc).
		String()

	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("DISTINCT phone_number").
		From("employees").
		OrderBy("phone_number", qb.Asc).
		String()

	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name").
		From("employees").
		OrderBy("first_name", qb.Asc).
		Limit(5, 0).
		String()

	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name").
		From("employees").
		OrderBy("first_name", qb.Asc).
		Limit(5, 3).
		String()

	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "salary").
		From("employees").
		Where("salary", qb.Eq,
			qb.NewQueryBuilder().
				Select("DISTINCT salary").
				From("employees").
				OrderBy("salary", qb.Desc).
				Limit(1, 1),
		).
		String()

	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "salary").
		From("employees").
		OrderBy("salary", qb.Desc).
		Fetch(0, 1).
		String()

	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "salary").
		From("employees").
		OrderBy("salary", qb.Desc).
		Fetch(5, 5).
		String()

	fmt.Println("SQL> ", sql)
}
