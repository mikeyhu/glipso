package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

// I (Integer)
type I int

// IsType for I
func (i I) IsType() {}

// String representation of I
func (i I) String() string {
	return fmt.Sprintf("%d", i.Int())
}

// Int unboxes an int from Int
func (i I) Int() int {
	return int(i)
}

// Equals checks equality with another item of type Type
func (i I) Equals(o interfaces.Type) B {
	if other, ok := o.(I); ok {
		return B(i.Int() == other.Int())
	}
	return B(false)
}

// B (Boolean)
type B bool

// IsType for B
func (b B) IsType() {}

// String for B
func (b B) String() string {
	return fmt.Sprintf("%t", b.Bool())
}

// Bool returns a bool representation of B
func (b B) Bool() bool {
	return bool(b)
}

// Equals checks equality with another item of type Type
func (b B) Equals(o interfaces.Type) B {
	if other, ok := o.(B); ok {
		return B(b.Bool() == other.Bool())
	}
	return B(false)
}

// P (PAIR)
type P struct {
	head interfaces.Type
	tail *P
}

// IsType for P
func (p P) IsType() {}

// String representation of P
func (p P) String() string {
	return fmt.Sprintf("￿`(%v %v)", p.head, p.tail)
}

// Iterate not supported yet
func (p P) Iterate(sco interfaces.Scope) interfaces.Iterable {
	panic("Iterate called on PAIR")
}

// ToSlice returns an array from a P
func (p P) ToSlice(sco interfaces.Scope) []interfaces.Type {
	slice := []interfaces.Type{p.head}
	tail := p.tail
	for {
		if tail == nil {
			return slice
		}
		slice = append(slice, tail.head)
		tail = tail.tail
	}
}

// REF (Reference)
// symbol for something in scope, variable or function
type REF string

// IsType for REF
func (r REF) IsType() {}

// String representation of a REF
func (r REF) String() string {
	return fmt.Sprintf("%v", string(r))
}
func (r REF) Evaluate(sco interfaces.Scope) interfaces.Type {
	return sco.ResolveRef(r)
}

// LAZYP (Lazily evaluated Pair)
// head operates like Pair, tail should be an expression that returns another LAZYP
type LAZYP struct {
	head interfaces.Type
	tail *EXP
}

// IsType for LAZYP
func (l LAZYP) IsType() {}

// String representation of LAZYP
func (l LAZYP) String() string {
	return fmt.Sprintf("￿`(%v %v)", l.head, l.tail)
}

// Iterate will evaluate the tail of the LAZYP
func (l LAZYP) Iterate(sco interfaces.Scope) interfaces.Iterable {
	if nextIter, ok := l.tail.Evaluate(sco).(LAZYP); ok {
		return nextIter
	}
	panic(fmt.Sprintf("Iterate : expected an LAZYP, got %v", l))
}

// ToSlice converts a LAZYP to a slice by iterating through it
func (l LAZYP) ToSlice(sco interfaces.Scope) []interfaces.Type {
	slice := []interfaces.Type{}
	next := l
	for {
		slice = append(slice, next.head)
		if next.tail == nil {
			return slice
		}
		next = next.Iterate(sco.NewChildScope()).(LAZYP)
	}
}

// VEC
// is a Vector or array
type VEC struct {
	Vector []interfaces.Type
}

// IsType for VEC
func (v VEC) IsType() {}

// String output for VEC
func (v VEC) String() string {
	return fmt.Sprintf("%v", v.Vector)
}

// Get a location within the VEC
func (v VEC) Get(loc int) interfaces.Type {
	return v.Vector[loc]
}

// FN
// acts as storage for a reusable Function by storing a set of arguments to a function and the function expression itself
type FN struct {
	Arguments  VEC
	Expression EXP
}

// IsType for FN
func (f FN) IsType() {}

// String output for FN
func (f FN) String() string {
	return fmt.Sprintf("FN %v %v", f)
}
