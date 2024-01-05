package fluentsql

import (
	"testing"
)

// TestQueryBasic
func TestQueryBasic(t *testing.T) {
	testCases := map[string]*QueryBuilder{
		"SELECT first_name, last_name, salary, salary * 1.05 AS new_salary FROM employees": NewQueryBuilder().
			Select("first_name", "last_name", "salary", "salary * 1.05 AS new_salary").
			From("employees"),
		"SELECT first_name, last_name, hire_date, salary FROM employees ORDER BY first_name ASC, last_name DESC": NewQueryBuilder().
			Select("first_name", "last_name", "hire_date", "salary").
			From("employees").
			OrderBy("first_name", Asc).
			OrderBy("last_name", Desc),
		"SELECT DISTINCT salary FROM employees ORDER BY salary DESC": NewQueryBuilder().
			Select("DISTINCT salary").
			From("employees").
			OrderBy("salary", Desc),
		"SELECT employee_id, first_name, last_name FROM employees ORDER BY first_name ASC LIMIT 5 OFFSET 3": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name").
			From("employees").
			OrderBy("first_name", Asc).
			Limit(5, 3),
		"SELECT employee_id, first_name, last_name, hire_date FROM employees WHERE DATE_PART('year', hire_date) = 1999 ORDER BY hire_date DESC": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name", "hire_date").
			From("employees").
			Where(FieldYear("hire_date"), Eq, 1999).
			OrderBy("hire_date", Desc),
		"SELECT employee_id, first_name, last_name, salary FROM employees WHERE salary = 7000 OR salary = 8000 ORDER BY salary ASC": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name", "salary").
			From("employees").
			Where("salary", Eq, 7000).
			WhereOr("salary", Eq, 8000).
			OrderBy("salary", Asc),
	}

	for expected, query := range testCases {
		if query.String() != expected {
			t.Fatalf(`Query %s != %s`, query.String(), expected)
		}
	}
}

// TestQueryLikeAndInAndNotNull
func TestQueryLikeAndInAndNotNull(t *testing.T) {
	testCases := map[string]*QueryBuilder{
		"SELECT first_name, last_name, salary, salary * 1.05 AS new_salary FROM employees": NewQueryBuilder().
			Select("first_name", "last_name", "salary", "salary * 1.05 AS new_salary").
			From("employees"),
		"SELECT employee_id, first_name, last_name FROM employees WHERE first_name LIKE 'Jo%' ORDER BY first_name ASC": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name").
			From("employees").
			Where("first_name", Like, "Jo%").
			OrderBy("first_name", Asc),
		"SELECT employee_id, first_name, last_name FROM employees WHERE first_name LIKE 'S%' AND first_name NOT LIKE 'Sh%' ORDER BY first_name ASC": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name").
			From("employees").
			Where("first_name", Like, "S%").
			Where("first_name", NotLike, "Sh%").
			OrderBy("first_name", Asc),
		"SELECT employee_id, first_name, last_name, department_id FROM employees WHERE department_id IN (8, 9) ORDER BY department_id ASC": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name", "department_id").
			From("employees").
			Where("department_id", In, []int{8, 9}).
			OrderBy("department_id", Asc),
		"SELECT employee_id, first_name, last_name, phone_number FROM employees WHERE phone_number IS NOT NULL": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name", "phone_number").
			From("employees").
			Where("phone_number", NotNull, nil),
	}

	for expected, query := range testCases {
		if query.String() != expected {
			t.Fatalf(`Query %s != %s`, query.String(), expected)
		}
	}
}

// TestQueryFieldYear
func TestQueryFieldYear(t *testing.T) {
	testCases := map[string]*QueryBuilder{
		"SELECT employee_id, first_name, last_name, hire_date FROM employees WHERE DATE_PART('year', hire_date) = 1999 ORDER BY hire_date DESC": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name", "hire_date").
			From("employees").
			Where(FieldYear("hire_date"), Eq, 1999).
			OrderBy("hire_date", Desc),
		"SELECT employee_id, first_name, last_name, DATE_PART('year', hire_date) joined_year FROM employees WHERE DATE_PART('year', hire_date) BETWEEN 1990 AND 1993 ORDER BY hire_date ASC": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name", FieldYear("hire_date").String()+" joined_year").
			From("employees").
			Where(FieldYear("hire_date"),
				Between, ValueBetween{
					Low:  1990,
					High: 1993,
				}).
			OrderBy("hire_date", Asc),
		"SELECT employee_id, first_name, last_name, DATE_PART('year', hire_date) FROM employees WHERE DATE_PART('year', hire_date) BETWEEN 1990 AND 1993 ORDER BY hire_date ASC": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name", FieldYear("hire_date")).
			From("employees").
			Where(FieldYear("hire_date"),
				Between, ValueBetween{
					Low:  1990,
					High: 1993,
				}).
			OrderBy("hire_date", Asc),
	}

	for expected, query := range testCases {
		if query.String() != expected {
			t.Fatalf(`Query %s != %s`, query.String(), expected)
		}
	}
}

