package common

import (
	"github.com/mikeyhu/glipso/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqualityNotEqual(t *testing.T) {
	//given
	exp := EXP{Function: REF("="), Arguments: []interfaces.Type{B(true), B(false)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func TestEqualityEqual(t *testing.T) {
	//given
	exp := EXP{Function: REF("="), Arguments: []interfaces.Type{B(true), B(true)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func TestEqualityErrorsIfTypFesNotValid(t *testing.T) {
	//given
	exp := EXP{Function: REF("="), Arguments: []interfaces.Type{P{}, I(1)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "Equals : unsupported type P(<nil> <nil>) or 1")
}

func TestConsCreatesPairWithNil(t *testing.T) {
	//given
	exp := EXP{Function: REF("cons"), Arguments: []interfaces.Type{I(1)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, P{I(1), ENDED}, result)
}

func TestConsCreatesPairWithTailPair(t *testing.T) {
	//given
	exp := EXP{Function: REF("cons"), Arguments: []interfaces.Type{I(1), P{I(2), ENDED}}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(1), result.(P).head)
	assert.Equal(t, I(2), result.(P).tail.Head())
	assert.False(t, result.(P).tail.HasTail())
}

func TestFirstRetrievesHeadOfPair(t *testing.T) {
	//given
	exp := EXP{Function: REF("first"), Arguments: []interfaces.Type{P{I(3), ENDED}}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(3), result)
}

func TestFirstUnsupportedType(t *testing.T) {
	//given
	exp := EXP{Function: REF("first"), Arguments: []interfaces.Type{B(true)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "first : true is not of type Iterable")
}

func TestTailRetrievesTailOfPair(t *testing.T) {
	//given
	exp := EXP{Function: REF("tail"), Arguments: []interfaces.Type{P{I(5), &P{I(6), ENDED}}}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(6), result.(*P).head)
}

func TestTailOfListWithoutTailRetrievesEND(t *testing.T) {
	//given
	exp := EXP{Function: REF("tail"), Arguments: []interfaces.Type{P{I(5), ENDED}}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, ENDED, result)
}

func TestTailUnsupportedType(t *testing.T) {
	//given
	exp := EXP{Function: REF("tail"), Arguments: []interfaces.Type{B(true)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "tail : true is not of type Iterable")
}

func TestApplySendsListToFunction(t *testing.T) {
	//given
	exp := EXP{Function: REF("apply"), Arguments: []interfaces.Type{REF("+"), P{I(2), &P{I(10), ENDED}}}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(12), result)
}

func TestApplyExpectedFunction(t *testing.T) {
	//given
	exp := EXP{Function: REF("apply"), Arguments: []interfaces.Type{B(true), B(false)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "apply : expected function, found true")
}

func TestApplyExpectedPair(t *testing.T) {
	//given
	exp := EXP{Function: REF("apply"), Arguments: []interfaces.Type{REF("+"), B(false)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "apply : expected pair, found false")
}

func TestApplyInvalidNumberOfArguments(t *testing.T) {
	//given
	exp := EXP{Function: REF("apply"), Arguments: []interfaces.Type{}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "apply : invalid number of arguments [0 of 2]")
}

func TestFilterInvalidNumberOfArguments(t *testing.T) {
	//given
	exp := EXP{Function: REF("filter"), Arguments: []interfaces.Type{}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "filter : invalid number of arguments [0 of 2]")
}

func TestFilterUnsupportedTypes(t *testing.T) {
	//given
	exp := EXP{Function: REF("filter"), Arguments: []interfaces.Type{B(true), B(false)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "filter : expected function and list. Recieved true, false")
}

func TestFilterExpectedBoolean(t *testing.T) {
	//given
	exp := EXP{Function: REF("filter"), Arguments: []interfaces.Type{
		FN{
			VEC{[]interfaces.Type{REF("a")}},
			&EXP{
				REF("+"),
				[]interfaces.Type{REF("a")},
			},
		}, P{I(1), ENDED}}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "filter : expected boolean value, recieved 1")
}

func TestMapUnsupportedTypes(t *testing.T) {
	//given
	exp := EXP{Function: REF("map"), Arguments: []interfaces.Type{B(true), B(false)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "map : expected function and list, recieved true, false")
}

func TestIfTrueReturnsSecondArgument(t *testing.T) {
	//given
	exp := EXP{Function: REF("if"), Arguments: []interfaces.Type{B(true), I(1), I(2)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(1), result)
}

func TestIfFalseReturnsThirdArgument(t *testing.T) {
	//given
	exp := EXP{Function: REF("if"), Arguments: []interfaces.Type{B(false), I(1), I(2)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(2), result)
}

func TestIfTrueEvaluatesRefRatherThanReturning(t *testing.T) {
	//given
	GlobalEnvironment.CreateRef(REF("a"), I(3))
	exp := EXP{Function: REF("if"), Arguments: []interfaces.Type{B(true), REF("a"), REF("b")}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(3), result)
}

func TestDefRecordsReferences(t *testing.T) {
	//given
	exp := EXP{Function: REF("def"), Arguments: []interfaces.Type{REF("one"), I(1)}}
	//when
	_, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	//when
	resolved, ok := GlobalEnvironment.ResolveRef(REF("one"))
	//then
	assert.Equal(t, I(1), resolved)
	assert.True(t, ok)
}

func TestDoReturnsLastArgument(t *testing.T) {
	//given
	exp := EXP{Function: REF("do"), Arguments: []interfaces.Type{I(1), I(2)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(2), result)
}

func TestRangeReturnsLazyPair(t *testing.T) {
	//given
	exp := EXP{Function: REF("range"), Arguments: []interfaces.Type{I(1), I(10)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t,
		LAZYP{I(1), &EXP{Function: REF("range"), Arguments: []interfaces.Type{I(2), I(10)}}},
		result)
}

func TestEvaluateMultiply(t *testing.T) {
	//given
	exp := EXP{Function: REF("*"), Arguments: []interfaces.Type{I(2), I(3)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(6), result)
}

func TestEvaluateModEven(t *testing.T) {
	//given
	exp := EXP{Function: REF("%"), Arguments: []interfaces.Type{I(4), I(2)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(0), result)
}

func TestEvaluateModOdd(t *testing.T) {
	//given
	exp := EXP{Function: REF("%"), Arguments: []interfaces.Type{I(5), I(2)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(1), result)
}

func TestEvaluateModUnsupportedType(t *testing.T) {
	//given
	exp := EXP{Function: REF("%"), Arguments: []interfaces.Type{B(true), B(false)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "mod : unsupported type")
}

func TestLessThanIntegersFirstIsHigher(t *testing.T) {
	//given
	exp := EXP{Function: REF("<"), Arguments: []interfaces.Type{I(6), I(1)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func TestLessThanIntegersFirstIsLower(t *testing.T) {
	//given
	//when
	//then
	exp := EXP{Function: REF("<"), Arguments: []interfaces.Type{I(1), I(6)}}
	result, err := exp.Evaluate(GlobalEnvironment)
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func TestLessThanIntegersArgumentsAreTheSame(t *testing.T) {
	//given
	exp := EXP{Function: REF("<"), Arguments: []interfaces.Type{I(6), I(6)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func TestLessThanUnsupportedType(t *testing.T) {
	//given
	exp := EXP{Function: REF("<"), Arguments: []interfaces.Type{B(true), B(false)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "lessThan : unsupported type true or false")
}

func TestGreaterThanIntegersFirstIsHigher(t *testing.T) {
	//given
	exp := EXP{Function: REF(">"), Arguments: []interfaces.Type{I(6), I(1)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func TestGreaterThanIntegersFirstIsLower(t *testing.T) {
	//given
	exp := EXP{Function: REF(">"), Arguments: []interfaces.Type{I(1), I(6)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func TestGreaterThanIntegersArgumentsAreTheSame(t *testing.T) {
	//given
	exp := EXP{Function: REF(">"), Arguments: []interfaces.Type{I(6), I(6)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func TestGreaterThanUnsupportedType(t *testing.T) {
	//given
	exp := EXP{Function: REF(">"), Arguments: []interfaces.Type{B(true), B(false)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "greaterThan : unsupported type true or false")
}

func TestLessThanEqualIntegersFirstIsHigher(t *testing.T) {
	//given
	exp := EXP{Function: REF("<="), Arguments: []interfaces.Type{I(6), I(1)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func TestLessThanEqualIntegersFirstIsLower(t *testing.T) {
	//given
	exp := EXP{Function: REF("<="), Arguments: []interfaces.Type{I(1), I(6)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func TestLessThanEqualIntegersArgumentsAreTheSame(t *testing.T) {
	//given
	exp := EXP{Function: REF("<="), Arguments: []interfaces.Type{I(6), I(6)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func TestLessThanEqualUnsupportedType(t *testing.T) {
	//given
	exp := EXP{Function: REF("<="), Arguments: []interfaces.Type{B(true), B(false)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "lessThanEqual : unsupported type true or false")
}

func TestGreaterThanEqualIntegersFirstIsHigher(t *testing.T) {
	//given
	exp := EXP{Function: REF(">="), Arguments: []interfaces.Type{I(6), I(1)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func TestGreaterThanEqualIntegersFirstIsLower(t *testing.T) {
	//given
	exp := EXP{Function: REF(">="), Arguments: []interfaces.Type{I(1), I(6)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func TestGreaterThanEqualIntegersArgumentsAreTheSame(t *testing.T) {
	//given
	exp := EXP{Function: REF(">="), Arguments: []interfaces.Type{I(6), I(6)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func TestGreaterEqualThanUnsupportedType(t *testing.T) {
	//given
	exp := EXP{Function: REF(">="), Arguments: []interfaces.Type{B(true), B(false)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "greaterThanEqual : unsupported type true or false")
}

func TestPrintReturnsNILL(t *testing.T) {
	//given
	exp := EXP{Function: REF("print"), Arguments: []interfaces.Type{I(1)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, NILL, result)
}

func TestEmptyReturnsFalseOnLongList(t *testing.T) {
	//given
	exp := EXP{Function: REF("empty"), Arguments: []interfaces.Type{P{I(1), P{I(2), ENDED}}}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func TestEmptyReturnsFalseOnNonEmptyList(t *testing.T) {
	//given
	exp := EXP{Function: REF("empty"), Arguments: []interfaces.Type{P{I(1), ENDED}}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func TestEmptyReturnsTrueOnEmptyList(t *testing.T) {
	//given
	exp := EXP{Function: REF("empty"), Arguments: []interfaces.Type{ENDED}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func TestTakeNumberReturnsLazyPairWhenGivenRange(t *testing.T) {
	//given
	exp := EXP{Function: REF("take"), Arguments: []interfaces.Type{
		I(3),
		&EXP{Function: REF("range"), Arguments: []interfaces.Type{I(1), I(5)}}},
	}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	lazyp, ok := result.(LAZYP)
	assert.True(t, ok)
	assert.Equal(t, I(1), lazyp.Head())
}

func TestTakeExpectsNumberAndPair(t *testing.T) {
	//given
	exp := EXP{Function: REF("take"), Arguments: []interfaces.Type{
		I(3),
		I(4),
	}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "take : expected number and list")
}

func TestLazyPairExpectsExpression(t *testing.T) {
	//given
	exp := EXP{Function: REF("lazypair"), Arguments: []interfaces.Type{B(true), B(false)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "lazypair : expected EXP got false")
}

func TestLetExpectsVectorAndExpression(t *testing.T) {
	//given
	exp := EXP{Function: REF("let"), Arguments: []interfaces.Type{B(true), B(false)}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "let : expected VEC and EXP, received: true, false")
}

func TestLetExpectsEvenNumberSizedVector(t *testing.T) {
	//given
	exp := EXP{Function: REF("let"), Arguments: []interfaces.Type{
		VEC{[]interfaces.Type{REF("a")}},
		&EXP{}}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "let : expected an even number of items in vector, recieved 1")
}

func TestIfErrorsIfNotBool(t *testing.T) {
	//given
	exp := EXP{Function: REF("if"), Arguments: []interfaces.Type{
		&EXP{Function: REF("+"), Arguments: []interfaces.Type{I(1)}},
	}}
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "if : expected first argument to evaluate to boolean, recieved 1")
}
