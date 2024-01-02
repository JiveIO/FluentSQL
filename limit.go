package fluentsql

import "fmt"

// Limit type struct
type Limit struct {
	Limit  int
	Offset int
}

func (l Limit) String() string {
	return fmt.Sprintf(" LIMIT %d OFFSET %d", l.Limit, l.Offset)
}
