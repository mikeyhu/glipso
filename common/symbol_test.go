package common

import (
	"github.com/mikeyhu/glipso/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SYM_Equals_SYM(t *testing.T) {
	//given
	a := SYM(":key")
	b := SYM(":key")

	//when
	result := a.Equals(b)

	//then
	assert.Equal(t, B(true), result)
}

func Test_SYM_Equals_DifferentSYM(t *testing.T) {
	//given
	a := SYM(":key")
	b := SYM(":differentKey")

	//when
	result := a.Equals(b)

	//then
	assert.Equal(t, B(false), result)
}

func Test_MAP_InitialiseAndLookup(t *testing.T) {
	//given
	k := SYM(":key")
	v := I(10)

	//when
	m, err := InitialiseMAP([]interfaces.Value{k, v})

	//then
	assert.NoError(t, err)
	result, ok := m.Lookup(SYM(":key"))
	assert.Equal(t, true, ok)
	assert.Equal(t, v, result)
}

func Test_MAP_AssociateAndLookup(t *testing.T) {
	//given
	m, _ := InitialiseMAP([]interfaces.Value{SYM(":a"), I(1)})

	//when
	n, err := m.Associate([]interfaces.Value{SYM(":b"), I(2)})
	//then
	assert.NoError(t, err)

	a, ok := n.Lookup(SYM(":a"))
	assert.Equal(t, true, ok)
	assert.Equal(t, I(1), a)

	b, ok := n.Lookup(SYM(":b"))
	assert.Equal(t, true, ok)
	assert.Equal(t, I(2), b)
}

func Test_MAP_AssociateIsImmutable(t *testing.T) {
	//given
	m, _ := InitialiseMAP([]interfaces.Value{SYM(":a"), I(1)})

	//when
	n, err := m.Associate([]interfaces.Value{SYM(":b"), I(2)})
	//then
	assert.NoError(t, err)

	_, ok := m.Lookup(SYM(":b"))
	assert.Equal(t, false, ok)

	b, ok := n.Lookup(SYM(":b"))
	assert.Equal(t, true, ok)
	assert.Equal(t, I(2), b)
}

func Test_SYM_Apply_FindsValue(t *testing.T) {
	//given
	k := SYM(":key")
	v := I(10)

	//when
	m, _ := InitialiseMAP([]interfaces.Value{k, v})
	exp := EXPBuild(k).withArgs(m).build()
	result, err := exp.Evaluate(GlobalEnvironment.NewChildScope())

	//then
	assert.NoError(t, err)
	assert.Equal(t, v, result)
}

func Test_SYM_Apply_FindsNILL(t *testing.T) {
	//given
	k := SYM(":key")
	m := &MAP{}

	//when
	exp := EXPBuild(k).withArgs(m).build()
	result, err := exp.Evaluate(GlobalEnvironment.NewChildScope())

	//then
	assert.NoError(t, err)
	assert.Equal(t, NILL, result)
}

func Test_SYM_Apply_InvalidType(t *testing.T) {
	//given
	k := SYM(":key")
	p := P{I(1), ENDED}

	//when
	exp := EXPBuild(k).withArgs(p).build()
	_, err := exp.Evaluate(GlobalEnvironment.NewChildScope())

	//then
	assert.EqualError(t, err, "SYM Apply : expected MAP, recieved P(1 <END>)")
}
