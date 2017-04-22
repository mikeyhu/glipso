package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

// P (PAIR)
type P struct {
	head interfaces.Value
	tail interfaces.Iterable
}

// IsType for P
func (p P) IsType()  {}
func (p P) IsValue() {}

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
func (p P) Iterate(sco interfaces.Scope) (interfaces.Iterable, error) {
	if p.HasTail() {
		return p.tail, nil
	}
	return ENDED, nil
}

// ToSlice returns an array from a P
func (p P) ToSlice(sco interfaces.Scope) ([]interfaces.Type, error) {
	slice := []interfaces.Type{}
	var tail interfaces.Iterable = p
	for {
		if tail != ENDED {
			slice = append(slice, tail.Head())
			if !tail.HasTail() {
				return slice, nil
			}
			var err error
			tail, err = tail.Iterate(sco)
			if err != nil {
				return slice, err
			}
		}
	}
}

// LAZYP (Lazily evaluated Pair)
// head operates like Pair, tail should be an expression that returns another LAZYP
type LAZYP struct {
	head interfaces.Value
	tail interfaces.Evaluatable
}

// IsType for LAZYP
func (l LAZYP) IsType()  {}
func (l LAZYP) IsValue() {}

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
func (l LAZYP) Iterate(sco interfaces.Scope) (interfaces.Iterable, error) {
	taileval, err := l.tail.Evaluate(sco)
	if err != nil {
		return ENDED, err
	}
	if nextIter, ok := taileval.(interfaces.Iterable); ok {
		return nextIter, nil
	}
	return ENDED, fmt.Errorf("lazypair : iterable expected got %v", taileval)
}

// ToSlice converts a LAZYP to a slice by iterating through it
func (l LAZYP) ToSlice(sco interfaces.Scope) ([]interfaces.Type, error) {
	slice := []interfaces.Type{}
	var next interfaces.Iterable = l
	for {
		slice = append(slice, next.Head())

		if !next.HasTail() {
			return slice, nil
		}
		res, err := next.Iterate(sco.NewChildScope())
		if err != nil {
			return slice, err
		}
		next = res

	}
}

func createLAZYP(sco interfaces.Scope, head interfaces.Value, function interfaces.Type, args ...interfaces.Type) LAZYP {
	return LAZYP{
		head,
		BindEvaluation(&EXP{function, args}, sco)}
}

// END acts as the end of a list
type END struct{}

func (e END) IsType()  {}
func (e END) IsValue() {}

func (e END) String() string {
	return "<END>"
}
func (e END) Head() interfaces.Value {
	return NILL
}
func (e END) HasTail() bool {
	return false
}
func (e END) Iterate(interfaces.Scope) (interfaces.Iterable, error) {
	panic("END : Iterate called on END")
}
func (e END) ToSlice(interfaces.Scope) ([]interfaces.Type, error) {
	return []interfaces.Type{}, nil
}

var ENDED = END{}
