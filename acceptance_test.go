package main

import (
	"github.com/mikeyhu/glipso/common"
	"github.com/mikeyhu/glipso/parser"
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
