package common

import (
	"errors"
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

type evaluator func([]interfaces.Value, interfaces.Scope) (interfaces.Value, error)
type lazyEvaluator func([]interfaces.Type, interfaces.Scope) (interfaces.Value, error)

var inbuilt map[REF]FI

func init() {
	inbuilt = map[REF]FI{}
	addInbuilt(FI{name: "=", evaluator: equals})
	addInbuilt(FI{name: "+", evaluator: plusAll})
	addInbuilt(FI{name: "-", evaluator: minusAll})
	addInbuilt(FI{name: "*", evaluator: multiplyAll})
	addInbuilt(FI{name: "%", evaluator: mod})
	addInbuilt(FI{name: "<", evaluator: lessThan})
	addInbuilt(FI{name: ">", evaluator: greaterThan})
	addInbuilt(FI{name: "<=", evaluator: lessThanEqual})
	addInbuilt(FI{name: ">=", evaluator: greaterThanEqual})
	addInbuilt(FI{name: "apply", lazyEvaluator: apply})
	addInbuilt(FI{name: "cons", evaluator: cons})
	addInbuilt(FI{name: "def", lazyEvaluator: def})
	addInbuilt(FI{name: "do", lazyEvaluator: do})
	addInbuilt(FI{name: "empty", evaluator: empty})
	addInbuilt(FI{name: "if", lazyEvaluator: iff})
	addInbuilt(FI{name: "filter", evaluator: filter})
	addInbuilt(FI{name: "first", evaluator: first})
	addInbuilt(FI{name: "fn", lazyEvaluator: fn})
	addInbuilt(FI{name: "lazypair", lazyEvaluator: lazypair})
	addInbuilt(FI{name: "let", lazyEvaluator: let})
	addInbuilt(FI{name: "macro", lazyEvaluator: macro})
	addInbuilt(FI{name: "map", evaluator: mapp})
	addInbuilt(FI{name: "print", evaluator: printt})
	addInbuilt(FI{name: "range", evaluator: rnge})
	addInbuilt(FI{name: "tail", evaluator: tail})
	addInbuilt(FI{name: "take", evaluator: take})
}

func addInbuilt(info FI) {
	inbuilt[REF(info.name)] = info
}

func plusAll(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	all := I(0)
	for _, v := range arguments {
		all += v.(I)
	}
	return all, nil
}

func minusAll(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
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
	return all, nil
}

func multiplyAll(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
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
	return all, nil
}

func mod(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	if len(arguments) != 2 {
		return NILL, fmt.Errorf("mod : expected 2 arguments, recieved %d", len(arguments))
	}
	a, aok := arguments[0].(I)
	b, bok := arguments[1].(I)
	if aok && bok {
		return I(a % b), nil
	}
	return NILL, errors.New("mod : unsupported type")
}

func equals(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	first, fok := arguments[0].(interfaces.Equalable)
	second, sok := arguments[1].(interfaces.Equalable)

	if fok && sok {
		return first.Equals(second), nil
	}
	return NILL, fmt.Errorf("Equals : unsupported type %v or %v", arguments[0], arguments[1])
}

func lessThan(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) < 0), nil
	}
	return NILL, fmt.Errorf("lessThan : unsupported type %v or %v", arguments[0], arguments[1])
}

func lessThanEqual(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) <= 0), nil
	}
	return NILL, fmt.Errorf("lessThanEqual : unsupported type %v or %v", arguments[0], arguments[1])
}

func greaterThan(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) > 0), nil
	}
	return NILL, fmt.Errorf("greaterThan : unsupported type %v or %v", arguments[0], arguments[1])
}

func greaterThanEqual(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		return B(first.CompareTo(second) >= 0), nil
	}
	return NILL, fmt.Errorf("greaterThanEqual : unsupported type %v or %v", arguments[0], arguments[1])
}

func cons(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	if len(arguments) == 0 {
		return P{}, nil
	} else if len(arguments) == 1 {
		return P{arguments[0], ENDED}, nil
	} else if len(arguments) == 2 {
		tail, ok := arguments[1].(P)
		if ok {
			return P{arguments[0], tail}, nil
		}
	}
	return P{}, nil
}

