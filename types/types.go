package types

import (
	"github.com/mikeyhu/mekkanism/interfaces"
)

type Argtype struct {
	Expression interfaces.Evaluatable
	Integer    int
}

func (argtype *Argtype) GetInteger() int {
	return argtype.Integer
}
