package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

type Environment struct {
	variables map[string]interfaces.Argument
	parent    *Environment
}

func (env Environment) ResolveRef(ref interfaces.Argument) interfaces.Argument {
	if result, ok := env.variables[ref.(REF).String()]; ok {
		return result
	}
	if env.parent != nil {
		return env.parent.ResolveRef(ref)
	}
	panic(fmt.Sprintf("Unable to resolve reference %q", ref))
}

func (env Environment) CreateRef(name interfaces.Argument, arg interfaces.Argument) interfaces.Argument {
	env.variables[name.(REF).String()] = arg
	return name
}

func (env Environment) NewChildScope() interfaces.Scope {
	return Environment{
		map[string]interfaces.Argument{},
		&env,
	}
}

var GlobalEnvironment Environment

func init() {
	GlobalEnvironment = Environment{
		variables: map[string]interfaces.Argument{},
	}
}
