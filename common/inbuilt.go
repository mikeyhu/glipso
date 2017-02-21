package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

func plusAll(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	all := I(0)
	for _, v := range arguments {
		all += v.(I)
	}
	return all
}

func minusAll(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
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

func multiplyAll(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	var all I
	head := true
	for _, v := range arguments {
		if head {
			all = v.(I)
			head = false
		} else {
			all *= v.(I)
		}
	}
	return all
}

func mod(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	a, aok := arguments[0].(I)
	b, bok := arguments[1].(I)
	if aok && bok {
		return I(a % b)
	}
	panic("Mod : unsupported type")
}

func equals(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	switch t := arguments[0].(type) {
	case B:
		return t.Equals(arguments[1])
	case I:
		return t.Equals(arguments[1])
	default:
		panic("Equals : unsupported type")
	}
}

func cons(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
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

func first(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	pair, ok := arguments[0].(P)
	if ok {
		return pair.head
	}
	panic("Panic - Cannot get head of non Pair type")
}

func tail(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	pair, ok := arguments[0].(P)
	if ok {
		return *pair.tail
	}
	panic("Panic - Cannot get tail of non Pair type")
}

func apply(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	if ap, okEval := arguments[1].(interfaces.Evaluatable); okEval {
		arguments[1] = ap.Evaluate(sco)
	}
	s, okRef := arguments[0].(REF)
	p, okPair := arguments[1].(interfaces.Iterable)
	if okRef && okPair {
		return EXP{Function: s, Arguments: p.ToSlice(sco.NewChildScope())}
	}
	panic(fmt.Sprintf("Panic - expected function, found %v", arguments[0]))
}

func iff(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	var test interfaces.Type
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

func def(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	var value interfaces.Type
	if eval, ok := arguments[1].(interfaces.Evaluatable); ok {
		value = eval.Evaluate(sco.NewChildScope())
	} else {
		value = arguments[1]
	}
	return GlobalEnvironment.CreateRef(arguments[0].(REF).EvaluateToRef(sco.NewChildScope()), value)
}

func do(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	var result interfaces.Type
	for _, a := range arguments {
		if e, ok := a.(interfaces.Evaluatable); ok {
			result = e.Evaluate(sco.NewChildScope())
		} else {
			result = a
		}
	}
	return result
}

func rnge(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	start := arguments[0].(I)
	end := arguments[1].(I)
	if start < end {
		return LAZYP{
			start,
			&EXP{Function: REF("range"), Arguments: []interfaces.Type{
				I(start.Int() + 1),
				end,
			}}}
	}
	return LAZYP{end, nil}

}

func fn(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	var argVec VEC
	if args, ok := arguments[0].(REF); ok {
		argVec = args.Evaluate(sco).(VEC)
	} else {
		argVec = arguments[0].(VEC)
	}

	if arg1, ok := arguments[1].(REF); ok {
		return FN{argVec, arg1.Evaluate(sco.NewChildScope()).(EXP)}
	} else {

	}
	return FN{argVec, arguments[1].(EXP)}
}

type evaluator func([]interfaces.Type, interfaces.Scope) interfaces.Type

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
		"*":     {multiplyAll, true},
		"%":     {mod, true},
		"apply": {apply, false},
		"if":    {iff, false},
		"def":   {def, false},
		"do":    {do, false},
		"range": {rnge, true},
		"fn":    {fn, false},
	}
}
