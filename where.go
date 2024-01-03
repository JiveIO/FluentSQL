package fluentsql

import (
	"fmt"
	"reflect"
	"strings"
)

// Where type struct
type Where struct {
	Conditions []Condition
}

// Condition type struct
type Condition struct {
	// Field name of Column
	Field string
	// Opt condition operators =, <>, >, <, >=, <=, LIKE, IN, NOT IN, BETWEEN
	Opt WhereOpt
	// Value data of condition
	Value any
	// AndOr Combination type with previous condition AND, OR. Default is AND
	AndOr WhereAndOr
	// Group conditions in parentheses `()`.
	Group []Condition
}

type WhereAndOr int

const (
	And WhereAndOr = iota
	Or
)

func (c *Condition) andOr() string {
	var sign string

	switch c.AndOr {
	case Or:
		sign = "OR"
	case And:
		sign = "AND"
	}

	return sign
}

type WhereOpt int

const (
	Eq WhereOpt = iota
	NotEq
	Diff
	Greater
	Lesser
	GrEq
	LeEq
	Like
	In
	NotIn
	Between
	NotBetween
	Null
	NotNull
	Exists
	NotExists
	EqAny
	NotEqAny
	DiffAny
	GreaterAny
	LesserAny
	GrEqAny
	LeEqAny
	EqAll
	NotEqAll
	DiffAll
	GreaterAll
	LesserAll
	GrEqAll
	LeEqAll
)

// BetweenValue for WhereOpt.Between or WhereOpt.NotBetween
type BetweenValue struct {
	Low  any
	High any
}

func (v BetweenValue) String() string {
	if _, ok := v.Low.(string); ok {
		// hire_date BETWEEN '1999-01-01' AND '2000-12-31'
		return fmt.Sprintf("'%v' AND '%v'", v.Low, v.High)
	}

	//  salary NOT BETWEEN 2500 AND 2900
	return fmt.Sprintf("%v AND %v", v.Low, v.High)
}

// FieldValue to handle condition `WHERE c.column <WhereOpt> c.column_1`
//
// So, When build condition Where("d.employee_id", qb.Eq, qb.FieldValue("e.employee_id")) to keep SQL string as
// d.employee_id = e.employee_id instead of
// d.employee_id = 'e.employee_id'
type FieldValue string

func (v FieldValue) String() string {
	return string(v)
}

func (c *Condition) opt() string {
	var sign string

	switch c.Opt {
	case Eq:
		sign = "="
	case NotEq:
		sign = "<>"
	case Diff:
		sign = "!="
	case Greater:
		sign = ">"
	case Lesser:
		sign = "<"
	case GrEq:
		sign = ">="
	case LeEq:
		sign = "<="
	case Like:
		sign = "LIKE"
	case In:
		sign = "IN"
	case NotIn:
		sign = "NOT IN"
	case Between:
		sign = "BETWEEN"
	case NotBetween:
		sign = "NOT BETWEEN"
	case Null:
		sign = "IS NULL"
	case NotNull:
		sign = "IS NOT NULL"
	case Exists:
		sign = "EXISTS"
	case NotExists:
		sign = "NOT EXISTS"
	case EqAny:
		sign = "= ANY"
	case NotEqAny:
		sign = "<> ANY"
	case DiffAny:
		sign = "!= ANY"
	case GreaterAny:
		sign = "> ANY"
	case LesserAny:
		sign = "< ANY"
	case GrEqAny:
		sign = ">= ANY"
	case LeEqAny:
		sign = "<= ANY"
	case EqAll:
		sign = "= ALL"
	case NotEqAll:
		sign = "<> ALL"
	case DiffAll:
		sign = "!= ALL"
	case GreaterAll:
		sign = "> ALL"
	case LesserAll:
		sign = "< ALL"
	case GrEqAll:
		sign = ">= ALL"
	case LeEqAll:
		sign = "<= ALL"
	}

	return sign
}

