package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

type evaluator func([]interfaces.Value, interfaces.Scope) interfaces.Value
type lazyEvaluator func([]interfaces.Type, interfaces.Scope) interfaces.Value

// FunctionInfo provides information about a built in function
type FunctionInfo struct {
	name          string
	evaluator     evaluator
	lazyEvaluator lazyEvaluator
}

// IsType for FunctionInfo
func (fi FunctionInfo) IsType()   {}
func (fi FunctionInfo) IsResult() {}

// String for FunctionInfo
func (fi FunctionInfo) String() string {
	return fmt.Sprintf("FI(%s)", fi.name)
}
func (fi FunctionInfo) Apply(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Value {
	if fi.evaluator != nil {
		evaluatedArgs := make([]interfaces.Value, len(arguments))
		for p, arg := range arguments {
			evaluatedArgs[p] = evaluateToResult(arg, sco)
		}
		return fi.evaluator(evaluatedArgs, sco)
	} else if fi.lazyEvaluator != nil {
		unevaluatedArgs := make([]interfaces.Type, len(arguments))
		copy(unevaluatedArgs, arguments)
		return fi.lazyEvaluator(unevaluatedArgs, sco)
	}
	panic(fmt.Sprintf("FunctionInfo : $v had neither an evaluator or lazy evaluator", fi.name))
}

var inbuilt map[REF]FunctionInfo

func init() {
	inbuilt = map[REF]FunctionInfo{
		REF("="):        {name: "=", evaluator: equals},
		REF("+"):        {name: "+", evaluator: plusAll},
		REF("-"):        {name: "-", evaluator: minusAll},
		REF("*"):        {name: "*", evaluator: multiplyAll},
		REF("%"):        {name: "%", evaluator: mod},
		REF("<"):        {name: "<", evaluator: lessThan},
		REF(">"):        {name: ">", evaluator: greaterThan},
		REF("<="):       {name: "<=", evaluator: lessThanEqual},
		REF(">="):       {name: ">=", evaluator: greaterThanEqual},
		REF("apply"):    {name: "apply", lazyEvaluator: apply},
		REF("cons"):     {name: "cons", evaluator: cons},
		REF("def"):      {name: "def", lazyEvaluator: def},
		REF("do"):       {name: "do", lazyEvaluator: do},
		REF("empty"):    {name: "empty", evaluator: empty},
		REF("if"):       {name: "if", lazyEvaluator: iff},
		REF("filter"):   {name: "filter", evaluator: filter},
		REF("first"):    {name: "first", evaluator: first},
		REF("fn"):       {name: "fn", lazyEvaluator: fn},
		REF("lazypair"): {name: "lazypair", lazyEvaluator: lazypair},
		REF("let"):      {name: "let", lazyEvaluator: let},
		REF("macro"):    {name: "macro", lazyEvaluator: macro},
		REF("map"):      {name: "map", lazyEvaluator: mapp},
		REF("print"):    {name: "print", evaluator: printt},
		REF("range"):    {name: "range", evaluator: rnge},
		REF("tail"):     {name: "tail", evaluator: tail},
		REF("take"):     {name: "take", evaluator: take},
	}
}

func evaluateToResult(value interfaces.Type, sco interfaces.Scope) interfaces.Value {
	switch v := value.(type) {
	case interfaces.Evaluatable:
		return v.Evaluate(sco)
	case interfaces.Value:
		return v
	default:
		panic(fmt.Sprintf("evaluateToResult : value %v of type %v is neither evaluatable or a result", value, v))
	}
}

func plusAll(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
	all := I(0)
	for _, v := range arguments {
		all += v.(I)
	}
	return all
}

func minusAll(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
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

func multiplyAll(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
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

func mod(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
	a, aok := arguments[0].(I)
	b, bok := arguments[1].(I)
	if aok && bok {
		return I(a % b)
	}
	panic("Mod : unsupported type")
}

func equals(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
	first, fok := arguments[0].(interfaces.Equalable)
	second, sok := arguments[1].(interfaces.Equalable)

	if fok && sok {
		return first.Equals(second)
	}
	panic(fmt.Sprintf("Equals : unsupported type %v  or %v\n", arguments[0], arguments[1]))
}

func lessThan(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) < 0)
	}
	panic(fmt.Sprintf("LessThan : unsupported type %v  or %v\n", arguments[0], arguments[1]))
}

func lessThanEqual(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) <= 0)
	}
	panic(fmt.Sprintf("LessThan : unsupported type %v  or %v\n", arguments[0], arguments[1]))
}

