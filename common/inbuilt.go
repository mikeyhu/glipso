package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

type evaluator func([]interfaces.Type, interfaces.Scope) interfaces.Type

// FunctionInfo provides information about a built in function
type FunctionInfo struct {
	name         string
	function     evaluator
	evaluateArgs bool
}

// IsType for FunctionInfo
func (fi FunctionInfo) IsType() {}

// String for FunctionInfo
func (fi FunctionInfo) String() string {
	return fmt.Sprintf("FI(%s)", fi.name)
}
func (fi FunctionInfo) Apply(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	evaluatedArgs := make([]interfaces.Type, len(arguments))
	if fi.evaluateArgs {
		for p, arg := range arguments {
			if r, ok := arg.(REF); ok {
				arg = r.Evaluate(sco)
			}
			if e, ok := arg.(interfaces.Evaluatable); ok {
				arg = e.Evaluate(sco)
			}
			evaluatedArgs[p] = arg
		}
	} else {
		copy(evaluatedArgs, arguments)
	}
	return fi.function(evaluatedArgs, sco)
}

var inbuilt map[REF]FunctionInfo

func init() {
	inbuilt = map[REF]FunctionInfo{
		REF("="):        {"=", equals, true},
		REF("+"):        {"+", plusAll, true},
		REF("-"):        {"-", minusAll, true},
		REF("*"):        {"*", multiplyAll, true},
		REF("%"):        {"%", mod, true},
		REF("<"):        {"<", lessThan, true},
		REF(">"):        {">", greaterThan, true},
		REF("<="):       {"<=", lessThanEqual, true},
		REF(">="):       {">=", greaterThanEqual, true},
		REF("apply"):    {"apply", apply, false},
		REF("cons"):     {"cons", cons, true},
		REF("def"):      {"def", def, false},
		REF("do"):       {"do", do, false},
		REF("empty"):    {"empty", empty, true},
		REF("if"):       {"if", iff, false},
		REF("filter"):   {"filter", filter, true},
		REF("first"):    {"first", first, true},
		REF("fn"):       {"fn", fn, false},
		REF("lazypair"): {"lazypair", lazypair, false},
		REF("macro"):    {"macro", macro, false},
		REF("map"):      {"map", mapp, false},
		REF("print"):    {"print", printt, true},
		REF("range"):    {"range", rnge, true},
		REF("tail"):     {"tail", tail, true},
		REF("take"):     {"take", take, true},
	}
}

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
	first, fok := arguments[0].(interfaces.Equalable)
	second, sok := arguments[1].(interfaces.Equalable)

	if fok && sok {
		return first.Equals(second)
	}
	panic(fmt.Sprintf("Equals : unsupported type %v  or %v\n", arguments[0], arguments[1]))
}

func lessThan(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) < 0)
	}
	panic(fmt.Sprintf("LessThan : unsupported type %v  or %v\n", arguments[0], arguments[1]))
}

func lessThanEqual(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) <= 0)
	}
	panic(fmt.Sprintf("LessThan : unsupported type %v  or %v\n", arguments[0], arguments[1]))
}

func greaterThan(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) > 0)
	}
	panic(fmt.Sprintf("GreaterThan : unsupported type %v  or %v\n", arguments[0], arguments[1]))
}

func greaterThanEqual(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) >= 0)
	}
	panic(fmt.Sprintf("GreaterThan : unsupported type %v  or %v\n", arguments[0], arguments[1]))
}

func cons(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	if len(arguments) == 0 {
		return P{}
	} else if len(arguments) == 1 {
		return P{arguments[0], ENDED}
	} else if len(arguments) == 2 {
		tail, ok := arguments[1].(P)
		if ok {
			return P{arguments[0], tail}
		}
	}
	return P{}
}

func first(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	pair, ok := arguments[0].(interfaces.Iterable)
	if ok {
		return pair.Head()
	}
	fmt.Printf("pair? %v : %v\n", arguments[0], pair)
	panic("Panic - Cannot get head of non Iterable type")
}

func tail(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	pair, ok := arguments[0].(interfaces.Iterable)
	if ok {
		if pair.HasTail() {
			return pair.Iterate(sco)
		}
		return ENDED
	}
	panic("Panic - Cannot get tail of non Pair type")
}

