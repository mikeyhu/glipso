package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_I_CompareTo_Equals(t *testing.T) {
	//given
	a := I(100)
	b := I(100)

	//when
	result, err := a.CompareTo(b)

	//then
	assert.Equal(t, 0, result)
	assert.NoError(t, err)
}

func Test_I_CompareTo_F(t *testing.T) {
	//given
	a := I(100)
	b := F(100)

	//when
	result, err := a.CompareTo(b)

	//then
	assert.Equal(t, 0, result)
	assert.NoError(t, err)
}

func Test_I_CompareTo_InvalidType(t *testing.T) {
	//given
	a := I(100)
	b := S("100")

	//when
	_, err := a.CompareTo(b)

	//then
	assert.EqualError(t, err, "CompareTo : Cannot compare 100 to 100")
}

func Test_F_CompareTo_Equals(t *testing.T) {
	//given
	a := F(100.01)
	b := F(100.01)

	//when
	result, err := a.CompareTo(b)

	//then
	assert.Equal(t, 0, result)
	assert.NoError(t, err)
}

func Test_F_CompareTo_I(t *testing.T) {
	//given
	a := F(100)
	b := I(100)

	//when
	result, err := a.CompareTo(b)

	//then
	assert.Equal(t, 0, result)
	assert.NoError(t, err)
}

func Test_F_CompareTo_InvalidType(t *testing.T) {
	//given
	a := F(100)
	b := S("100")

	//when
	_, err := a.CompareTo(b)

	//then
	assert.EqualError(t, err, "CompareTo : Cannot compare 100.000000 to 100")
}
