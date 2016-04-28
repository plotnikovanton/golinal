package la

import (
	"fmt"
)

type Column []Vector

// Creates Column filled with clones of given Vector
func (v Vector) NewColumn(length int) Column {
	result := make(Column, length)
	for i := 0; i < length; i++ {
		result[i] = v.CloneVector()
	}
	return result
}

// Rase an error if given Columns got not same length
func AssertColumnLength(a, b Column) {
	if len(a) != len(b) {
		panic(fmt.Sprintf("Length of given columns is not same %v and %v", len(a), len(b)))
	}
}

// Produces new Column by applying an
// Vector -> Vector -> Vector function value by value of given columns
func (a Column) WithColumn(f func(Vector, Vector) Vector, b Column) Column {
	AssertColumnLength(a, b)
	result := make(Column, len(a))
	for i, val := range a {
		result[i] = f(val, b[i])
	}
	return result
}

// Apply Vector -> Vector function to each element of Coulmn
// and produce new column
func (a Column) ToColumn(f func(Vector) Vector) Column {
	result := make(Column, len(a))
	for i, val := range a {
		result[i] = f(val)
	}
	return result
}
