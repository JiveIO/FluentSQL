package fluentsql

import (
	"fmt"
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
	// Opt condition operators =, <>, >, <, >=, <=, LIKE, IN, NOT IN
	Opt WhereOpt
	// Value data of condition
	Value any
	// AndOr Combination type with previous condition AND, OR. Default is AND
	AndOr WhereAndOr
	// Group condition A=1 AND (T=2 OR M=3)
	Group []Condition
}

type WhereAndOr int

const (
	And WhereAndOr = iota
	Or
)

func (c Condition) andOr() string {
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
	Greater
	Lesser
	GrEq
	LeEq
	Like
	In
	NotIn
)

func (c Condition) opt() string {
	var sign string

	switch c.Opt {
	case Eq:
		sign = "="
	case NotEq:
		sign = "<>"
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
	}

	return sign
}

func (w Where) String() string {
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

func (c Condition) String() string {
	if len(c.Group) > 0 {
		var conditions []string

		for _, cond := range c.Group {
			var _condition = cond.String()

			if cond.AndOr == Or && len(conditions) > 0 {
				_orCondition := fmt.Sprint(" OR ", _condition)

				last := len(conditions) - 1

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

	if _, ok := c.Value.(string); ok {
		val := c.Value

		if c.Opt == Like {
			val = fmt.Sprint("%", val, "%")
		}

		//if c.Opt == In {
		//	val = fmt.Sprint("(", strings.Join(val, ", "), ")")
		//}

		return fmt.Sprintf("%s %s '%s'", c.Field, c.opt(), val)
	}

	return fmt.Sprintf("%s %s %s", c.Field, c.opt(), c.Value)
}