// TestQueryLikeAndInAndNotNull
func TestQueryBetween(t *testing.T) {
	testCases := map[string]*QueryBuilder{
		"SELECT employee_id, first_name, last_name, salary FROM employees WHERE salary NOT BETWEEN 3000 AND 5000": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name", "salary").
			Where("salary", NotBetween, ValueBetween{Low: 3000, High: 5000}).
			From("employees"),
		"SELECT employee_id, first_name, last_name, DATE_PART('year', hire_date) joined_year FROM employees WHERE DATE_PART('year', hire_date) BETWEEN 1990 AND 1993 ORDER BY hire_date ASC": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name", FieldYear("hire_date").String()+" joined_year").
			From("employees").
			Where(FieldYear("hire_date"),
				Between, ValueBetween{
					Low:  1990,
					High: 1993,
				}).
			OrderBy("hire_date", Asc),
	}

	for expected, query := range testCases {
		if query.String() != expected {
			t.Fatalf(`Query %s != %s`, query.String(), expected)
		}
	}
}

// TestQuerySubquery
func TestQuerySubquery(t *testing.T) {
	testCases := map[string]*QueryBuilder{
		"SELECT employee_id, first_name, last_name, salary FROM employees WHERE salary = (SELECT DISTINCT salary FROM employees ORDER BY salary DESC LIMIT 1 OFFSET 1)": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name", "salary").
			From("employees").
			Where("salary", Eq,
				NewQueryBuilder().
					Select("DISTINCT salary").
					From("employees").
					OrderBy("salary", Desc).
					Limit(1, 1),
			),
		"SELECT employee_id, first_name, last_name, salary FROM employees WHERE salary >= ALL (SELECT salary FROM employees WHERE department_id = 8) ORDER BY salary DESC": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name", "salary").
			From("employees").
			Where("salary", GrEqAll,
				NewQueryBuilder().
					Select("salary").
					From("employees").
					Where("department_id", Eq, 8),
			).
			OrderBy("salary", Desc),
		"SELECT employee_id, first_name, last_name, salary FROM employees e WHERE  EXISTS (SELECT 1 FROM dependents d WHERE d.employee_id = e.employee_id) ORDER BY first_name ASC, last_name ASC": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name", "salary").
			From("employees", "e").
			Where(FieldEmpty(""), Exists,
				NewQueryBuilder().
					Select("1").
					From("dependents", "d").
					Where("d.employee_id", Eq, ValueField("e.employee_id")),
			).
			OrderBy("first_name", Asc).
			OrderBy("last_name", Asc),
		"SELECT employee_id, first_name, last_name, salary FROM employees WHERE department_id IN (SELECT department_id FROM departments WHERE department_name = 'Marketing' OR department_name = 'Sales')": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name", "salary").
			From("employees").
			Where("department_id", In,
				NewQueryBuilder().
					Select("department_id").
					From("departments").
					Where("department_name", Eq, "Marketing").
					WhereOr("department_name", Eq, "Sales"),
			),
		"SELECT employee_id, first_name, last_name FROM employees e WHERE  NOT EXISTS (SELECT employee_id FROM dependents d WHERE d.employee_id = e.employee_id)": NewQueryBuilder().
			Select("employee_id", "first_name", "last_name").
			From("employees", "e").
			Where(FieldEmpty(""), NotExists,
				NewQueryBuilder().
					Select("employee_id").
					From("dependents", "d").
					Where("d.employee_id", Eq, ValueField("e.employee_id")),
			),
	}

	for expected, query := range testCases {
		if query.String() != expected {
			t.Fatalf(`Query %s != %s`, query.String(), expected)
		}
	}
}

// TestQueryAlias
func TestQueryAlias(t *testing.T) {
	testCases := map[string]*QueryBuilder{
		"SELECT inv_no AS invoice_no, amount, due_date AS 'Due date', cust_no 'Customer No' FROM invoices": NewQueryBuilder().
			Select("inv_no AS invoice_no", "amount", "due_date AS 'Due date'", "cust_no 'Customer No'").
			From("invoices"),
		"SELECT first_name, last_name, salary * 1.1 AS new_salary FROM employees WHERE new_salary > 5000": NewQueryBuilder().
			Select("first_name", "last_name", "salary * 1.1 AS new_salary").
			From("employees").
			Where("new_salary", Greater, 5000),
	}

	for expected, query := range testCases {
		if query.String() != expected {
			t.Fatalf(`Query %s != %s`, query.String(), expected)
		}
	}
}

