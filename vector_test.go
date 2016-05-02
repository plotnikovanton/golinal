package la

import (
	"math"
	"testing"
)

func TestNewVector(t *testing.T) {
	v := NewVector(1, 2, 4)
	if len(v) != 3 {
		t.Errorf("Expecting vector size %d, but got %d", 3, len(v))
	}
	if v[0] != 1 || v[1] != 2 || v[2] != 4 {
		t.Errorf(
			"Expecting vector values 1, 2, 4, but got %d %d %d",
			v[0], v[1], v[2],
		)
	}
}

func TestVectorString(t *testing.T) {
	a := NewVector(1, 2, 3)
	if a.String() != "[1,2,3]" {
		t.Errorf("Expecting [1,2,3], but got %s", a.String())
	}
}

func TestEquals(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(1, 2, 3)
	c := NewVector(1, 2, 4)
	if !(a.Equals(b) && b.Equals(a)) {
		t.Errorf(
			"Expecting that vectors %s and %s are equals, but not",
			a.String(), b.String(),
		)
	}
	if a.Equals(c) || c.Equals(a) {
		t.Errorf(
			"Expecting that vectors %s and %s are diferent, but not",
			a.String(), c.String(),
		)
	}
}

func TestCloneVector(t *testing.T) {
	v := NewVector(1, 2, 3)
	v_clone := v.CloneVector()
	if !v.Equals(v_clone) {
		t.Errorf(
			"Expecting that vector %s and clone %s are same at creation, but not",
			v.String(), v_clone.String(),
		)
	}

	v_clone[0] = 42
	if v.Equals(v_clone) {
		t.Errorf(
			"Expecting that vector %s and clone %s are diferent after changing, but not",
			v.String(), v_clone.String(),
		)
	}
}

func TestVectorLen(t *testing.T) {
	v := NewVector(1, 2, -2)
	if v.Len() != 3. {
		t.Errorf(
			"Expecting that vector length is 3, but got",
			v.Len(),
		)
	}
}

func TestScalMul(t *testing.T) {
	v := NewVector(1, 2, -3)
	result := NewVector(-2., -4., 6.)
	if !v.ScalMul(-2.).Equals(result) {
		t.Errorf(
			"Expecting that scalar mult would equal %s, but got %s ",
			result.String(), v.ScalMul(-2.).String(),
		)
	}
}

func TestNeg(t *testing.T) {
	v := NewVector(1, 2, -3)
	result := NewVector(-1, -2, 3)
	if !v.Neg().Equals(result) {
		t.Errorf(
			"Expecting that negate would equal %s, but got %s ",
			result.String(), v.Neg().String(),
		)
	}
}

func TestUnit(t *testing.T) {
	v := NewVector(2, -2, 2)
	result := NewVector(1/math.Sqrt(3), -1/math.Sqrt(3), 1/math.Sqrt(3))
	if !v.Unit().Equals(result) {
		t.Errorf(
			"Expecting that utin vector would equal %s, but got %s ",
			result.String(), v.Unit().String(),
		)
	}
}

func TestVectorAdd(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(-3, -3, -3)
	actual := a.Add(b)
	except := NewVector(-2, -1, 0)
	if !actual.Equals(except) {
		t.Errorf(
			"Expecting that %s adding %s would equal %s, but got %s ",
			a.String(), b.String(), except.String(), actual.String(),
		)
	}
}

func TestVectorSub(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(-3, -3, -3)
	actual := a.Sub(b)
	except := NewVector(4, 5, 6)
	if !actual.Equals(except) {
		t.Errorf(
			"Expecting that %s substruct %s would equal %s, but got %s ",
			a.String(), b.String(), except.String(), actual.String(),
		)
	}
}

func TestVectorDotProduct(t *testing.T) {
	a := NewVector(1, -1, 2)
	b := NewVector(1, 2, 2)
	actual := a.DotProd(b)
	except := NewVector(1, -2, 4)
	if !actual.Equals(except) {
		t.Errorf(
			"Expecting that %s dot product %s would equal %s, but got %s ",
			a.String(), b.String(), except.String(), actual.String(),
		)
	}
}

func TestVectorInnerProduct(t *testing.T) {
	a := NewVector(1, -1, 2)
	b := NewVector(1, 2, 2)
	actual := a.InnerProd(b)
	except := 3.
	if actual != except {
		t.Errorf(
			"Expecting that %s dot product %s would equal %s, but got %s ",
			a.String(), b.String(), except, actual,
		)
	}
}

func TestVectorCrossProduct(t *testing.T) {
	a := NewVector(1, -1, 2)
	b := NewVector(1, 2, 2)
	actual := a.CrossProd(b)
	except := NewVector(-6, 0, 3)
	if !actual.Equals(except) {
		t.Errorf(
			"Expecting that %s cross product %s would equal %s, but got %s ",
			a.String(), b.String(), except.String(), actual.String(),
		)
	}
}
