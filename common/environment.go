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
	if env.variables != nil {
		if result, ok := env.variables[ref.(REF).String()]; ok {
			return result, true
		}
	}
	if env.parent != nil {
		return env.parent.ResolveRef(ref)
	}
	return nil, false
}

// CreateRef will create a variable in this scope
func (env *Environment) CreateRef(name interfaces.Type, arg interfaces.Type) interfaces.Type {
	if DEBUG {
		fmt.Printf("Adding %v %v to %v\n", name, arg, env)
	}
	if env.variables == nil {
		env.variables = map[string]interfaces.Type{}
	}
	env.variables[name.(REF).String()] = arg
	return name
}

// NewChildScope creates new scope that inherits from this one
func (env Environment) NewChildScope() interfaces.Scope {
	id := nextScopeID()
	if DEBUG {
		fmt.Printf("New scope %d from %d\n", id, env.id)
	}
	return &Environment{
		id:     id,
		parent: &env,
	}
}

// DisplayEnvironment is used to display environment information for internal debugging
func (env Environment) DisplayEnvironment() {
	if DEBUG {
		env.displayEnvironment(0)
	}
}

func (env Environment) displayEnvironment(i int) {
	for k := range env.variables {
		fmt.Printf("Scope[%d %d] %v := %v\n", env.id, i, k, env.variables[k])
	}
	if env.parent != nil {
		env.parent.displayEnvironment(i + 1)
	}
}

func (env Environment) String() string {
	return fmt.Sprintf("ENV{%d}", env.id)
}

// GlobalEnvironment acts as the global scope for variables
var GlobalEnvironment *Environment

func init() {
	GlobalEnvironment = &Environment{
		id:        nextScopeID(),
		variables: map[string]interfaces.Type{},
	}
}

var scopeID = 0

func nextScopeID() int {
	scopeID = scopeID + 1
	return scopeID
}

// DisplayDiagnostics outputs Information about the Scopes
func (env Environment) DisplayDiagnostics() {
	if DEBUG {
		fmt.Printf("Total number of scopes created: %v", scopeID)
	}
}
