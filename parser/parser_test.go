package parser

import (
	"github.com/mikeyhu/glipso/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Parser_ExpressionFunctionName(t *testing.T) {
	result, err := Parse("(+ 1 2)")
	assert.NoError(t, err)
	assert.Equal(t, result.Function, common.REF("+"))
}

func Test_Parser_ExpressionArguments(t *testing.T) {
	result, err := Parse("(+ 1 3)")
	assert.NoError(t, err)
	assert.Equal(t, result.Arguments[0].(common.I).Int(), 1, "1st element")
	assert.Equal(t, result.Arguments[1].(common.I).Int(), 3, "2nd element")
	assert.Equal(t, len(result.Arguments), 2, "array length of arguments")
}

func Test_Parser_ErrorWhenNoClosingBrackets(t *testing.T) {
	_, err := Parse("(+ 1")
	assert.EqualError(t, err, "Unexpected EOF while parsing EXP")
}

func Test_Parser_NestedExpression(t *testing.T) {
	result, err := Parse("(+ 1 (+ 1 3))")
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, 1, args[0].(common.I).Int(), "1st element")
	assert.Equal(t, 2, len(result.Arguments), "array length of outer arguments")
	assert.Equal(t, args[1].(*common.EXP).Function, common.REF("+"), "Nested Expression")
}

func Test_Parser_REF(t *testing.T) {
	result, err := Parse("(+ ref)")
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0].(common.REF).String(), "ref")
}

func Test_Parser_VectorContainingIntegers(t *testing.T) {
	result, err := Parse("(+ [1 2])")
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0].(common.VEC).Get(0), common.I(1))
	assert.Equal(t, args[0].(common.VEC).Get(1), common.I(2))
}

func Test_Parser_VectorContainingStrings(t *testing.T) {
	result, err := Parse(`(+ ["hello" "world"])`)
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0].(common.VEC).Get(0), common.S("hello"))
	assert.Equal(t, args[0].(common.VEC).Get(1), common.S("world"))
}

func Test_Parser_VectorWithSymbolsInside(t *testing.T) {
	result, err := Parse("(fn [a])")
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0].(common.VEC).Get(0), common.REF("a"))
}

func Test_Parser_QuotedTextAsStrings(t *testing.T) {
	result, err := Parse(`(+ "hello" "world")`)
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0], common.S("hello"))
	assert.Equal(t, args[1], common.S("world"))
}

func Test_Parser_FunctionNameWithDashes(t *testing.T) {
	result, err := Parse(`(a-function 1 2)`)
	assert.NoError(t, err)
	assert.Equal(t, common.REF("a-function"), result.Function)
	args := result.Arguments
	assert.Equal(t, common.I(1), args[0])
	assert.Equal(t, common.I(2), args[1])
}

func Test_Parser_Booleans(t *testing.T) {
	result, err := Parse(`(= true false)`)
	assert.NoError(t, err)
	assert.Equal(t, common.REF("="), result.Function)
	args := result.Arguments
	assert.Equal(t, args[0], common.B(true))
	assert.Equal(t, args[1], common.B(false))
}

func Test_Parser_QuotedTextWithEscapedCharacters(t *testing.T) {
	result, err := Parse(`(+ "hi \"there\"")`)
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0], common.S(`hi "there"`))
}

func Test_Parser_Floats(t *testing.T) {
	result, err := Parse(`(+ 1.01 2.02)`)
	assert.NoError(t, err)
	args := result.Arguments
	assert.Equal(t, args[0], common.F(1.01))
	assert.Equal(t, args[1], common.F(2.02))
}

func Test_Parser_Symbol(t *testing.T) {
	result, err := Parse(`(:key :value)`)
	assert.NoError(t, err)
	assert.Equal(t, result.Function, common.SYM(":key"))
	args := result.Arguments
	assert.Equal(t, args[0], common.SYM(":value"))
}
