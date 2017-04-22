package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

// B (Boolean)
type B bool

// IsType for B
func (b B) IsType() {}

// IsValue for B
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
func (v VEC) IsType() {}

// IsValue for VEC
func (v VEC) IsValue() {}

// String output for VEC
func (v VEC) String() string {
	return fmt.Sprintf("%v", v.Vector)
}

// Get a location within the VEC
func (v VEC) Get(loc int) interfaces.Type {
	return v.Vector[loc]
}

func (v VEC) count() int {
	return len(v.Vector)
}

// S provides a type for string values
type S string

// IsType for S
func (s S) IsType() {}

// IsValue for S
func (s S) IsValue() {}

//String output for S
func (s S) String() string {
	return string(s)
}

//CompareTo for String
func (s S) CompareTo(o interfaces.Comparable) (int, error) {
	panic("not implemented")
}

// NIL generally acts as a return type when a function performs a side effect
type NIL struct{}

// IsType for NIL
func (n NIL) IsType() {}

// IsValue for NIL
func (n NIL) IsValue() {}

// String output for NIL
func (n NIL) String() string {
	return "<NIL>"
}

// NILL should be used for NIL values rather than creating a new NIL each time
var NILL = NIL{}
