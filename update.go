package fluentsql

import (
	"fmt"
	"strings"
)

// Update clause
type Update struct {
	Table any
}

func (u *Update) String() string {
	return fmt.Sprintf("UPDATE %s", u.Table)
}

type UpdateItem struct {
	// Field name of column type string
	Field any
	// Value data of string, int, ValueField, QueryBuilder
	Value any
}

func (s *UpdateItem) String() string {
	if valueQueryBuilder, ok := s.Value.(*QueryBuilder); ok { // Value type is QueryBuilder.
		return fmt.Sprintf("%s = (%v)", s.Field, valueQueryBuilder)
	}

	return fmt.Sprintf("%s = %s", s.Field, s.Value)
}

type UpdateSet struct {
	Items []UpdateItem
}

func (s *UpdateSet) Append(field, value any) {
	s.AppendItems(UpdateItem{
		Field: field,
		Value: value,
	})
}

func (s *UpdateSet) AppendItems(items ...UpdateItem) {
	s.Items = append(s.Items, items...)
}

func (s *UpdateSet) String() string {
	var setColumns []string

	for _, item := range s.Items {
		setColumns = append(setColumns, item.String())
	}

	return strings.Join(setColumns, ", ")
}
