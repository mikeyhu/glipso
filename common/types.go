package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

type I int

func (i I) IsArg() {}
func (i I) String() string {
	return fmt.Sprintf("%d", i.Int())
}
func (i I) Int() int {
	return int(i)
}
func (i I) Equals(o interfaces.Argument) B {
	if other, ok := o.(I); ok {
		return B(i.Int() == other.Int())
	}
	return B(false)
}

type B bool

func (b B) IsArg() {}
func (b B) String() string {
	return fmt.Sprintf("%t", b.Bool())
}
func (b B) Bool() bool {
	return bool(b)
}
func (b B) Equals(o interfaces.Argument) B {
	if other, ok := o.(B); ok {
		return B(b.Bool() == other.Bool())
	}
	return B(false)
}

type P struct {
	head interfaces.Argument
	tail *P
}

func (p P) IsArg() {}
func (p P) String() string {
	return fmt.Sprintf("￿`(%v %v)", p.head, p.tail)
}
func (p P) Iterate(sco interfaces.Scope) interfaces.Iterable {
	panic("Iterate called on PAIR")
}
func (p P) ToSlice(sco interfaces.Scope) []interfaces.Argument {
	slice := []interfaces.Argument{p.head}
	tail := p.tail
	for {
		if tail == nil {
			return slice
		}
		slice = append(slice, tail.head)
		tail = tail.tail
	}
}

type REF string //symbol for something in scope, variable or function

func (r REF) IsArg() {}
func (r REF) String() string {
	return fmt.Sprintf("%v", string(r))
}
func (r REF) Evaluate(sco interfaces.Scope) interfaces.Argument {
	return sco.ResolveRef(r)
}

type LAZYP struct {
	head interfaces.Argument
	tail *EXP
}

func (l LAZYP) IsArg() {}
func (l LAZYP) String() string {
	return fmt.Sprintf("￿`(%v %v)", l.head, l.tail)
}
func (l LAZYP) Iterate(sco interfaces.Scope) interfaces.Iterable {
	if nextIter, ok := l.tail.Evaluate(sco).(LAZYP); ok {
		return nextIter
	}
	panic(fmt.Sprintf("Iterate : expected an LAZYP, got %v", l))
}
func (l LAZYP) ToSlice(sco interfaces.Scope) []interfaces.Argument {
	slice := []interfaces.Argument{}
	next := l
	for {
		slice = append(slice, next.head)
		if next.tail == nil {
			return slice
		}
		next = next.Iterate(sco.NewChildScope()).(LAZYP)
	}
}

type VEC struct {
	Vector []interfaces.Argument
}
func (v VEC) IsArg() {}
func (v VEC) String() string {
	return fmt.Sprintf("%v", v.Vector)
}
func (v VEC) Get(loc int) interfaces.Argument {
	return v.Vector[loc]
}

type FN struct {
	Arguments VEC
	Expression EXP
}
func (f FN) IsArg() {}

