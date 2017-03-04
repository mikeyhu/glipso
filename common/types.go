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
func (i I) Equals(o interfaces.Equalable) interfaces.Type {
	if other, ok := o.(I); ok {
		return B(i.Int() == other.Int())
	}
	return B(false)
}

// CompareTo compares one I to another I and returns -1, 0 or 1
func (i I) CompareTo(o interfaces.Comparable) int {
	if other, ok := o.(I); ok {
		if i.Int() < other.Int() {
			return -1
		} else if i.Int() == other.Int() {
			return 0
		}
		return 1
	}
	panic(fmt.Sprintf("CompareTo : Cannot compare %v to %v", i, o))
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
func (b B) Equals(o interfaces.Equalable) interfaces.Type {
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
	return fmt.Sprintf("P(%v %v)", p.head, p.tail)
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

// Evaluate resolves a REF to something in scope
func (r REF) Evaluate(sco interfaces.Scope) interfaces.Type {
	if resolved, ok := sco.ResolveRef(r); ok {
		if evaluatable, ok := resolved.(*EXP); ok {
			return evaluatable //.Evaluate(sco)
		}
		return resolved
	}
	panic(fmt.Sprintf("Unable to resolve REF('%v')\n", r))
}

// EvaluateToRef resolves a REF down to another REF
func (r REF) EvaluateToRef(sco interfaces.Scope) REF {
	resolved, ok := sco.ResolveRef(r)
	if ok {
		if resolvedRef, ok := resolved.(REF); ok {
			return resolvedRef.EvaluateToRef(sco)
		}
		return r
	}
	return r
}

// LAZYP (Lazily evaluated Pair)
// head operates like Pair, tail should be an expression that returns another LAZYP
type LAZYP struct {
	head interfaces.Type
	tail interfaces.Evaluatable
}

// IsType for LAZYP
func (l LAZYP) IsType() {}

// String representation of LAZYP
func (l LAZYP) String() string {
	return fmt.Sprintf("LAZYP(%v %v)", l.head, l.tail)
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

// VEC is a Vector or array
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

// FN acts as storage for a reusable Function by storing a set of arguments to a function and the function expression itself
type FN struct {
	Arguments  VEC
	Expression *EXP
}

// IsType for FN
func (f FN) IsType() {}

// String output for FN
func (f FN) String() string {
	return fmt.Sprintf("FN(%v, %v)", f.Arguments, f.Expression)
}

// S provides a type for string values
type S string

// IsType for S
func (s S) IsType() {}

//String output for S
func (s S) String() string {
	return string(s)
}
