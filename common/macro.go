package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

// MAC expands the Expression with Arguments prior to evaluation
type MAC struct {
	Arguments  VEC
	Expression *EXP
}

// IsType for MAC
func (m MAC) IsType() {}

func (m MAC) IsValue() {}

// String representation of MAC
func (m MAC) String() string {
	return fmt.Sprintf("MAC(%v %v)", m.Arguments, m.Expression)
}

//Expand will replace references to Arguments with arguments provided and then return it without evaluation
func (m MAC) Expand(arguments []interfaces.Type) interfaces.Evaluatable {
	m.printStartExpand()
	if len(arguments) != len(m.Arguments.Vector) {
		panic(fmt.Sprintf("Expand : expected %v arguments\n", len(m.Arguments.Vector)))
	}

	variables := make(map[REF]interfaces.Type)

	for p, ref := range m.Arguments.Vector {
		variables[ref.(REF)] = arguments[p]
	}

	var expandEXP func(*EXP) *EXP
	expandEXP = func(exp *EXP) *EXP {
		newArgs := make([]interfaces.Type, len(exp.Arguments))
		for p, arg := range exp.Arguments {
			if ref, ok := arg.(REF); ok {
				if newVal, found := variables[ref]; found {
					newArgs[p] = newVal
				} else {
					newArgs[p] = ref
				}
			} else if exp, ok := arg.(*EXP); ok {
				newArgs[p] = expandEXP(exp)
			} else {
				newArgs[p] = arg
			}
		}
		return &EXP{exp.Function, newArgs}
	}
	result := expandEXP(m.Expression)
	m.printEndExpand(result)
	return result
}

func (m MAC) printStartExpand() {
	if DEBUG {
		fmt.Printf("Expanding %v\n", m)
	}
}

func (m MAC) printEndExpand(expanded interfaces.Type) {
	if DEBUG {
		fmt.Printf("Expanded to %v\n", expanded)
	}
}
