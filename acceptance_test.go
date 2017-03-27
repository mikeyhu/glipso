package main

import (
	"github.com/mikeyhu/glipso/common"
	"github.com/mikeyhu/glipso/parser"
	"github.com/mikeyhu/glipso/prelude"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	exp, _ := parser.Parse("(+ 1 2 3 4 5)")
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(15), result)
}

func TestApplyAddNumbers(t *testing.T) {
	exp, _ := parser.Parse("(apply + (cons 1 (cons 2 (cons 3))))")
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(6), result)
}

func TestIfEvaluatesSecondExpression(t *testing.T) {
	code := `
	(if (= 1 1) (+ 2 2) (+ 3 3))
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(4), result)
}

func TestIfEvaluatesThirdExpression(t *testing.T) {
	code := `
	(if (= 1 2) (+ 2 2) (+ 3 3))
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(6), result)
}

func TestCreatesAndUsesVariable(t *testing.T) {
	code := `
	(do
		(def one 1)
		(def two 2)
		(+ one two))
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(3), result)
}

func TestSummingRange(t *testing.T) {
	exp, _ := parser.Parse("(apply + (range 1 5))")
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(15), result)
}

func TestGlobalAdd1Function(t *testing.T) {
	code := `
	(do
		(def add1 (fn [a] (+ 1 a)))
		(add1 5))
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(6), result)
}

func TestAnonymousAdd1Function(t *testing.T) {
	code := `
	((fn [a] (+ 1 a)) 5)
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(6), result)
}

func TestEvenFunctionEvaluatesEvenNumber(t *testing.T) {
	code := `
	(do
		(def even (fn [a] (= (% a 2) 0)))
		(even 2)
	)
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.B(true), result)
}

func TestEvenFunctionEvaluatesOddNumber(t *testing.T) {
	code := `
	(do
		(def even (fn [a] (= (% a 2) 0)))
		(even 1)
	)
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.B(false), result)
}

func TestFilterEvenNumbers(t *testing.T) {
	code := `
	(do
		(def even (fn [a] (= (% a 2) 0)))
		(apply + (filter even (cons 1 (cons 2 (cons 3 (cons 4))))))
	)
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(6), result)
}

func TestMapAdd1(t *testing.T) {
	code := `
	(do
		(def add1 (fn [a] (+ a 1)))
		(first (map add1 (cons 1)))
	)
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(2), result)
}

func TestLazyPairHasAccessToClosure(t *testing.T) {
	code := `
	(do
		(def hasclosure (fn [a b] (lazypair a (lazypair b))))
		(apply + (hasclosure 1 10))
	)
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(11), result)
}

func TestLazyPairCanBeUsedToCreateRange(t *testing.T) {
	code := `
	(do
		(def rangefn
			(fn [s e]
				(if (< s e)
					(lazypair s (rangefn (+ s 1) e))
					(cons s)
				)
			)
		)
		(=
			(apply + (rangefn 1 5))
			(apply + (range 1 5))
		)
	)
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.B(true), result)
}

func TestEmptyReturnsFalseWhenGivenAListWithContents(t *testing.T) {
	code := `(empty (range 1 5))`

	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.B(false), result)
}

func TestEmptyReturnsTrueWhenGivenAListWithNoContents(t *testing.T) {
	code := `(empty (filter (fn [num] (> num 10)) (range 1 5)))`

	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.B(true), result)
}

func BenchmarkSumRange(b *testing.B) {
	code := "(apply + (range 1 15))"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		exp, _ := parser.Parse(code)
		result := exp.Evaluate(common.GlobalEnvironment)
		assert.Equal(b, common.I(120), result)
	}
}

func BenchmarkSumRangefn(b *testing.B) {
	prelude.ParsePrelude(common.GlobalEnvironment)

	fn := `
	(do
		(def rangefn
			(fn [s e]
				(if (< s e)
					(lazypair s (rangefn (+ s 1) e))
					(cons s)
				)
			)
		)
	)`

	exp, _ := parser.Parse(fn)
	exp.Evaluate(common.GlobalEnvironment)

	code := "(apply + (rangefn 1 15))"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		exp, _ := parser.Parse(code)
		result := exp.Evaluate(common.GlobalEnvironment)
		assert.Equal(b, common.I(120), result)
	}
}

func BenchmarkParseCode(b *testing.B) {
	code := `
	(do
	    (defn notdivbyany [num listofdivs]
		(empty
		    (filter
			(fn [z] (= 0 z))
			(map (fn [head] (% num head)) listofdivs)
		    )
		)
	    )

	    (defn getprimes [num listofprimes]
		(if
		    (notdivbyany num listofprimes)
		    (lazypair num (getprimes (+ num 1) (cons num listofprimes)))
		    (getprimes (+ num 1) listofprimes)
		)
	    )

	    (map print (take 100 (getprimes 3 (cons 2))))
	)`
	for i := 0; i < b.N; i++ {
		_, err := parser.Parse(code)
		assert.NoError(b, err)
	}
}

func TestTakeReturnsFullListWhenSmallerThanTakeArgument(t *testing.T) {
	prelude.ParsePrelude(common.GlobalEnvironment)
	code := `(last (take 5 (range 1 3)))`

	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(3), result)
}

func TestTakeReturnsPartialList(t *testing.T) {
	prelude.ParsePrelude(common.GlobalEnvironment)
	code := `(last (take 3 (range 1 5)))`

	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(3), result)
}

func TestPanicsWhenFunctionNotFound(t *testing.T) {
	code := `(notafunction 1)`
	exp, _ := parser.Parse(code)
	assert.Panics(t, func() {
		exp.Evaluate(common.GlobalEnvironment)
	})
}

func TestReturningRefReturnsCorrectlyScopedValue(t *testing.T) {
	prelude.ParsePrelude(common.GlobalEnvironment)
	code := `
	(do
		(defn returnsA [A B] A)
		(defn hasA [A B] (returnsA B A))
		(hasA 1 2)
	)
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.I(2), result)
}

func TestLetAcceptsExpressionsInVectors(t *testing.T) {
	code := `(let [a (+ 1 2)] (= 3 a))`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.B(true), result)
}

func TestLetAcceptsValuesInVectors(t *testing.T) {
	code := `(let [a 3] (= 3 a))`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.B(true), result)
}
