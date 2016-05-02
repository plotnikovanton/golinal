package la

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Vector []float64

// Cunstruct new Vector
func NewVector(elems ...float64) Vector {
	return elems
}

// Returns random vector with given length
func NewRandomVector(length int) Vector {
	rand.Seed(time.Now().UTC().UnixNano())
	result := make(Vector, length)

	for i := 0; i < length; i++ {
		result[i] = rand.Float64()*2 - 1
	}

	return result
}

func (v Vector) String() string {
	buff := bytes.NewBufferString("[")
	for i, val := range v {
		buff.WriteString(strconv.FormatFloat(val, 'g', -1, 64))
		if i < len(v)-1 {
			buff.WriteRune(',')
		}
	}
	buff.WriteRune(']')
	return buff.String()
}

// Equals checks if given vectors is equal to current vector
func (a Vector) Equals(b Vector) bool {
	if len(a) != len(b) {
		return false
	} else {
		for i, val := range a {
			if val != b[i] {
				return false
			}
		}
		return true
	}
}

// Clones vector
func (v Vector) CloneVector() Vector {
	ret := make([]float64, len(v))
	copy(ret, v)
	return ret
}

// Rase an error if given Vectors got not same length
func AssertVectorsLength(a, b Vector) {
	if len(a) != len(b) {
		panic(fmt.Sprintf("Length of given vectors is not same %v and %v", len(a), len(b)))
	}
}

// Euqlidian norm of vector
func (v Vector) Len() float64 {
	var sum float64
	for _, val := range v {
		sum += val * val
	}
	return math.Sqrt(sum)
}

// Scalar multiply vector with given constant
func (v Vector) ScalMul(s float64) Vector {
	result := make([]float64, len(v))
	for i, val := range v {
		result[i] = val * s
	}
	return result
}

// Negate vector
// It is alias for
//		vector.ScalMul(-1.)
func (v Vector) Neg() Vector {
	return v.ScalMul(-1.)
}

// Returns normalization of the vector
func (v Vector) Unit() Vector {
	return v.ScalMul(1 / v.Len())
}

// Sum of two vectors
func (a Vector) Add(b Vector) Vector {
	AssertVectorsLength(a, b)
	result := make([]float64, len(a))
	for i, val := range a {
		result[i] = val + b[i]
	}
	return result
}

// Returns substractions of two vectors
func (a Vector) Sub(b Vector) Vector {
	AssertVectorsLength(a, b)
	result := make([]float64, len(a))
	for i, val := range a {
		result[i] = val - b[i]
	}
	return result
}

// DotProduct
func (a Vector) DotProd(b Vector) Vector {
	AssertVectorsLength(a, b)
	result := make([]float64, len(a))
	for i, val := range a {
		result[i] = val * b[i]
	}
	return result
}

// Returns inner product for pair of vectors
func (a Vector) InnerProd(b Vector) float64 {
	AssertVectorsLength(a, b)
	res := 0.
	for i, val := range a {
		res += val * b[i]
	}
	return res
}

// CrossProduct for Vectors length 3
func (a Vector) CrossProd(b Vector) Vector {
	AssertVectorsLength(a, b)
	if len(a) != 3 {
		panic(fmt.Sprintf("Given vectors length should be '3', but got", len(a)))
	}

	return NewVector(
		a[1]*b[2]-b[1]*a[2],
		-1*(a[0]*b[2]-b[0]*a[2]),
		a[0]*b[1]-b[0]*a[1],
	)
}
