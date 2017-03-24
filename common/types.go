package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

// I (Integer)
type I int

// IsType for I
func (i I) IsType()   {}
func (i I) IsResult() {}

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
func (b B) IsType()   {}
func (b B) IsResult() {}

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

// P (PAIR)
type P struct {
	head interfaces.Value
	tail interfaces.Iterable
}

// IsType for P
func (p P) IsType()   {}
func (p P) IsResult() {}

// String representation of P
func (p P) String() string {
	return fmt.Sprintf("P(%v %v)", p.head, p.tail)
}

// Head returns the Head or the P
func (p P) Head() interfaces.Value {
	return p.head
}

// HasTail returns bool showing whether the P has a tail
func (p P) HasTail() bool {
	return p.tail != ENDED
}

// Iterate not supported yet
func (p P) Iterate(sco interfaces.Scope) interfaces.Iterable {
	if p.HasTail() {
		return p.tail
	}
	return ENDED
}

// ToSlice returns an array from a P
func (p P) ToSlice(sco interfaces.Scope) []interfaces.Type {
	slice := []interfaces.Type{}
	var tail interfaces.Iterable = p
	for {
		if tail != ENDED {
			slice = append(slice, tail.Head())
			if !tail.HasTail() {
				return slice
			}
			tail = tail.Iterate(sco)
		}
	}
	return slice
}

// REF (Reference)
// symbol for something in scope, variable or function
type REF string

// IsType for REF
func (r REF) IsType() {}

// String representation of a REF
func (r REF) String() string {
	return string(r)
}

// Evaluate resolves a REF to something in scope
func (r REF) Evaluate(sco interfaces.Scope) interfaces.Value {
	if DEBUG {
		env := sco.(*Environment)
		fmt.Printf("%v being looked up in scope %v:\n", r, env.id)
		env.DisplayEnvironment()
	}
	if resolved, ok := sco.ResolveRef(r); ok {
		return resolved
	}
	panic(fmt.Sprintf("Unable to resolve REF('%v')\n", r))
}

// LAZYP (Lazily evaluated Pair)
// head operates like Pair, tail should be an expression that returns another LAZYP
type LAZYP struct {
	head interfaces.Value
	tail interfaces.Evaluatable
}

// IsType for LAZYP
func (l LAZYP) IsType()   {}
func (l LAZYP) IsResult() {}

// String representation of LAZYP
func (l LAZYP) String() string {
	return fmt.Sprintf("LAZYP(%v %v)", l.head, l.tail)
}

// Head returns the Head of the LAZYP
func (l LAZYP) Head() interfaces.Value {
	return l.head
}

// HasTail returns true if the LAZYP has an evaluatable tail
func (l LAZYP) HasTail() bool {
	return l.tail != nil
}

// Iterate will evaluate the tail of the LAZYP
func (l LAZYP) Iterate(sco interfaces.Scope) interfaces.Iterable {
	if nextIter, ok := l.tail.Evaluate(sco).(interfaces.Iterable); ok {
		return nextIter
	}
	panic(fmt.Sprintf("Iterate : expected an LAZYP, got %v", l))
}

// ToSlice converts a LAZYP to a slice by iterating through it
func (l LAZYP) ToSlice(sco interfaces.Scope) []interfaces.Type {
	slice := []interfaces.Type{}
	var next interfaces.Iterable = l
	for {
		slice = append(slice, next.Head())

		if !next.HasTail() {
			return slice
		}
		next = next.Iterate(sco.NewChildScope()).(interfaces.Iterable)
	}
}

// VEC is a Vector or array
type VEC struct {
	Vector []interfaces.Type
}

// IsType for VEC
func (v VEC) IsType()   {}
func (v VEC) IsResult() {}

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

// FN acts as storage for a reusable Function by storing a set of arguments to a function and the function expression itself
type FN struct {
	Arguments  VEC
	Expression interfaces.Evaluatable
}

// IsType for FN
func (f FN) IsType()   {}
func (f FN) IsResult() {}

// String output for FN
func (f FN) String() string {
	return fmt.Sprintf("FN(%v, %v)", f.Arguments, f.Expression)
}

func (f FN) Apply(arguments []interfaces.Type, env interfaces.Scope) interfaces.Value {
	if len(f.Arguments.Vector) != len(arguments) {
		panic("Invalid number of arguments")
	}
	fnenv := env.NewChildScope()
	for i, v := range f.Arguments.Vector {
		fnenv.CreateRef(v.(REF), evaluateToResult(arguments[i], env))
	}
	return f.Expression.Evaluate(fnenv)
}

// S provides a type for string values
type S string

// IsType for S
func (s S) IsType()   {}
func (s S) IsResult() {}

//String output for S
func (s S) String() string {
	return string(s)
}

// NIL generally acts as a return type when a function performs a side effect
type NIL struct{}

// IsType for NIL
func (n NIL) IsType()   {}
func (n NIL) IsResult() {}

// String output for NIL
func (n NIL) String() string {
	return "<NIL>"
}

var NILL = NIL{}

// END acts as the end of a list
type END struct{}

func (e END) IsType()   {}
func (e END) IsResult() {}

func (e END) String() string {
	return "<END>"
}
func (e END) Head() interfaces.Value {
	return NILL
}
func (e END) HasTail() bool {
	return false
}
func (e END) Iterate(interfaces.Scope) interfaces.Iterable {
	panic("END : Iterate called on END")
}
func (e END) ToSlice(interfaces.Scope) []interfaces.Type {
	return []interfaces.Type{}
}

var ENDED = END{}