func (c *Condition) String() string {
	// Handle group conditions
	if len(c.Group) > 0 {
		var conditions []string

		for _, cond := range c.Group {
			var _condition = cond.String()

			if cond.AndOr == Or && len(conditions) > 0 {
				_orCondition := fmt.Sprint(" OR ", _condition)

				last := len(conditions) - 1

				// OR with previous condition
				conditions[last] = conditions[last] + _orCondition
			} else {
				conditions = append(conditions, _condition)
			}
		}

		// No WHERE condition
		if len(conditions) == 0 {
			return ""
		}

		return fmt.Sprintf("(%s)", strings.Join(conditions, " AND "))
	}

	// WHERE Address IS NULL
	// WHERE Address IS NOT NULL
	if c.Opt == Null || c.Opt == NotNull {
		return fmt.Sprintf("%s %s", c.Field, c.opt())
	}

	// WHERE Country IN ('Germany', 'France', 'UK')
	// WHERE Age NOT IN (12, 31, 21)
	if c.Opt == In || c.Opt == NotIn {
		// Type of value
		typ := reflect.TypeOf(c.Value)

		if typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
			valuesStr := ""
			if values, ok := c.Value.([]string); ok {
				valuesStr = "'" + strings.Join(values, "', '") + "'"
			}
			if values, ok := c.Value.([]int); ok {
				valuesStr = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(values)), ", "), "[]")
			}

			return fmt.Sprintf("%s %s (%s)", c.Field, c.opt(), valuesStr)
		}
	}

	// WHERE Price BETWEEN 10 AND 20;
	// WHERE ProductName BETWEEN 'Carnation Tigers' AND 'Mozzarella di Giovanni'
	// WHERE Price NOT BETWEEN 10 AND 20;
	// WHERE ProductName NOT BETWEEN 'Carnation Tigers' AND 'Mozzarella di Giovanni'
	// WHERE Price BETWEEN 10 AND 20
	if c.Opt == Between || c.Opt == NotBetween {
		return fmt.Sprintf("%s %s %s", c.Field, c.opt(), c.Value.(BetweenValue).String())
	}

	// WHERE salary = (SELECT DISTINCT salary FROM employees ORDER BY salary DESC LIMIT 1 , 1);
	// WHERE CustomerID IN (SELECT CustomerID FROM Orders);
	// WHERE CustomerID NOT IN (SELECT CustomerID FROM Orders);
	// WHERE EXISTS (SELECT ProductName FROM Products);
	// WHERE NOT EXISTS (SELECT ProductName FROM Products);
	// WHERE ProductID = ANY (SELECT ProductID FROM OrderDetails WHERE Quantity = 10);
	// WHERE ProductID > ALL (SELECT ProductID FROM OrderDetails WHERE Quantity = 10);
	if _, ok := c.Value.(*QueryBuilder); ok { // Column type is a complex query.
		selectQuery := c.Value.(*QueryBuilder).String()

		if c.Opt == Eq || c.Opt == In || c.Opt == NotIn || c.Opt == Exists || c.Opt == NotExists ||
			c.Opt == EqAny || c.Opt == NotEqAny || c.Opt == DiffAny || c.Opt == GreaterAny || c.Opt == LesserAny || c.Opt == GrEqAny || c.Opt == LeEqAny ||
			c.Opt == EqAll || c.Opt == NotEqAll || c.Opt == DiffAll || c.Opt == GreaterAll || c.Opt == LesserAll || c.Opt == GrEqAll || c.Opt == LeEqAll {
			return fmt.Sprintf("%s %s (%s)", c.Field, c.opt(), selectQuery)
		}
	}

	if _, ok := c.Value.(string); ok {
		return fmt.Sprintf("%s %s '%v'", c.Field, c.opt(), c.Value)
	}

	return fmt.Sprintf("%s %s %v", c.Field, c.opt(), c.Value)
}

func (w *Where) String() string {
	var conditions []string

	if len(w.Conditions) > 0 {
		for _, cond := range w.Conditions {
			var _condition = cond.String()

			if cond.AndOr == Or && len(conditions) > 0 {
				_orCondition := fmt.Sprint(" OR ", _condition)

				last := len(conditions) - 1

				conditions[last] = conditions[last] + _orCondition
			} else {
				conditions = append(conditions, _condition)
			}
		}
	}

	// No WHERE condition
	if len(conditions) == 0 {
		return ""
	}

	return fmt.Sprintf("WHERE %s", strings.Join(conditions, " AND "))
}