func greaterThan(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) > 0)
	}
	panic(fmt.Sprintf("GreaterThan : unsupported type %v  or %v\n", arguments[0], arguments[1]))
}

func greaterThanEqual(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) >= 0)
	}
	panic(fmt.Sprintf("GreaterThan : unsupported type %v  or %v\n", arguments[0], arguments[1]))
}

func cons(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
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

func first(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
	pair, ok := arguments[0].(interfaces.Iterable)
	if ok {
		return pair.Head()
	}
	fmt.Printf("pair? %v : %v\n", arguments[0], pair)
	panic("Panic - Cannot get head of non Iterable type")
}

func tail(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
	pair, ok := arguments[0].(interfaces.Iterable)
	if ok {
		if pair.HasTail() {
			return pair.Iterate(sco)
		}
		return ENDED
	}
	panic("Panic - Cannot get tail of non Pair type")
}

func apply(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Value {
	list := evaluateToResult(arguments[1], sco)
	s, okRef := arguments[0].(REF)
	p, okPair := list.(interfaces.Iterable)
	if !okRef {
		panic(fmt.Sprintf("Panic - expected function, found %v", arguments[0]))
	} else if !okPair {
		panic(fmt.Sprintf("Panic - expected pair, found %v", list))
	}
	return evaluateToResult(&EXP{Function: s, Arguments: p.ToSlice(sco.NewChildScope())}, sco)
}

func iff(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Value {
	test := evaluateToResult(arguments[0], sco)
	if test.(B).Bool() {
		return evaluateToResult(arguments[1], sco)
	}
	return evaluateToResult(arguments[2], sco)
}

func def(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Value {
	value := evaluateToResult(arguments[1], sco)
	GlobalEnvironment.CreateRef(arguments[0].(REF), value)
	return NILL
}

func do(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Value {
	var result interfaces.Value
	for _, a := range arguments {
		result = evaluateToResult(a, sco.NewChildScope())
	}
	return result
}

func rnge(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
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

func fn(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Value {
	var argVec VEC
	if args, ok := arguments[0].(REF); ok {
		argVec = args.Evaluate(sco).(VEC)
	} else {
		argVec = arguments[0].(VEC)
	}

	return FN{argVec, arguments[1].(interfaces.Evaluatable)}
}

func filter(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
	fn, fnok := arguments[0].(FN)
	list := evaluateToResult(arguments[1], sco)

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

func mapp(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Value {

	var mp func(interfaces.Type, interfaces.Iterable) *P
	mp = func(fn interfaces.Type, iterable interfaces.Iterable) *P {
		head := iterable.Head()
		res := (&EXP{Function: fn, Arguments: []interfaces.Type{head}}).Evaluate(sco.NewChildScope())
		if iterable.HasTail() {
			return &P{res, mp(fn, iterable.Iterate(sco))}
		}
		return &P{res, ENDED}
	}

	list := evaluateToResult(arguments[1], sco)

	if pair, pok := list.(interfaces.Iterable); pok {
		return *mp(arguments[0], pair)
	}

	panic(fmt.Sprintf("map : expected function and list, recieved %v %v", arguments[0], arguments[1]))
}

func lazypair(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Value {
	head := evaluateToResult(arguments[0], sco)
	if len(arguments) > 1 {
		if tail, ok := arguments[1].(interfaces.Evaluatable); ok {
			return LAZYP{head, BindEvaluation(tail, sco)}
		}
		panic(fmt.Sprintf("lazypair : expected EXP got %v", arguments[1]))
	}
	return LAZYP{head, nil}
}

func macro(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Value {
	return MAC{arguments[0].(VEC), arguments[1].(*EXP)}
}

func printt(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
	for _, arg := range arguments {
		fmt.Printf("%v\n", arg)
	}
	return NILL
}

func empty(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
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

func take(arguments []interfaces.Value, sco interfaces.Scope) interfaces.Value {
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

func let(arguments []interfaces.Type, sco interfaces.Scope) interfaces.Value {
	vectors, vok := arguments[0].(VEC)
	exp, eok := arguments[1].(interfaces.Evaluatable)

	childScope := sco.NewChildScope()

	if vok && eok {
		count := vectors.Count()
		if count%2 > 0 {
			panic(fmt.Sprintf("let : expected an even number of items in vector, recieved %v", count))
		}
		for i := 0; i < count/2; i++ {
			childScope.CreateRef(vectors.Get(i), evaluateToResult(vectors.Get(i+1), sco))
		}
		return exp.Evaluate(childScope)
	}
	panic(fmt.Sprintf("let : expected VEC and EXP, received: %v %v", arguments[0], arguments[1]))
}
