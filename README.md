# Fluent SQL

Fluent SQL - flexible and powerful SQL string builder for Go

## Setup

Install module 

```bash
# Latest version
go get github.com/jivegroup/fluentsql@latest

# Specific version
go get github.com/jivegroup/fluentsql@v1.3.5
```

Use module
```go
import (
    qb "github.com/jivegroup/fluentsql"	
)
```

## QueryBuilder
QueryBuilder: SELECT - extracts data from a database

```go
import (
    qb "github.com/jivegroup/fluentsql"	
)

// ------------- Simple query -------------
sql := qb.QueryInstance().
    Select("employee_id", "first_name", "last_name", "department_id").
    From("employees").
    Where("department_id", qb.NotEq, 8).
    OrderBy("first_name", qb.Asc).
    OrderBy("last_name", qb.Asc).
    Limit(3, 0).
    String()

// ------------- Sub-query -------------
sql = qb.QueryInstance().
    Select("employee_id", "first_name", "last_name", "salary").
    From("employees").
    Where("salary", qb.Eq,
        qb.QueryInstance().
            Select("DISTINCT salary").
            From("employees").
            OrderBy("salary", qb.Desc).
            Limit(1, 1),
    ).
    String()

// ------------- Where Group -------------
sql = qb.QueryInstance().
    Select("employee_id", "first_name", "last_name", "salary").
    From("employees").
    Where("salary", qb.In, qb.ValueBetween{
        Low:  9000,
        High: 12000,
    }).
    WhereGroup(func(whereBuilder WhereBuilder) *WhereBuilder {
        whereBuilder.Where("age", qb.Eq, 25).
        WhereOr("work_year", qb.Eq, 10)

        return &whereBuilder
    }),
    String()

// ------------- JOIN query -------------
sql = qb.QueryInstance().
    Select("r.region_name", "c.country_name", "l.street_address", "l.city").
    From("regions", "r").
    Join(qb.LeftJoin, "countries c", qb.Condition{
        Field: "c.region_id",
        Opt:   qb.Eq,
        Value: qb.ValueField("r.region_id"),
    }).
    Join(qb.LeftJoin, "locations l", qb.Condition{
        Field: "l.country_id",
        Opt:   qb.Eq,
        Value: qb.ValueField("c.country_id"),
    }).
    Where("c.country_id", qb.In, []string{"US", "UK", "CN"}).
    String()

// ------------- ALL | ANY -------------
sql = qb.QueryInstance().
    Select("employee_id", "first_name", "last_name", "salary").
    From("employees").
    Where("salary", qb.GrEqAll,
        qb.QueryInstance().
        Select("salary").
        From("employees").
        Where("department_id", qb.Eq, 8),
    ).
	OrderBy("salary", qb.Desc).
    String()

sql = qb.QueryInstance().
    Select("employee_id", "first_name", "last_name", "salary").
    From("employees").
    Where("salary", qb.GreaterAny,
        qb.QueryInstance().
        Select("AVG(salary)").
        From("employees").
        GroupBy("department_id"),
    ).
    OrderBy("first_name", qb.Asc).
    OrderBy("last_name", qb.Asc).
    String()

// ------------- EXISTS -------------
sql = qb.QueryInstance().
    Select("employee_id", "first_name", "last_name", "salary").
    From("employees", " e").
    Where(qb.FieldEmpty(""), qb.Exists,
        qb.QueryInstance().
        Select("1").
        From("dependents", "d").
        Where("d.employee_id", qb.Eq, qb.ValueField("e.employee_id")),
    ).
    OrderBy("first_name", qb.Asc).
    OrderBy("last_name", qb.Asc).
    String()

// ------------- BETWEEN | NOT BETWEEN -------------
sql = qb.QueryInstance().
    Select("employee_id", "first_name", "last_name", "salary").
    From("employees").
    Where("salary", qb.Between, qb.ValueBetween{
        Low:  9000,
        High: 12000,
    }).
    OrderBy("salary", qb.Asc).
    String()

sql = qb.QueryInstance().
    Select("employee_id", "first_name", "last_name", qb.FieldYear("hire_date")+" joined_year").
    From("employees").
    Where(qb.FieldYear("hire_date"),
        qb.Between, 
        qb.ValueBetween{
            Low:  1990,
            High: 1993,
        }).
    OrderBy("hire_date", qb.Asc).
    String()

// ------------- IN | NOT IN -------------
sql = qb.QueryInstance().
    Select("employee_id", "first_name", "last_name", "job_id").
    From("employees").
    Where("job_id", qb.In, []int{8, 9, 10}).
    OrderBy("job_id", qb.Asc).
    String()

sql = qb.QueryInstance().
    Select("employee_id", "first_name", "last_name", "job_id").
    From("employees").
    Where("job_id", qb.NotIn, []int{7, 8, 9}).
    OrderBy("job_id", qb.Asc).
    String()

// ------------- LIKE | NOT LIKE -------------
sql = qb.QueryInstance().
    Select("employee_id", "first_name", "last_name").
    From("employees").
    Where("first_name", qb.Like, "S%").
    Where("first_name", qb.NotLike, "Sh%").
    OrderBy("first_name", qb.Asc).
    String()

// ------------- IS NULL | IS NOT NUL -------------
sql = qb.QueryInstance().
    Select("employee_id", "first_name", "last_name", "phone_number").
    From("employees").
    Where("phone_number", qb.Null, nil).
    String()

sql = qb.QueryInstance().
    Select("employee_id", "first_name", "last_name", "phone_number").
    From("employees").
    Where("phone_number", qb.NotNull, nil).
    String()

// ------------- NOT -------------
sql = qb.QueryInstance().
    Select("employee_id", "first_name", "last_name", "salary").
    From("employees").
    Where("department_id", qb.Eq, 5).
    Where(qb.FieldNot("salary"), qb.Greater, 5000).
    String()

sql = qb.QueryInstance().
    Select("employee_id", "first_name", "last_name", "salary").
    From("employees").
    Where("salary", qb.NotBetween, qb.ValueBetween{Low: 3000, High: 5000}).
    String()

sql = qb.QueryInstance().
    Select("employee_id", "first_name", "last_name").
    From("employees", "e").
    Where(qb.FieldEmpty(""), qb.NotExists,
        qb.QueryInstance().
        Select("employee_id").
        From("dependents", "d").
        Where("d.employee_id", qb.Eq, qb.ValueField("e.employee_id")),
    ).
    String()
```

