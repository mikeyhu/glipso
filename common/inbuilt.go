package common

import (
	"fmt"
	"github.com/mikeyhu/mekkanism/interfaces"
)

func PlusAll(arguments []interfaces.Argument) interfaces.Argument {
	all := I(0)
	for _, v := range arguments {
		all += v.(I)
	}
	return all
}

func MinusAll(arguments []interfaces.Argument) interfaces.Argument {
	var all I
	head := true
	for _, v := range arguments {
		if head {
			all = v.(I)
			head = false
		} else {
			all -= v.(I)
		}
	}
	return all
}

func Equals(arguments []interfaces.Argument) interfaces.Argument {
	switch t := arguments[0].(type) {
	case B:
		return t.Equals(arguments[1])
	case I:
		return t.Equals(arguments[1])
	default:
		panic("Equals : unsupported type")
	}
}

func allArgumentsAreB(arguments []interfaces.Argument) bool {
	for _, a := range arguments {
		if !a.(B) {
			return false
		}
	}
	return true
}

func Cons(arguments []interfaces.Argument) interfaces.Argument {
	if len(arguments) == 0 {
		return P{}
	} else if len(arguments) == 1 {
		return P{arguments[0], nil}
	} else if len(arguments) == 2 {
		tail, ok := arguments[1].(P)
		if ok {
			return P{arguments[0], &tail}
		}
	}
	return P{}
}

func First(arguments []interfaces.Argument) interfaces.Argument {
	pair, ok := arguments[0].(P)
	if ok {
		return pair.head
	} else {
		panic("Panic - Cannot get head of non Pair type")
	}
}

func Tail(arguments []interfaces.Argument) interfaces.Argument {
	pair, ok := arguments[0].(P)
	if ok {
		return *pair.tail
	} else {
		panic("Panic - Cannot get tail of non Pair type")
	}
}

func Apply(arguments []interfaces.Argument) interfaces.Argument {
	s, okScope := arguments[0].(SCOPE)
	p, okPair := arguments[1].(P)
	if okScope && okPair {
		return Expression{FunctionName: s.String(), Arguments: p.ToSlice()}
	} else {
		panic(fmt.Sprintf("Panic - expected function, found %v", arguments[0]))
	}
}

func If(arguments []interfaces.Argument) interfaces.Argument {
	var test interfaces.Argument
	fmt.Printf("If Args: %v\n", arguments)
	fmt.Printf("If 1st arg %v", arguments[0])
	if exp, ok := arguments[0].(Expression); ok {
		fmt.Printf("If evaluating\n")
		test = exp.Evaluate()
	} else {
		fmt.Printf("If skipping evaluation\n")
		test = arguments[0]
	}
	fmt.Printf("If test %v\n", test)
	if test.(B).Bool() {
		return arguments[1]
	} else {
		return arguments[2]
	}
}

type evaluations func([]interfaces.Argument) interfaces.Argument

type FunctionInfo struct {
	function     evaluations
	evaluateArgs bool
}

var inbuilt map[string]FunctionInfo

func init() {
	inbuilt = map[string]FunctionInfo{
		"cons":  {Cons, true},
		"first": {First, true},
		"tail":  {Tail, true},
		"=":     {Equals, true},
		"+":     {PlusAll, true},
		"-":     {MinusAll, true},
		"apply": {Apply, true},
		"if":    {If, false},
	}
}
