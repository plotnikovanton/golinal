package la

import (
	"fmt"
)

type Column []Vector

// Rase an error if given Columns got not same length
func AssertColumnLength(a, b Column) {
	if len(a) != len(b) {
		panic(fmt.Sprintf("Length of given columns is not same %v and %v", len(a), len(b)))
	}
}

// Produces new Column by applying an
// Vector -> Vector -> Vector function value by value of given columns
func (a Column) ToColumn(f func(Vector, Vector) Vector, b Column) Column {
	AssertColumnLength(a, b)
	result := make(Column, len(a))
	for i, val := range a {
		result[i] = f(val, b[i])
	}
	return result
}
