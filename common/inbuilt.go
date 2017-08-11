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
	addInbuilt(FI{name: "%", evaluator: mod, argumentCount: 2})
	addInbuilt(FI{name: "<", evaluator: lessThan})
	addInbuilt(FI{name: ">", evaluator: greaterThan})
	addInbuilt(FI{name: "<=", evaluator: lessThanEqual})
	addInbuilt(FI{name: ">=", evaluator: greaterThanEqual})
	addInbuilt(FI{name: "and", evaluator: and})
	addInbuilt(FI{name: "assoc", evaluator: assoc})
	addInbuilt(FI{name: "apply", lazyEvaluator: apply, argumentCount: 2})
	addInbuilt(FI{name: "cons", evaluator: cons})
	addInbuilt(FI{name: "def", lazyEvaluator: def, argumentCount: 2})
	addInbuilt(FI{name: "do", lazyEvaluator: do})
	addInbuilt(FI{name: "empty", evaluator: empty, argumentCount: 1})
	addInbuilt(FI{name: "if", lazyEvaluator: iff, argumentCount: 3})
	addInbuilt(FI{name: "filter", evaluator: filter, argumentCount: 2})
	addInbuilt(FI{name: "first", evaluator: first, argumentCount: 1})
	addInbuilt(FI{name: "get", evaluator: get, argumentCount: 2})
	addInbuilt(FI{name: "fn", lazyEvaluator: fn, argumentCount: 2})
	addInbuilt(FI{name: "hash-map", evaluator: hashmap})
	addInbuilt(FI{name: "lazypair", lazyEvaluator: lazypair})
	addInbuilt(FI{name: "let", lazyEvaluator: let, argumentCount: 2})
	addInbuilt(FI{name: "macro", lazyEvaluator: macro, argumentCount: 2})
	addInbuilt(FI{name: "map", evaluator: mapp, argumentCount: 2})
	addInbuilt(FI{name: "or", evaluator: or})
	addInbuilt(FI{name: "print", evaluator: printt})
	addInbuilt(FI{name: "panic", evaluator: panicc, argumentCount: 1})
	addInbuilt(FI{name: "range", evaluator: rnge, argumentCount: 2})
	addInbuilt(FI{name: "tail", evaluator: tail, argumentCount: 1})
	addInbuilt(FI{name: "take", evaluator: take, argumentCount: 2})
}

func addInbuilt(info FI) {
	inbuilt[REF(info.name)] = info
}

func plusAll(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	return numericFlatten(arguments, func(a interfaces.Numeric, b interfaces.Numeric) interfaces.Numeric {
		return a.Add(b)
	})
}

func minusAll(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	return numericFlatten(arguments, func(a interfaces.Numeric, b interfaces.Numeric) interfaces.Numeric {
		return a.Subtract(b)
	})
}

func multiplyAll(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	return numericFlatten(arguments, func(a interfaces.Numeric, b interfaces.Numeric) interfaces.Numeric {
		return a.Multiply(b)
	})
}

func mod(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	a, aok := arguments[0].(I)
	b, bok := arguments[1].(I)
	if aok && bok {
		return a.Mod(b), nil
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
		compare, err := first.CompareTo(second)
		if err != nil {
			return NILL, err
		}
		if err != nil {
			return NILL, err
		}
		return B(compare < 0), nil
	}
	return NILL, fmt.Errorf("lessThan : unsupported type %v or %v", arguments[0], arguments[1])
}

func lessThanEqual(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		compare, err := first.CompareTo(second)
		if err != nil {
			return NILL, err
		}
		return B(compare <= 0), nil
	}
	return NILL, fmt.Errorf("lessThanEqual : unsupported type %v or %v", arguments[0], arguments[1])
}

func greaterThan(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		compare, err := first.CompareTo(second)
		if err != nil {
			return NILL, err
		}
		return B(compare > 0), nil
	}
	return NILL, fmt.Errorf("greaterThan : unsupported type %v or %v", arguments[0], arguments[1])
}

func greaterThanEqual(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	first, fok := arguments[0].(interfaces.Comparable)
	second, sok := arguments[1].(interfaces.Comparable)
	if fok && sok {
		compare, err := first.CompareTo(second)
		if err != nil {
			return NILL, err
		}
		return B(compare >= 0), nil
	}
	return NILL, fmt.Errorf("greaterThanEqual : unsupported type %v or %v", arguments[0], arguments[1])
}

