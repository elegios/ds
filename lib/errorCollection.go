package lib

import (
	"strings"
)

type ErrorCollection []error

func (e *ErrorCollection) Error() string {
	strs := make([]string, len(*e), len(*e))
	for i, err := range *e {
		strs[i] = err.Error()
	}
	return strings.Join(strs, "\n")
}
