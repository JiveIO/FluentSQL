package fluentsql

import (
	"fmt"
	"strings"
)

// From type struct
type From struct {
	Table string
	Alias string
}

func (f From) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("FROM %s", f.Table))
	if f.Alias != "" {
		sb.WriteString(f.Alias)
	}

	return sb.String()
}
