package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

// Environment Provides a mechanism for creating and resolving variables
type Environment struct {
	id        int
	variables map[string]interfaces.Type
	parent    *Environment
}

// ResolveRef will try to resolve a provided reference to a value in this or parent scope
func (env Environment) ResolveRef(ref interfaces.Type) (interfaces.Type, bool) {
	if result, ok := env.variables[ref.(REF).String()]; ok {
		return result, true
	}
	if env.parent != nil {
		return env.parent.ResolveRef(ref)
	}
	return nil, false
}

// CreateRef will create a variable in this scope
func (env Environment) CreateRef(name interfaces.Type, arg interfaces.Type) interfaces.Type {
	env.variables[name.(REF).String()] = arg
	return name
}

// NewChildScope creates new scope that inherits from this one
func (env Environment) NewChildScope() interfaces.Scope {
	return Environment{
		NextScopeId(),
		map[string]interfaces.Type{},
		&env,
	}
}

// DisplayEnvironment is used to display environment information for internal debugging
func (env Environment) DisplayEnvironment() {
	env.displayEnvironment(0)
}

func (env Environment) displayEnvironment(i int) {
	for k := range env.variables {
		fmt.Printf("Scope[%d %d] %v := %v\n", env.id, i, k, env.variables[k])
	}
	if env.parent != nil {
		env.parent.displayEnvironment(i + 1)
	}
}

// GlobalEnvironment acts as the global scope for variables
var GlobalEnvironment Environment

func init() {
	GlobalEnvironment = Environment{
		id:        NextScopeId(),
		variables: map[string]interfaces.Type{},
	}
}

var scopeId = 0

func NextScopeId() int {
	scopeId = scopeId + 1
	return scopeId
}
