package main

import (
	"github.com/mikeyhu/glipso/common"
	"github.com/mikeyhu/glipso/parser"
	"github.com/mikeyhu/glipso/prelude"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AddNumbers(t *testing.T) {
	exp, err := parser.Parse("(+ 1 2 3 4 5)")
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.I(15), result)
}

func Test_ApplyAddNumbers(t *testing.T) {
	exp, err := parser.Parse("(apply + (cons 1 (cons 2 (cons 3))))")
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.I(6), result)
}

func Test_IfEvaluatesSecondExpression(t *testing.T) {
	code := `
	(if (= 1 1) (+ 2 2) (+ 3 3))
	`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.I(4), result)
}

func Test_IfEvaluatesThirdExpression(t *testing.T) {
	code := `
	(if (= 1 2) (+ 2 2) (+ 3 3))
	`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.I(6), result)
}

func Test_CreatesAndUsesVariable(t *testing.T) {
	code := `
	(do
		(def one 1)
		(def two 2)
		(+ one two))
	`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.I(3), result)
}

func Test_SummingRange(t *testing.T) {
	exp, err := parser.Parse("(apply + (range 1 5))")
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.I(15), result)
}

func Test_CreateAdd1Function(t *testing.T) {
	code := `
	(do
		(def add1 (fn [a] (+ 1 a)))
		(add1 5))
	`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.I(6), result)
}

func Test_AnonymousAdd1Function(t *testing.T) {
	code := `
	((fn [a] (+ 1 a)) 5)
	`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.I(6), result)
}

func Test_EvenFunctionEvaluatesEvenNumber(t *testing.T) {
	code := `
	(do
		(def even (fn [a] (= (% a 2) 0)))
		(even 2)
	)
	`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.B(true), result)
}

func Test_EvenFunctionEvaluatesOddNumber(t *testing.T) {
	code := `
	(do
		(def even (fn [a] (= (% a 2) 0)))
		(even 1)
	)
	`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.B(false), result)
}

func Test_FilterEvenNumbers(t *testing.T) {
	code := `
	(do
		(def even (fn [a] (= (% a 2) 0)))
		(apply + (filter even (cons 1 (cons 2 (cons 3 (cons 4))))))
	)
	`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.I(6), result)
}

func Test_MappingAdd1(t *testing.T) {
	code := `
	(do
		(def add1 (fn [a] (+ a 1)))
		(first (map add1 (cons 1)))
	)
	`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.I(2), result)
}

func Test_LazyPairHasAccessToClosure(t *testing.T) {
	code := `
	(do
		(def hasclosure (fn [a b] (lazypair a (lazypair b))))
		(apply + (hasclosure 1 10))
	)
	`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.I(11), result)
}

func Test_LazyPairCanBeUsedToCreateRange(t *testing.T) {
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
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.B(true), result)
}

func Test_EmptyReturnsFalseWhenGivenAListWithContents(t *testing.T) {
	code := `(empty (range 1 5))`

	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.B(false), result)
}

func Test_EmptyReturnsTrueWhenGivenAListWithNoContents(t *testing.T) {
	code := `(empty (filter (fn [num] (> num 10)) (range 1 5)))`

	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.B(true), result)
}

func BenchmarkSumRange(b *testing.B) {
	code := "(apply + (range 1 15))"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		exp, err := parser.Parse(code)
		assert.NoError(b, err)
		result, err := exp.Evaluate(common.GlobalEnvironment)
		assert.NoError(b, err)
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

	_, _ = parser.Parse(fn)

	code := "(apply + (rangefn 1 15))"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		exp, err := parser.Parse(code)
		assert.NoError(b, err)
		result, err := exp.Evaluate(common.GlobalEnvironment)
		assert.NoError(b, err)
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

	    (apply print (take 100 (getprimes 3 (cons 2))))
	)`
	for i := 0; i < b.N; i++ {
		_, err := parser.Parse(code)
		assert.NoError(b, err)
	}
}

func Test_TakeReturnsFullListWhenSmallerThanTakeArgument(t *testing.T) {
	prelude.ParsePrelude(common.GlobalEnvironment)
	code := `(last (take 5 (range 1 3)))`

	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.I(3), result)
}

func Test_TakeReturnsPartialList(t *testing.T) {
	prelude.ParsePrelude(common.GlobalEnvironment)
	code := `(last (take 3 (range 1 5)))`

	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.I(3), result)
}

func Test_ErrorsWhenFunctionNotFound(t *testing.T) {
	code := `(notafunction 1)`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)

	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.Equal(t, common.NILL, result)
	assert.EqualError(t, err, "evaluate : function 'notafunction' not found")

}

func Test_ReturningRefReturnsCorrectlyScopedValue(t *testing.T) {
	prelude.ParsePrelude(common.GlobalEnvironment)
	code := `
	(do
		(defn returnsA [A B] A)
		(defn hasA [A B] (returnsA B A))
		(hasA 1 2)
	)
	`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.I(2), result)
}

func Test_LetAcceptsExpressionsInVectors(t *testing.T) {
	code := `(let [a (+ 1 2)] (= 3 a))`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.B(true), result)
}

func Test_LetAcceptsValuesInVectors(t *testing.T) {
	code := `(let [a 3] (= 3 a))`
	exp, err := parser.Parse(code)
	assert.NoError(t, err)
	result, err := exp.Evaluate(common.GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, common.B(true), result)
}