func apply(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	if ap, okEval := arguments[1].(interfaces.Evaluatable); okEval {
		arguments[1] = ap.Evaluate(sco)
	}
	s, okRef := arguments[0].(REF)
	p, okPair := arguments[1].(interfaces.Iterable)
	if !okRef {
		panic(fmt.Sprintf("Panic - expected function, found %v", arguments[0]))
	} else if !okPair {
		panic(fmt.Sprintf("Panic - expected pair, found %v", arguments[1]))
	}
	return &EXP{Function: s, Arguments: p.ToSlice(sco.NewChildScope())}
}

func iff(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	var test interfaces.Type
	if exp, ok := arguments[0].(*EXP); ok {
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
	return P{end, ENDED}

}

func fn(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	var argVec VEC
	if args, ok := arguments[0].(REF); ok {
		argVec = args.Evaluate(sco).(VEC)
	} else {
		argVec = arguments[0].(VEC)
	}

	if arg1, ok := arguments[1].(REF); ok {
		return FN{argVec, arg1.Evaluate(sco.NewChildScope()).(*EXP)}
	}
	return FN{argVec, arguments[1].(*EXP)}
}

func filter(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	fn, fnok := arguments[0].(FN)
	var list interfaces.Type = arguments[1]

	if exp, eok := list.(interfaces.Evaluatable); eok {
		list = exp.Evaluate(sco)
	}

	iter, iok := list.(interfaces.Iterable)

	var flt func(interfaces.Iterable) interfaces.Iterable
	flt = func(it interfaces.Iterable) interfaces.Iterable {
		head := it.Head()
		res := (&EXP{Function: fn, Arguments: []interfaces.Type{head}}).Evaluate(sco.NewChildScope())
		if include, iok := res.(B); iok {
			if bool(include) {
				if it.HasTail() {
					return &P{head, flt(it.Iterate(sco))}
				}
				return &P{head, ENDED}
			} else if it.HasTail() {
				return flt(it.Iterate(sco))
			}
			return ENDED
		}
		panic(fmt.Sprintf("filter : expected boolean value, recieved %v", res))
	}

	if fnok && iok {
		return flt(iter)
	}
	panic(fmt.Sprintf("filter : expected function and list. Recieved %v, %v\n", arguments[0], arguments[1]))
}

func mapp(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {

	var mp func(interfaces.Type, interfaces.Iterable) *P
	mp = func(fn interfaces.Type, iterable interfaces.Iterable) *P {
		head := iterable.Head()
		res := (&EXP{Function: fn, Arguments: []interfaces.Type{head}}).Evaluate(sco.NewChildScope())
		if iterable.HasTail() {
			return &P{res, mp(fn, iterable.Iterate(sco))}
		}
		return &P{res, ENDED}
	}

	list := arguments[1]
	if eval, ok := list.(interfaces.Evaluatable); ok {
		list = eval.Evaluate(sco)
	}

	if pair, pok := list.(interfaces.Iterable); pok {
		return *mp(arguments[0], pair)
	}

	panic(fmt.Sprintf("map : expected function and list, recieved %v %v", arguments[0], arguments[1]))
}

func lazypair(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	head := arguments[0]
	if h, ok := head.(interfaces.Evaluatable); ok {
		head = h.Evaluate(sco)
	}
	if len(arguments) > 1 {
		if tail, ok := arguments[1].(interfaces.Evaluatable); ok {
			return LAZYP{head, BindEvaluation(tail, sco)}
		}
		panic(fmt.Sprintf("lazypair : expected EXP got %v", arguments[1]))
	}
	return LAZYP{head, nil}
}

func macro(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	return MAC{arguments[0].(VEC), arguments[1].(*EXP)}
}

func printt(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	for _, arg := range arguments {
		fmt.Printf("%v\n", arg)
	}
	return NILL
}

func empty(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	var arg interfaces.Type = arguments[0]
	if arg == nil {
		return B(true)
	}
	list := arg.(interfaces.Iterable)
	if list != ENDED {
		return B(false)
	}
	return B(true)
}

func take(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Type {
	num, nok := arguments[0].(I)
	list, lok := arguments[1].(interfaces.Iterable)

	if nok && lok {
		if num > 1 && list.HasTail() {
			return LAZYP{
				list.Head(),
				&EXP{
					Function: REF("take"),
					Arguments: []interfaces.Type{
						I(num - 1),
						list.Iterate(sco),
					},
				},
			}
		}
		return P{list.Head(), ENDED}

	}
	panic("take : expected number and list")
}
