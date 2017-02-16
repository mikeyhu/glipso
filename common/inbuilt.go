package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

func plusAll(arguments []interfaces.Argument, sco interfaces.Scope) interfaces.Argument {
	all := I(0)
	for _, v := range arguments {
		all += v.(I)
	}
	return all
}

func minusAll(arguments []interfaces.Argument, sco interfaces.Scope) interfaces.Argument {
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

func equals(arguments []interfaces.Argument, sco interfaces.Scope) interfaces.Argument {
	switch t := arguments[0].(type) {
	case B:
		return t.Equals(arguments[1])
	case I:
		return t.Equals(arguments[1])
	default:
		panic("Equals : unsupported type")
	}
}

func cons(arguments []interfaces.Argument, sco interfaces.Scope) interfaces.Argument {
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

func first(arguments []interfaces.Argument, sco interfaces.Scope) interfaces.Argument {
	pair, ok := arguments[0].(P)
	if ok {
		return pair.head
	}
	panic("Panic - Cannot get head of non Pair type")
}

func tail(arguments []interfaces.Argument, sco interfaces.Scope) interfaces.Argument {
	pair, ok := arguments[0].(P)
	if ok {
		return *pair.tail
	}
	panic("Panic - Cannot get tail of non Pair type")
}

func apply(arguments []interfaces.Argument, sco interfaces.Scope) interfaces.Argument {
	if ap, okEval := arguments[1].(interfaces.Evaluatable); okEval {
		arguments[1] = ap.Evaluate(sco)
	}
	s, okRef := arguments[0].(REF)
	p, okPair := arguments[1].(interfaces.Iterable)
	if okRef && okPair {
		return EXP{FunctionName: s.String(), Arguments: p.ToSlice(sco.NewChildScope())}
	}
	panic(fmt.Sprintf("Panic - expected function, found %v", arguments[0]))
}

func iff(arguments []interfaces.Argument, sco interfaces.Scope) interfaces.Argument {
	var test interfaces.Argument
	if exp, ok := arguments[0].(EXP); ok {
		test = exp.Evaluate(sco)
	} else {
		test = arguments[0]
	}
	if test.(B).Bool() {
		return arguments[1]
	}
	return arguments[2]
}

func def(arguments []interfaces.Argument, sco interfaces.Scope) interfaces.Argument {
	if eval, ok := arguments[1].(interfaces.Evaluatable); ok {
		arguments[1] = eval.Evaluate(sco.NewChildScope())
	}
	return GlobalEnvironment.CreateRef(arguments[0].(REF), arguments[1])
}

func do(arguments []interfaces.Argument, sco interfaces.Scope) interfaces.Argument {
	var result interfaces.Argument
	for _, a := range arguments {
		if e, ok := a.(interfaces.Evaluatable); ok {
			result = e.Evaluate(sco)
		} else {
			result = a
		}
	}
	return result
}

func rnge(arguments []interfaces.Argument, sco interfaces.Scope) interfaces.Argument {
	start := arguments[0].(I)
	end := arguments[1].(I)
	if start < end {
		return LAZYP{
			start,
			&EXP{FunctionName: "range", Arguments: []interfaces.Argument{
				I(start.Int() + 1),
				end,
			}}}
	}
	return LAZYP{end, nil}

}

func fn(arguments []interfaces.Argument, sco interfaces.Scope) interfaces.Argument {
	return FN{arguments[0].(VEC), arguments[1].(EXP)}
}

type evaluator func([]interfaces.Argument, interfaces.Scope) interfaces.Argument

type FunctionInfo struct {
	function     evaluator
	evaluateArgs bool
}

var inbuilt map[string]FunctionInfo

func init() {
	inbuilt = map[string]FunctionInfo{
		"cons":  {cons, true},
		"first": {first, true},
		"tail":  {tail, true},
		"=":     {equals, true},
		"+":     {plusAll, true},
		"-":     {minusAll, true},
		"apply": {apply, false},
		"if":    {iff, false},
		"def":   {def, false},
		"do":    {do, false},
		"range": {rnge, true},
		"fn":    {fn, false},
	}
}