## UpdateBuilder
UpdateBuilder: UPDATE - updates data in a database

```go
import (
    qb "github.com/jivegroup/fluentsql"
)

// Simple
sql := qb.UpdateInstance().
    Update("employees").
    Set("first_name", "Steven").
    Set("last_name", "King").
    Where("employee_id", qb.Eq, 100).
    String()

// Complex
sql = qb.UpdateInstance().
    Update("employees").
    Set([]string{"first_name", "last_name", "salary", "department_id"}, []any{"Steven - Modified", "King - Modified", 25500, 11}).
    Where("employee_id", qb.Eq, 100).
    String()
```

## InsertBuilder
InsertBuilder: INSERT - inserts new data into a database

```go
import (
    qb "github.com/jivegroup/fluentsql"
)

// Insert multi-rows
sql := qb.InsertInstance().
    Insert("countries", "country_id", "country_name", "region_id").
    Row("VN", "Vietnam", 4).
    Row("VI", "Vieata", 4).
    Row("VM", "VieatMonda", 4).
	String()

// Insert from query
sql = qb.InsertInstance().
    Insert("countries_temp", "country_id", "country_name", "region_id").
    Query(qb.QueryInstance().
        Select("c.country_id", "c.country_name", "c.region_id").
        From("countries", "c").
        Where("c.country_id", qb.NotIn, []string{"VN", "VI", "VM"}),
    ).
    String()
```

## DeleteBuilder
DeleteBuilder: DELETE - deletes data from a database

```go
import (
    qb "github.com/jivegroup/fluentsql"
)

sql := qb.DeleteInstance().
    Delete("countries").
    Where("country_id", qb.Eq, "VN").
    WhereOr("country_id", qb.Eq, "VI").
    WhereOr("country_id", qb.Eq, "VM").
    String()
```
