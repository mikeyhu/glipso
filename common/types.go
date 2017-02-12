package common

import (
	"fmt"
	"github.com/mikeyhu/mekkanism/interfaces"
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
func (p P) ToSlice() []interfaces.Argument {
	slice := []interfaces.Argument{p.head}
	tail := p.tail
	for {
		if tail == nil {
			return slice
		} else {
			slice = append(slice, tail.head)
			tail = tail.tail
		}
	}
	return slice
}

type REF string //symbol for something in scope, variable or function

func (r REF) IsArg() {}
func (r REF) String() string {
	return fmt.Sprintf("%v", string(r))
}
func (r REF) Evaluate() interfaces.Argument {
	return GlobalEnvironment.resolveRef(r)
}
