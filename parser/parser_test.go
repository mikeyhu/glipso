package parser

import (
	"github.com/mikeyhu/glipso/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParserReturnsExpressionFunctionName(t *testing.T) {
	result, err := Parse("(+ 1 2)")
	assert.NoError(t, err)
	assert.Equal(t, result.Function, common.REF("+"))
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
	assert.EqualError(t, err, "Unexpected EOF while parsing EXP")
}

func TestParserReturnsNestedExpression(t *testing.T) {
	result, err := Parse("(+ 1 (+ 1 3))")
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, 1, args[0].(common.I).Int(), "1st element")
	assert.Equal(t, 2, len(result.Arguments), "array length of outer arguments")
	assert.Equal(t, args[1].(*common.EXP).Function, common.REF("+"), "Nested Expression")
}

func TestParserReturnsSymbolsAsScopes(t *testing.T) {
	result, err := Parse("(+ symbol)")
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0].(common.REF).String(), "symbol")
}

func TestParserReturnsVectorsContainingIntegers(t *testing.T) {
	result, err := Parse("(+ [1 2])")
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0].(common.VEC).Get(0), common.I(1))
	assert.Equal(t, args[0].(common.VEC).Get(1), common.I(2))
}

func TestParserReturnsVectorsContainingStrings(t *testing.T) {
	result, err := Parse(`(+ ["hello" "world"])`)
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0].(common.VEC).Get(0), common.S("hello"))
	assert.Equal(t, args[0].(common.VEC).Get(1), common.S("world"))
}

func TestParserReturnsVectorWithSymbolsInside(t *testing.T) {
	result, err := Parse("(fn [a])")
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0].(common.VEC).Get(0), common.REF("a"))
}

func TestParserReturnsQuotedTextAsStrings(t *testing.T) {
	result, err := Parse(`(+ "hello" "world")`)
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0], common.S("hello"))
	assert.Equal(t, args[1], common.S("world"))
}

func TestParserHandlesFunctionsWithDashes(t *testing.T) {
	result, err := Parse(`(a-function 1 2)`)
	assert.NoError(t, err)
	assert.Equal(t, common.REF("a-function"), result.Function)
	args := result.Arguments
	assert.Equal(t, common.I(1), args[0])
	assert.Equal(t, common.I(2), args[1])
}

func TestParserHandlesBooleans(t *testing.T) {
	result, err := Parse(`(= true false)`)
	assert.NoError(t, err)
	assert.Equal(t, common.REF("="), result.Function)
	args := result.Arguments
	assert.Equal(t, args[0], common.B(true))
	assert.Equal(t, args[1], common.B(false))
}
