package fluentsql

import (
	"fmt"
	"strings"
)

// joinSlice joins a slice of any type into a single string with a specified separator.
//
// Parameters:
//   - values: []T - A slice of any type to be joined into a string.
//   - separator: string - A string to be used as the separator between elements in the slice.
//
// Returns:
//   - string: A string representation of the slice with elements separated by the given separator.
//
// Variables:
//   - values: The input slice whose elements are to be joined.
//   - separator: The string used to separate the elements in the joined result.
func joinSlice[T any](values []T, separator string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(values)), fmt.Sprintf("%s ", separator)), "[]")
}
