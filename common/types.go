package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

// I (Integer)
type I int

// IsType for I
func (i I) IsType()  {}
func (i I) IsValue() {}

// String representation of I
func (i I) String() string {
	return fmt.Sprintf("%d", i.Int())
}

// Int unboxes an int from Int
func (i I) Int() int {
	return int(i)
}

// Equals checks equality with another item of type Type
func (i I) Equals(o interfaces.Equalable) interfaces.Value {
	if other, ok := o.(I); ok {
		return B(i.Int() == other.Int())
	}
	return B(false)
}

// CompareTo compares one I to another I and returns -1, 0 or 1
func (i I) CompareTo(o interfaces.Comparable) (int, error) {
	if other, ok := o.(I); ok {
		if i.Int() < other.Int() {
			return -1, nil
		} else if i.Int() == other.Int() {
			return 0, nil
		}
		return 1, nil
	}
	if other, ok := o.(F); ok {
		f := F(i.Int())
		return f.CompareTo(other)
	}
	return 0, fmt.Errorf("CompareTo : Cannot compare %v to %v", i, o)
}

// B (Boolean)
type B bool

// IsType for B
func (b B) IsType()  {}
func (b B) IsValue() {}

// String for B
func (b B) String() string {
	return fmt.Sprintf("%t", b.Bool())
}

// Bool returns a bool representation of B
func (b B) Bool() bool {
	return bool(b)
}

// Equals checks equality with another item of type Type
func (b B) Equals(o interfaces.Equalable) interfaces.Value {
	if other, ok := o.(B); ok {
		return B(b.Bool() == other.Bool())
	}
	return B(false)
}

// VEC is a Vector or array
type VEC struct {
	Vector []interfaces.Type
}

// IsType for VEC
func (v VEC) IsType()  {}
func (v VEC) IsValue() {}

// String output for VEC
func (v VEC) String() string {
	return fmt.Sprintf("%v", v.Vector)
}

// Get a location within the VEC
func (v VEC) Get(loc int) interfaces.Type {
	return v.Vector[loc]
}

func (v VEC) Count() int {
	return len(v.Vector)
}

// S provides a type for string values
type S string

// IsType for S
func (s S) IsType()  {}
func (s S) IsValue() {}

//String output for S
func (s S) String() string {
	return string(s)
}

func (s S) CompareTo(o interfaces.Comparable) (int, error) {
	panic("not implemented")
}

// NIL generally acts as a return type when a function performs a side effect
type NIL struct{}

// IsType for NIL
func (n NIL) IsType()  {}
func (n NIL) IsValue() {}

// String output for NIL
func (n NIL) String() string {
	return "<NIL>"
}

var NILL = NIL{}

// F (Float)
type F float64

// IsType for F
func (f F) IsType()  {}
func (f F) IsValue() {}

// String representation of F
func (f F) String() string {
	return fmt.Sprintf("%f", f.float())
}

// float unboxes a float from F
func (f F) float() float64 {
	return float64(f)
}

// Equals checks equality with another item of type Type
func (f F) Equals(o interfaces.Equalable) interfaces.Value {
	if other, ok := o.(F); ok {
		return B(f.float() == other.float())
	}
	return B(false)
}

// CompareTo compares one F to another F and returns -1, 0 or 1
func (f F) CompareTo(o interfaces.Comparable) (int, error) {
	if other, ok := o.(F); ok {
		if f.float() < other.float() {
			return -1, nil
		} else if f.float() == other.float() {
			return 0, nil
		}
		return 1, nil
	}
	if other, ok := o.(I); ok {
		return f.CompareTo(F(other.Int()))
	}
	return 0, fmt.Errorf("CompareTo : Cannot compare %v to %v", f, o)
}
