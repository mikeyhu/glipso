package main

import (
	"github.com/mikeyhu/mekkanism/common"
	"github.com/mikeyhu/mekkanism/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	exp, _ := parser.Parse("(+ 1 2 3 4 5)")
	result := exp.Evaluate()
	assert.Equal(t, common.I(15), result)
}

func TestApplyAddNumbers(t *testing.T) {
	exp, _ := parser.Parse("(apply + (cons 1 (cons 2 (cons 3))))")
	result := exp.Evaluate()
	assert.Equal(t, common.I(6), result)
}

func TestIfEvaluatesSecondExpression(t *testing.T) {
	code := `
	(if (= 1 1) (+ 2 2) (+ 3 3))
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate()
	assert.Equal(t, common.I(4), result)
}

func TestIfEvaluatesThirdExpression(t *testing.T) {
	code := `
	(if (= 1 2) (+ 2 2) (+ 3 3))
	`
	exp, _ := parser.Parse(code)
	result := exp.Evaluate()
	assert.Equal(t, common.I(6), result)
}