func first(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	pair, ok := arguments[0].(interfaces.Iterable)
	if ok {
		return pair.Head(), nil
	}
	return NILL, fmt.Errorf("first : %v is not of type Iterable", arguments[0])
}

func tail(arguments []interfaces.Value, sco interfaces.Scope) (interfaces.Value, error) {
	pair, ok := arguments[0].(interfaces.Iterable)
	if ok {
		if pair.HasTail() {
			return pair.Iterate(sco)
		}
		return ENDED, nil
	}
	return NILL, fmt.Errorf("tail : %v is not of type Iterable", arguments[0])
}

func apply(arguments []interfaces.Type, sco interfaces.Scope) (interfaces.Value, error) {
	if len(arguments) != 2 {
		return NILL, fmt.Errorf("apply : invalid number of arguments [%d of 2]", len(arguments))
	}
	list, err := evaluateToValue(arguments[1], sco)
	if err != nil {
		return NILL, err
	}
	s, okRef := arguments[0].(REF)
	p, okPair := list.(interfaces.Iterable)
	if !okRef {
		return NILL, fmt.Errorf("apply : expected function, found %v", arguments[0])
	} else if !okPair {
		return NILL, fmt.Errorf("apply : expected pair, found %v", list)
	}
	slice, err := p.ToSlice(sco.NewChildScope())
	if err != nil {
		return NILL, err
	}
	return evaluateToValue(&EXP{Function: s, Arguments: slice}, sco)
}

func iff(arguments []interfaces.Type, sco interfaces.Scope) (interfaces.Value, error) {
	test, err := evaluateToValue(arguments[0], sco)
	if err != nil {
		return NILL, err
	}
	if iff, iok := test.(B); iok {
		if iff {
			return evaluateToValue(arguments[1], sco)
		}
		return evaluateToValue(arguments[2], sco)
	}
	return NILL, fmt.Errorf("if : expected first argument to evaluate to boolean, recieved %v", test)
}

func def(arguments []interfaces.Type, sco interfaces.Scope) (interfaces.Value, error) {
	value, err := evaluateToValue(arguments[1], sco)
	if err != nil {
		return NILL, err
	}
	GlobalEnvironment.CreateRef(arguments[0].(REF), value)
	return NILL, nil
}

func do(arguments []interfaces.Type, sco interfaces.Scope) (interfaces.Value, error) {
	var result interfaces.Value
	for _, a := range arguments {
		next, err := evaluateToValue(a, sco.NewChildScope())
		if err != nil {
			return NILL, err
		}
		result = next
	}
	return result, nil
}

func rnge(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	start := arguments[0].(I)
	end := arguments[1].(I)
	if start < end {
		return LAZYP{
			start,
			&EXP{Function: REF("range"), Arguments: []interfaces.Type{
				I(start.Int() + 1),
				end,
			}}}, nil
	}
	return P{end, ENDED}, nil

}

func fn(arguments []interfaces.Type, sco interfaces.Scope) (interfaces.Value, error) {
	var argVec VEC
	if args, ok := arguments[0].(REF); ok {
		arg, err := args.Evaluate(sco)
		if err != nil {
			return NILL, err
		}
		argVec = arg.(VEC)
	} else {
		argVec = arguments[0].(VEC)
	}

	return FN{argVec, arguments[1].(interfaces.Evaluatable)}, nil
}

