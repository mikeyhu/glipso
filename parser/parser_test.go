package parser

import (
	"github.com/mikeyhu/glipso/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParserReturnsExpressionFunctionName(t *testing.T) {
	result, err := Parse("(+ 1 2)")
	assert.NoError(t, err)
	assert.Equal(t, result.FunctionName, "+")
}

func TestParserReturnsExpressionArguments(t *testing.T) {
	result, err := Parse("(+ 1 3)")
	assert.NoError(t, err)
	assert.Equal(t, result.Arguments[0].(common.I).Int(), 1, "1st element")
	assert.Equal(t, result.Arguments[1].(common.I).Int(), 3, "2nd element")
	assert.Equal(t, len(result.Arguments), 2, "array length of arguments")
}

func TestParserReturnsErrorWhenNoClosingBrackets(t *testing.T) {
	_, err := Parse("(+ 1")
	assert.EqualError(t, err, "Expected end of Expression")
}

func TestParserReturnsNestedExpression(t *testing.T) {
	result, err := Parse("(+ 1 (+ 1 3))")
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0].(common.I).Int(), 1, "1st element")
	assert.Equal(t, len(result.Arguments), 2, "array length of outer arguements")
	assert.Equal(t, args[1].(common.EXP).FunctionName, "+", "Nested Expression")
}

func TestParserReturnsSymbolsAsScopes(t *testing.T) {
	result, err := Parse("(+ symbol)")
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0].(common.REF).String(), "symbol")
}

func TestParserReturnsVector(t *testing.T) {
	result, err := Parse("(+ [1 2])")
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0].(common.VEC).Get(0), common.I(1))
	assert.Equal(t, args[0].(common.VEC).Get(1), common.I(2))
}

func TestParserReturnsVectorWithSymbolsInside(t *testing.T) {
	result, err := Parse("(fn [a])")
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0].(common.VEC).Get(0), common.REF("a"))
}