func cons(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	if len(arguments) == 0 {
		return ENDED, nil
	} else if len(arguments) == 1 {
		return P{arguments[0], ENDED}, nil
	} else if len(arguments) == 2 {
		tail, ok := arguments[1].(P)
		if ok {
			return P{arguments[0], tail}, nil
		}
	}
	return ENDED, nil
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
	list, err := evaluateToValue(arguments[1], sco)
	if err != nil {
		return NILL, err
	}
	s, okRef := arguments[0].(REF)
	p, okPair := list.(interfaces.Sliceable)
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

func rnge(arguments []interfaces.Value, sco interfaces.Scope) (interfaces.Value, error) {
	start := arguments[0].(I)
	end := arguments[1].(I)
	if start < end {
		return createLAZYP(sco, start, REF("range"), I(start.Int()+1), end), nil
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
	ap, apok := arguments[0].(interfaces.Appliable)
	iter, iok := arguments[1].(interfaces.Iterable)

	var flt func(interfaces.Iterable) (interfaces.Iterable, error)
	flt = func(it interfaces.Iterable) (interfaces.Iterable, error) {
		head := it.Head()
		res, err := evaluateToValue(&EXP{Function: ap, Arguments: []interfaces.Type{head}}, sco.NewChildScope())
		if err != nil {
			return ENDED, err
		}
		if include, iok := res.(B); iok {
			if it.HasTail() {
				next, err := it.Iterate(sco)
				if err != nil {
					return ENDED, err
				}
				if bool(include) {
					return createLAZYP(sco, head, REF("filter"), ap, next), nil
				}
				return flt(next)
			}
			if bool(include) {
				return &P{head, ENDED}, nil
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
	fn, fnok := arguments[0].(interfaces.Appliable)
	list, lok := arguments[1].(interfaces.Iterable)

	if fnok && lok {
		head := list.Head()
		res, err := evaluateToValue(&EXP{Function: fn, Arguments: []interfaces.Type{head}}, sco.NewChildScope())
		if err == nil {
			if !list.HasTail() {
				return &P{res, ENDED}, nil
			}
			next, err := list.Iterate(sco)
			if err == nil {
				return createLAZYP(sco, res, REF("map"), fn, next), nil
			}
		}
		return ENDED, err
	}
	return ENDED, fmt.Errorf("map : expected function and list, recieved %v, %v", arguments[0], arguments[1])
}

func lazypair(arguments []interfaces.Type, sco interfaces.Scope) (interfaces.Value, error) {
	head, err := evaluateToValue(arguments[0], sco)
	if err != nil {
		return NILL, err
	}
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
	list, ok := arguments[0].(interfaces.Iterable)
	if !ok {
		return ENDED, fmt.Errorf("empty : expected Iterable got %v", arguments[0])
	}
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
			next, err := list.Iterate(sco)
			if err != nil {
				return NILL, err
			}
			return createLAZYP(sco, list.Head(), REF("take"), num-1, next), nil
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
		count := vectors.count()
		if count%2 > 0 {
			return NILL, fmt.Errorf("let : expected an even number of items in vector, recieved %v", count)
		}
		for i := 0; i < count; i += 2 {
			val, err := evaluateToValue(vectors.Get(i+1), childScope)
			if err != nil {
				return NILL, err
			}
			childScope.CreateRef(vectors.Get(i), val)
		}
		return exp.Evaluate(childScope)
	}
	return NILL, fmt.Errorf("let : expected VEC and EXP, received: %v, %v", arguments[0], arguments[1])
}

func panicc(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	panic(arguments[0].String())
}

func hashmap(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	return initialiseMAP(arguments)
}

func assoc(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	mp, ok := arguments[0].(*MAP)

	if !ok {
		return NILL, fmt.Errorf("assoc : first argument should be a MAP")
	}
	return mp.associate(arguments[1:])
}

func get(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	k, kok := arguments[0].(interfaces.Equalable)
	mp, mok := arguments[1].(*MAP)

	if kok && mok {
		v, found := mp.lookup(k)
		if found {
			return v, nil
		}
		return NILL, nil
	}
	return NILL, fmt.Errorf("get : expected key and map")
}

func and(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	for _, a := range arguments {
		if asB, ok := a.(B); ok {
			if !asB {
				return B(false), nil
			}
		} else {
			return NILL, fmt.Errorf("and : expected all arguments to be B")
		}
	}
	return B(true), nil
}

func or(arguments []interfaces.Value, _ interfaces.Scope) (interfaces.Value, error) {
	for _, a := range arguments {
		if asB, ok := a.(B); ok {
			if asB {
				return B(true), nil
			}
		} else {
			return NILL, fmt.Errorf("or : expected all arguments to be B")
		}
	}
	return B(false), nil
}