// TestQueryJoin
func TestQueryJoin(t *testing.T) {
	testCases := map[string]*QueryBuilder{
		"SELECT first_name, last_name, employees.department_id, departments.department_id, department_name FROM employees INNER JOIN departments ON departments.department_id = employees.department_id WHERE employees.department_id IN (1, 2, 3)": NewQueryBuilder().
			Select("first_name", "last_name", "employees.department_id", "departments.department_id", "department_name").
			From("employees").
			Join(InnerJoin, "departments", Condition{
				Field: "departments.department_id",
				Opt:   Eq,
				Value: ValueField("employees.department_id"),
			}).
			Where("employees.department_id", In, []int{1, 2, 3}),
		"SELECT first_name, last_name, job_title, department_name FROM employees e INNER JOIN departments d ON d.department_id = e.department_id INNER JOIN jobs j ON j.job_id = e.job_id WHERE e.department_id IN (1, 2, 3)": NewQueryBuilder().
			Select("first_name", "last_name", "job_title", "department_name").
			From("employees", "e").
			Join(InnerJoin, "departments d", Condition{
				Field: "d.department_id",
				Opt:   Eq,
				Value: ValueField("e.department_id"),
			}).
			Join(InnerJoin, "jobs j", Condition{
				Field: "j.job_id",
				Opt:   Eq,
				Value: ValueField("e.job_id"),
			}).
			Where("e.department_id", In, []int{1, 2, 3}),
		"SELECT c.country_name, c.country_id, l.country_id, l.street_address, l.city FROM countries c LEFT JOIN locations l ON l.country_id = c.country_id WHERE c.country_id IN ('US', 'UK', 'CN')": NewQueryBuilder().
			Select("c.country_name", "c.country_id", "l.country_id", "l.street_address", "l.city").
			From("countries", "c").
			Join(LeftJoin, "locations l", Condition{
				Field: "l.country_id",
				Opt:   Eq,
				Value: ValueField("c.country_id"),
			}).
			Where("c.country_id", In, []string{"US", "UK", "CN"}),
		"SELECT country_name FROM countries c LEFT JOIN locations l ON l.country_id = c.country_id WHERE l.location_id IS NULL ORDER BY country_name ASC": NewQueryBuilder().
			Select("country_name").
			From("countries", "c").
			Join(LeftJoin, "locations l", Condition{
				Field: "l.country_id",
				Opt:   Eq,
				Value: ValueField("c.country_id"),
			}).
			Where("l.location_id", Null, nil).
			OrderBy("country_name", Asc),
		"SELECT r.region_name, c.country_name, l.street_address, l.city FROM regions r LEFT JOIN countries c ON c.region_id = r.region_id LEFT JOIN locations l ON l.country_id = c.country_id WHERE c.country_id IN ('US', 'UK', 'CN')": NewQueryBuilder().
			Select("r.region_name", "c.country_name", "l.street_address", "l.city").
			From("regions", "r").
			Join(LeftJoin, "countries c", Condition{
				Field: "c.region_id",
				Opt:   Eq,
				Value: ValueField("r.region_id"),
			}).
			Join(LeftJoin, "locations l", Condition{
				Field: "l.country_id",
				Opt:   Eq,
				Value: ValueField("c.country_id"),
			}).
			Where("c.country_id", In, []string{"US", "UK", "CN"}),
		"SELECT e.first_name || ' ' || e.last_name AS employee, m.first_name || ' ' || m.last_name AS manager FROM employees e INNER JOIN employees m ON m.employee_id = e.manager_id ORDER BY manager ASC": NewQueryBuilder().
			Select("e.first_name || ' ' || e.last_name AS employee", "m.first_name || ' ' || m.last_name AS manager").
			From("employees", "e").
			Join(InnerJoin, "employees m", Condition{
				Field: "m.employee_id",
				Opt:   Eq,
				Value: ValueField("e.manager_id"),
			}).
			OrderBy("manager", Asc),
		"SELECT basket_name, fruit_name FROM fruits FULL OUTER JOIN baskets ON baskets.basket_id = fruits.basket_id WHERE fruit_name IS NULL": NewQueryBuilder().
			Select("basket_name", "fruit_name").
			From("fruits").
			Join(FullOuterJoin, "baskets", Condition{
				Field: "baskets.basket_id",
				Opt:   Eq,
				Value: ValueField("fruits.basket_id"),
			}).
			Where("fruit_name", Null, nil),
		"SELECT sales_org, channel FROM sales_organization CROSS JOIN sales_channel": NewQueryBuilder().
			Select("sales_org", "channel").
			From("sales_organization").
			Join(CrossJoin, "sales_channel", Condition{}),
	}

	for expected, query := range testCases {
		if query.String() != expected {
			t.Fatalf(`Query %s != %s`, query.String(), expected)
		}
	}
}
