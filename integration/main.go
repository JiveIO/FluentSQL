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

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "salary").
		From("employees").
		Where("salary", qb.Greater, 14000).
		OrderBy("salary", qb.Desc).
		String()
	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "department_id").
		From("employees").
		Where("department_id", qb.Eq, 5).
		OrderBy("first_name", qb.Asc).
		String()
	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name").
		From("employees").
		Where("last_name", qb.Eq, "Chen").
		OrderBy("salary", qb.Desc).
		String()
	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "hire_date").
		From("employees").
		Where("hire_date", qb.GrEq, "1999-01-01").
		OrderBy("hire_date", qb.Desc).
		String()
	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "hire_date").
		From("employees").
		Where("YEAR (hire_date)", qb.Eq, 1999).
		OrderBy("hire_date", qb.Desc).
		String()
	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "phone_number").
		From("employees").
		WhereNull("phone_number").
		String()
	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "department_id").
		From("employees").
		Where("department_id", qb.NotEq, 8).
		OrderBy("first_name", qb.Asc).
		OrderBy("last_name", qb.Asc).
		String()
	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "salary").
		From("employees").
		Where("salary", qb.Greater, 10000).
		Where("department_id", qb.Eq, 8).
		OrderBy("salary", qb.Desc).
		String()
	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "salary").
		From("employees").
		Where("salary", qb.Eq, 7000).
		WhereOr("salary", qb.Eq, 8000).
		OrderBy("salary", qb.Asc).
		String()
	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "salary").
		From("employees").
		Where("salary", qb.Between, qb.BetweenValue{
			From: 9000,
			To:   12000,
		}).
		OrderBy("salary", qb.Asc).
		String()
	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "department_id").
		From("employees").
		Where("department_id", qb.In, []int{8, 9}).
		OrderBy("department_id", qb.Asc).
		String()
	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name").
		From("employees").
		Where("first_name", qb.Like, "Jo%").
		OrderBy("first_name", qb.Asc).
		String()
	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "salary").
		From("employees").
		Where("salary", qb.GrEqAll,
			qb.NewQueryBuilder().
				Select("salary").
				From("employees").
				Where("department_id", qb.Eq, 8),
		).
		OrderBy("salary", qb.Desc).
		String()
	fmt.Println("SQL> ", sql)

	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "salary").
		From("employees").
		Where("salary", qb.GreaterAny,
			qb.NewQueryBuilder().
				Select("AVG(salary)").
				From("employees").
				GroupBy("department_id"),
		).
		OrderBy("first_name", qb.Asc).
		OrderBy("last_name", qb.Asc).
		String()
	fmt.Println("SQL> ", sql)

	// EXISTS
	sql = qb.NewQueryBuilder().
		Select("employee_id", "first_name", "last_name", "salary").
		From("employees", " e").
		Where("", qb.Exists,
			qb.NewQueryBuilder().
				Select("1").
				From("dependents", "d").
				Where("d.employee_id", qb.Eq, qb.FieldValue("e.employee_id")),
		).
		OrderBy("first_name", qb.Asc).
		OrderBy("last_name", qb.Asc).
		String()
	fmt.Println("SQL> ", sql)
}
