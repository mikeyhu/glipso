package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

type Environment struct {
	variables map[string]interfaces.Argument
}

func (env Environment) resolveRef(ref REF) interfaces.Argument {
	if result, ok := env.variables[ref.String()]; ok {
		return result
	} else {
		panic(fmt.Sprintf("Unable to resolve reference %q", ref))
	}
}

func (env Environment) createRef(name REF, arg interfaces.Argument) REF {
	env.variables[name.String()] = arg
	return name
}

var GlobalEnvironment Environment

func init() {
	GlobalEnvironment = Environment{
		variables: map[string]interfaces.Argument{},
	}
}
