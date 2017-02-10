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

type B bool

func (b B) IsArg() {}
func (b B) String() string {
	return fmt.Sprintf("%t", b.Bool())
}
func (b B) Bool() bool {
	return bool(b)
}
func (b B) Equals(o B) B {
	return B(b.Bool() == o.Bool())
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
	for i:=1;i<10;i++ {
		if tail == nil {
			return slice
		} else {
			slice = append(slice, tail.head)
			tail = tail.tail
		}
	}
	return slice
}

type SCOPE string //symbol for something in scope, variable or function

func (s SCOPE) IsArg() {}
func (s SCOPE) String() string {
	return fmt.Sprintf("%v", string(s))
}
