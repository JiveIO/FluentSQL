package main

import (
	"fmt"
	qb "github.com/JiveIO/FluentSQL"
	"time"
)

func main() {
	rb := qb.NewQueryBuilder().
		Select("first_name", "last_name", "email").
		From("users", "usr").
		Where("usr.first_name", qb.Like, "john").
		Where("usr.email", qb.Eq, "john@mail.com").
		WhereGroup(func(wb qb.QueryBuilder) *qb.QueryBuilder {
			return wb.WhereOr("usr.created_at", qb.Greater, time.Now().Format("2006-01-12")).
				WhereOr("usr.update_at", qb.GrEq, time.Now().Format("2006-01-12"))
		}).
		Limit(10, 0)

	fmt.Print("SQL>", rb.String())
}
