package common

import (
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
	return arguments[0].(B).Equals(arguments[1].(B))
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

type F func([]interfaces.Argument) interfaces.Argument

var inbuilt = map[string]F{
	"cons":  Cons,
	"first": First,
	"tail":  Tail,
	"=":     Equals,
	"+":     PlusAll,
	"-":     MinusAll,
}
