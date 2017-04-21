package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// I Integer

func Test_I_Equals_I(t *testing.T) {
	//given
	a := I(100)
	b := I(100)

	//when
	result := a.Equals(b)

	//then
	assert.Equal(t, B(true), result)
}

func Test_I_Equals_F(t *testing.T) {
	//given
	a := I(200)
	b := F(200)

	//when
	result := a.Equals(b)

	//then
	assert.Equal(t, B(true), result)
}

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

func Test_I_Add_I(t *testing.T) {
	//given
	a := I(100)
	b := I(200)

	//when
	result := a.Add(b)

	//then
	assert.Equal(t, I(300), result)
}

func Test_I_Subtract_I(t *testing.T) {
	//given
	a := I(100)
	b := I(50)

	//when
	result := a.Subtract(b)

	//then
	assert.Equal(t, I(50), result)
}

func Test_I_Multiply_I(t *testing.T) {
	//given
	a := I(2)
	b := I(3)

	//when
	result := a.Multiply(b)

	//then
	assert.Equal(t, I(6), result)
}

func Test_I_Divide_I(t *testing.T) {
	//given
	a := I(100)
	b := I(2)

	//when
	result := a.Divide(b)

	//then
	assert.Equal(t, I(50), result)
}

func Test_I_Mod_I(t *testing.T) {
	//given
	a := I(15)
	b := I(7)

	//when
	result := a.Mod(b)

	//then
	assert.Equal(t, I(1), result)
}

// F Float

func Test_F_Equals_F(t *testing.T) {
	//given
	a := F(100.1)
	b := F(100.1)

	//when
	result := a.Equals(b)

	//then
	assert.Equal(t, B(true), result)
}

func Test_F_Equals_I(t *testing.T) {
	//given
	a := F(200)
	b := I(200)

	//when
	result := a.Equals(b)

	//then
	assert.Equal(t, B(true), result)
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