func filter(arguments []interfaces.Value, sco interfaces.Scope) (interfaces.Value, error) {
	if len(arguments) != 2 {
		return NILL, fmt.Errorf("filter : invalid number of arguments [%d of 2]", len(arguments))
	}
	ap, apok := arguments[0].(interfaces.Appliable)
	iter, iok := arguments[1].(interfaces.Iterable)

	var flt func(interfaces.Iterable) (interfaces.Iterable, error)
	flt = func(it interfaces.Iterable) (interfaces.Iterable, error) {
		head := it.Head()
		res, err := (&EXP{Function: ap, Arguments: []interfaces.Type{head}}).Evaluate(sco.NewChildScope())
		if err != nil {
			return ENDED, err
		}
		if include, iok := res.(B); iok {
			if bool(include) {
				if it.HasTail() {
					next, err := it.Iterate(sco)
					if err != nil {
						return ENDED, err
					}
					return LAZYP{
						head,
						&EXP{
							Function:  REF("filter"),
							Arguments: []interfaces.Type{ap, next},
						},
					}, nil
				}
				return &P{head, ENDED}, nil
			} else if it.HasTail() {
				next, err := it.Iterate(sco)
				if err != nil {
					return ENDED, err
				}
				return flt(next)
			}
			return ENDED, nil
		}
		return ENDED, fmt.Errorf("filter : expected boolean value, recieved %v", res)
	}

	if apok && iok {
		return flt(iter)
	}
	return NILL, fmt.Errorf("filter : expected function and list. Recieved %v, %v", arguments[0], arguments[1])
}

func mapp(arguments []interfaces.Value, sco interfaces.Scope) (interfaces.Value, error) {

	var mp func(interfaces.Value, interfaces.Iterable) *P
	mp = func(fn interfaces.Value, iterable interfaces.Iterable) *P {
		head := iterable.Head()
		res, _ := (&EXP{Function: fn, Arguments: []interfaces.Type{head}}).Evaluate(sco.NewChildScope())
		if iterable.HasTail() {
			next, _ := iterable.Iterate(sco)
			return &P{res, mp(fn, next)}
		}
		return &P{res, ENDED}
	}

	list := arguments[1]

	if pair, pok := list.(interfaces.Iterable); pok {
		return *mp(arguments[0], pair), nil
	}

	return ENDED, fmt.Errorf("map : expected function and list, recieved %v, %v", arguments[0], arguments[1])
}

func lazypair(arguments []interfaces.Type, sco interfaces.Scope) (interfaces.Value, error) {
	head, _ := evaluateToValue(arguments[0], sco)
	if len(arguments) > 1 {
		if tail, ok := arguments[1].(interfaces.Evaluatable); ok {
			return LAZYP{head, BindEvaluation(tail, sco)}, nil
		}
		return NILL, fmt.Errorf("lazypair : expected EXP got %v", arguments[1])
	}
	return LAZYP{head, nil}, nil
}

func macro(arguments []interfaces.Type, _ interfaces.Scope) (interfaces.Value, error) {
	return MAC{arguments[0].(VEC), arguments[1].(*EXP)}, nil
}

func printt(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	for _, arg := range arguments {
		fmt.Printf("%v\n", arg)
	}
	return NILL, nil
}

func empty(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	var arg interfaces.Value = arguments[0]
	if arg == nil {
		return B(true), nil
	}
	list := arg.(interfaces.Iterable)
	if list != ENDED {
		return B(false), nil
	}
	return B(true), nil
}

func take(arguments []interfaces.Value, sco interfaces.Scope) (interfaces.Value, error) {
	num, nok := arguments[0].(I)
	list, lok := arguments[1].(interfaces.Iterable)

	if nok && lok {
		if num > 1 && list.HasTail() {
			next, _ := list.Iterate(sco)
			return LAZYP{
				list.Head(),
				&EXP{
					Function: REF("take"),
					Arguments: []interfaces.Type{
						I(num - 1),
						next,
					},
				},
			}, nil
		}
		return P{list.Head(), ENDED}, nil

	}
	return ENDED, errors.New("take : expected number and list")
}

func let(arguments []interfaces.Type, sco interfaces.Scope) (interfaces.Value, error) {
	vectors, vok := arguments[0].(VEC)
	exp, eok := arguments[1].(interfaces.Evaluatable)

	childScope := sco.NewChildScope()

	if vok && eok {
		count := vectors.Count()
		if count%2 > 0 {
			return NILL, fmt.Errorf("let : expected an even number of items in vector, recieved %v", count)
		}
		for i := 0; i < count/2; i++ {
			val, _ := evaluateToValue(vectors.Get(i+1), sco)
			childScope.CreateRef(vectors.Get(i), val)
		}
		return exp.Evaluate(childScope)
	}
	return NILL, fmt.Errorf("let : expected VEC and EXP, received: %v, %v", arguments[0], arguments[1])
}
