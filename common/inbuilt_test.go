package common

import (
	"github.com/mikeyhu/glipso/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

// equals =

func Test_equals_NotEqual(t *testing.T) {
	//given
	exp := EXPBuild(REF("=")).withArgs(B(true), B(false)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func Test_equals_Equal(t *testing.T) {
	//given
	exp := EXPBuild(REF("=")).withArgs(B(true), B(true)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func Test_equals_ErrorsIfTypFesNotValid(t *testing.T) {
	//given
	exp := EXPBuild(REF("=")).withArgs(P{}, I(1)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "Equals : unsupported type P(<nil> <nil>) or 1")
}

// cons

func Test_cons_CreatesPairWithNil(t *testing.T) {
	//given
	exp := EXPBuild(REF("cons")).withArgs(I(1)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, P{I(1), ENDED}, result)
}

func Test_cons_CreatesPairWithTailPair(t *testing.T) {
	//given
	exp := EXPBuild(REF("cons")).withArgs(I(1), P{I(2), ENDED}).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(1), result.(P).head)
	assert.Equal(t, I(2), result.(P).tail.Head())
	assert.False(t, result.(P).tail.HasTail())
}

// first

func Test_first_RetrievesHeadOfPair(t *testing.T) {
	//given
	exp := EXPBuild(REF("first")).withArgs(P{I(3), ENDED}).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(3), result)
}

func Test_first_UnsupportedType(t *testing.T) {
	//given
	exp := EXPBuild(REF("first")).withArgs(B(true)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "first : true is not of type Iterable")
}

// tail

func Test_tail_RetrievesTailOfPair(t *testing.T) {
	//given
	exp := EXPBuild(REF("tail")).withArgs(P{I(5), &P{I(6), ENDED}}).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(6), result.(*P).head)
}

func Test_tail_OfListWithoutTailRetrievesEND(t *testing.T) {
	//given
	exp := EXPBuild(REF("tail")).withArgs(P{I(5), ENDED}).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, ENDED, result)
}

func Test_tail_UnsupportedType(t *testing.T) {
	//given
	exp := EXPBuild(REF("tail")).withArgs(B(true)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "tail : true is not of type Iterable")
}

// apply

func Test_apply_SendsListToFunction(t *testing.T) {
	//given
	exp := EXPBuild(REF("apply")).withArgs(REF("+"), P{I(2), &P{I(10), ENDED}}).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(12), result)
}

func Test_apply_ExpectedFunction(t *testing.T) {
	//given
	exp := EXPBuild(REF("apply")).withArgs(B(true), B(false)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "apply : expected function, found true")
}

func Test_apply_ExpectedPair(t *testing.T) {
	//given
	exp := EXPBuild(REF("apply")).withArgs(REF("+"), B(false)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "apply : expected pair, found false")
}

func Test_apply_InvalidNumberOfArguments(t *testing.T) {
	//given
	exp := EXPBuild(REF("apply")).withArgs().build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "apply : invalid number of arguments [0 of 2]")
}

// filter

func Test_filter_InvalidNumberOfArguments(t *testing.T) {
	//given
	exp := EXPBuild(REF("filter")).withArgs().build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "filter : invalid number of arguments [0 of 2]")
}

func Test_filter_UnsupportedTypes(t *testing.T) {
	//given
	exp := EXPBuild(REF("filter")).withArgs(B(true), B(false)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "filter : expected function and list. Recieved true, false")
}

func Test_filter_ExpectedBoolean(t *testing.T) {
	//given
	exp := EXPBuild(REF("filter")).withArgs(
		FNBuild().
			withArgs(REF("a")).
			withEXPBuilder(EXPBuild(REF("+")).withArgs(REF("a"))).build(),
		P{I(1), ENDED},
	).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "filter : expected boolean value, recieved 1")
}

// map

func Test_map_InvalidNumberOfArguments(t *testing.T) {
	//given
	exp := EXPBuild(REF("map")).withArgs().build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "map : invalid number of arguments [0 of 2]")
}

func Test_map_UnsupportedTypes(t *testing.T) {
	//given
	exp := EXPBuild(REF("map")).withArgs(B(true), B(false)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "map : expected function and list, recieved true, false")
}

// if

func Test_if_TrueReturnsSecondArgument(t *testing.T) {
	//given
	exp := EXPBuild(REF("if")).withArgs(B(true), I(1), I(2)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(1), result)
}

func Test_if_FalseReturnsThirdArgument(t *testing.T) {
	//given
	exp := EXPBuild(REF("if")).withArgs(B(false), I(1), I(2)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(2), result)
}

func Test_if_TrueEvaluatesRefRatherThanReturning(t *testing.T) {
	//given
	GlobalEnvironment.CreateRef(REF("a"), I(3))
	exp := EXPBuild(REF("if")).withArgs(B(true), REF("a"), REF("b")).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(3), result)
}

func Test_if_ErrorsIfNotBool(t *testing.T) {
	//given
	exp := EXPBuild(REF("if")).withArgs(
		EXPBuild(REF("+")).withArgs(I(1)).build(), I(1), I(2),
	).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "if : expected first argument to evaluate to boolean, recieved 1")
}

func Test_if_InvalidNumberOfArguments(t *testing.T) {
	//given
	exp := EXPBuild(REF("if")).withArgs(
		EXPBuild(REF("+")).withArgs(I(1)).build(),
	).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "if : invalid number of arguments [1 of 3]")
}

// def

func Test_def_RecordsReferences(t *testing.T) {
	//given
	exp := EXPBuild(REF("def")).withArgs(REF("one"), I(1)).build()
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

// do

func Test_do_ReturnsLastArgument(t *testing.T) {
	//given
	exp := EXPBuild(REF("do")).withArgs(I(1), I(2)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(2), result)
}

// range

func Test_range_ReturnsLazyPair(t *testing.T) {
	//given
	exp := EXPBuild(REF("range")).withArgs(I(1), I(10)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	lazyp, ok := result.(LAZYP)
	assert.True(t, ok)
	assert.NotNil(t, lazyp)
}

// multiply *

func Test_multiply_TwoIntegers(t *testing.T) {
	//given
	exp := EXPBuild(REF("*")).withArgs(I(2), I(3)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(6), result)
}

func Test_multiply_TwoFloats(t *testing.T) {
	//given
	exp := EXPBuild(REF("*")).withArgs(F(1.5), F(2.5)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, F(3.75), result)
}

func Test_multiply_IByF(t *testing.T) {
	//given
	exp := EXPBuild(REF("*")).withArgs(I(2), F(2.5)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, F(5), result)
}

func Test_multiply_FByI(t *testing.T) {
	//given
	exp := EXPBuild(REF("*")).withArgs(F(2.5), I(2)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, F(5), result)
}

// mod %

func Test_mod_Even(t *testing.T) {
	//given
	exp := EXPBuild(REF("%")).withArgs(I(4), I(2)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(0), result)
}

func Test_mod_Odd(t *testing.T) {
	//given
	exp := EXPBuild(REF("%")).withArgs(I(5), I(2)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, I(1), result)
}

func Test_mod_UnsupportedType(t *testing.T) {
	//given
	exp := EXPBuild(REF("%")).withArgs(B(true), B(false)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "mod : unsupported type")
}

func Test_mod_IncorrectNumberOfArguments(t *testing.T) {
	//given
	exp := EXPBuild(REF("%")).withArgs(I(7)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "% : invalid number of arguments [1 of 2]")
}

// lessThan <

func Test_lessThan_IntegersFirstIsHigher(t *testing.T) {
	//given
	exp := EXPBuild(REF("<")).withArgs(I(6), I(1)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func Test_lessThan_IntegersFirstIsLower(t *testing.T) {
	//given
	exp := EXPBuild(REF("<")).withArgs(I(1), I(6)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func Test_lessThan_IntegersArgumentsAreTheSame(t *testing.T) {
	//given
	exp := EXPBuild(REF("<")).withArgs(I(6), I(6)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func Test_lessThan_UnsupportedType(t *testing.T) {
	//given
	exp := EXPBuild(REF("<")).withArgs(B(true), B(false)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "lessThan : unsupported type true or false")
}

// greaterThan >

func Test_greaterThan_IntegersFirstIsHigher(t *testing.T) {
	//given
	exp := EXPBuild(REF(">")).withArgs(I(6), I(1)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func Test_greaterThan_IntegersFirstIsLower(t *testing.T) {
	//given
	exp := EXPBuild(REF(">")).withArgs(I(1), I(6)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func Test_greaterThan_IntegersArgumentsAreTheSame(t *testing.T) {
	//given
	exp := EXPBuild(REF(">")).withArgs(I(6), I(6)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func Test_greaterThan_UnsupportedType(t *testing.T) {
	//given
	exp := EXPBuild(REF(">")).withArgs(B(true), B(false)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "greaterThan : unsupported type true or false")
}

// lessThanEqual <=

func Test_lessThanEqual_IntegersFirstIsHigher(t *testing.T) {
	//given
	exp := EXPBuild(REF("<=")).withArgs(I(6), I(1)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func Test_lessThanEqual_IntegersFirstIsLower(t *testing.T) {
	//given
	exp := EXPBuild(REF("<=")).withArgs(I(1), I(6)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func Test_lessThanEqual_IntegersArgumentsAreTheSame(t *testing.T) {
	//given
	exp := EXPBuild(REF("<=")).withArgs(I(6), I(6)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func Test_lessThanEqual_UnsupportedType(t *testing.T) {
	//given
	exp := EXPBuild(REF("<=")).withArgs(B(true), B(false)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "lessThanEqual : unsupported type true or false")
}

// greaterThanEqual >=

func Test_greaterThanEqual_IntegersFirstIsHigher(t *testing.T) {
	//given
	exp := EXPBuild(REF(">=")).withArgs(I(6), I(1)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func Test_greaterThanEqual_IntegersFirstIsLower(t *testing.T) {
	//given
	exp := EXPBuild(REF(">=")).withArgs(I(1), I(6)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func Test_greaterThanEqual_IntegersArgumentsAreTheSame(t *testing.T) {
	//given
	exp := EXPBuild(REF(">=")).withArgs(I(6), I(6)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func Test_greaterThanEqual_UnsupportedType(t *testing.T) {
	//given
	exp := EXPBuild(REF(">=")).withArgs(B(true), B(false)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "greaterThanEqual : unsupported type true or false")
}

// print

func Test_print_ReturnsNILL(t *testing.T) {
	//given
	exp := EXPBuild(REF("print")).withArgs(I(1)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, NILL, result)
}

// empty

func Test_empty_ReturnsFalseOnLongList(t *testing.T) {
	//given
	exp := EXPBuild(REF("empty")).withArgs(P{I(1), P{I(2), ENDED}}).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func Test_empty_ReturnsFalseOnNonEmptyList(t *testing.T) {
	//given
	exp := EXPBuild(REF("empty")).withArgs(P{I(1), ENDED}).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(false), result)
}

func Test_empty_ReturnsTrueOnEmptyList(t *testing.T) {
	//given
	exp := EXPBuild(REF("empty")).withArgs(ENDED).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	assert.Equal(t, B(true), result)
}

func Test_empty_InvalidType(t *testing.T) {
	//given
	exp := EXPBuild(REF("empty")).withArgs(I(1)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "empty : expected Iterable got 1")
}

// take

func Test_take_NumberReturnsLazyPairWhenGivenRange(t *testing.T) {
	//given
	exp := EXPBuild(REF("take")).withArgs(
		I(3),
		EXPBuild(REF("range")).withArgs(I(1), I(5)).build(),
	).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.NoError(t, err)
	lazyp, ok := result.(LAZYP)
	assert.True(t, ok)
	assert.Equal(t, I(1), lazyp.Head())
}

func Test_take_ExpectsNumberAndPair(t *testing.T) {
	//given
	exp := EXPBuild(REF("take")).withArgs(I(3), I(4)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "take : expected number and list")
}

// lazypair

func Test_lazypair_ExpectsExpression(t *testing.T) {
	//given
	exp := EXPBuild(REF("lazypair")).withArgs(B(true), B(false)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "lazypair : expected EXP got false")
}

// let

func Test_let_ExpectsVectorAndExpression(t *testing.T) {
	//given
	exp := EXPBuild(REF("let")).withArgs(B(true), B(false)).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "let : expected VEC and EXP, received: true, false")
}

func Test_let_ExpectsEvenNumberSizedVector(t *testing.T) {
	//given
	exp := EXPBuild(REF("let")).withArgs(
		VEC{[]interfaces.Type{REF("a")}},
		&EXP{}).build()
	//when
	result, err := exp.Evaluate(GlobalEnvironment)
	//then
	assert.Equal(t, NILL, result)
	assert.EqualError(t, err, "let : expected an even number of items in vector, recieved 1")
}

// panic

func Test_panic_PanicsWithMessage(t *testing.T) {
	//given
	exp := EXPBuild(REF("panic")).withArgs(S("a message")).build()
	//when
	assert.Panics(t, func() {
		_, _ = exp.Evaluate(GlobalEnvironment)
	}, "Expected panic")
}

//assoc

func Test_assoc_ReturnsNewMAP(t *testing.T) {
	//given
	m, _ := InitialiseMAP([]interfaces.Value{SYM(":key"), I(10)})

	//when
	n, err := EXPBuild(REF("assoc")).withArgs(m, SYM(":key"), I(20)).build().Evaluate(GlobalEnvironment)

	//then
	assert.NoError(t, err)
	l, err := EXPBuild(SYM(":key")).withArgs(n).build().Evaluate(GlobalEnvironment)

	assert.NoError(t, err)

	assert.Equal(t, I(20), l)

}
