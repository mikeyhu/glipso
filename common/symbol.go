package common

import (
	"fmt"
	"github.com/mikeyhu/glipso/interfaces"
)

// SYM is a symbol, beginning with a : and normally used as keys within maps
type SYM string

// IsType for SYM
func (s SYM) IsType() {}

// IsValue for SYM
func (s SYM) IsValue() {}

// String for SYM
func (s SYM) String() string {
	return string(s)
}

// Equals checks equality with another item of type Type
func (s SYM) Equals(o interfaces.Equalable) interfaces.Value {
	if other, ok := o.(SYM); ok {
		return B(s == other)
	}
	return B(false)
}

// Apply for SYM only works on a single argument of MAP, and looks up a value in the MAP keyed to the SYM
func (s SYM) Apply(arguments []interfaces.Type, env interfaces.Scope) (interfaces.Value, error) {
	if len(arguments) != 1 {
		return NILL, fmt.Errorf("SYM Apply : expected 1 argument, recieved %d", len(arguments))
	}
	val, _ := evaluateToValue(arguments[0], env)

	if mp, ok := val.(*MAP); ok {
		if v, found := mp.lookup(s); found {
			return v, nil
		}
		return NILL, nil

	}
	return NILL, fmt.Errorf("SYM Apply : expected MAP, recieved %v", arguments[0])
}

// MAP is an immutable hash-map, associating new entries with a MAP returns a new MAP
type MAP struct {
	store  map[interfaces.Equalable]interfaces.Value
	parent *MAP
}

// IsType for MAP
func (m *MAP) IsType() {}

// IsValue for MAP
func (m *MAP) IsValue() {}

// String representation of MAP
func (m *MAP) String() string {
	return fmt.Sprintf("%v", m.store)
}

func (m *MAP) lookup(k interfaces.Equalable) (interfaces.Value, bool) {
	if result, ok := m.store[k]; ok {
		return result, true
	}
	if m.parent != nil {
		return m.parent.lookup(k)
	}
	return NILL, false
}

func initialiseMAP(arguments []interfaces.Value) (*MAP, error) {
	count := len(arguments)
	if count%2 > 0 {
		return nil, fmt.Errorf("MAP Initialise : expected an even number of arguments, recieved %v", count)
	}
	mp := &MAP{map[interfaces.Equalable]interfaces.Value{}, nil}
	for i := 0; i < count; i += 2 {
		mp.store[arguments[i].(interfaces.Equalable)] = arguments[i+1]
	}
	return mp, nil
}

func (m *MAP) associate(arguments []interfaces.Value) (*MAP, error) {
	count := len(arguments)
	if count%2 > 0 {
		return nil, fmt.Errorf("MAP Initialise : expected an even number of arguments, recieved %v", count)
	}
	mp := &MAP{map[interfaces.Equalable]interfaces.Value{}, m}
	for i := 0; i < count; i += 2 {
		mp.store[arguments[i].(interfaces.Equalable)] = arguments[i+1]
	}
	return mp, nil
}

func (m *MAP) ToSlice(interfaces.Scope) ([]interfaces.Type, error) {
	keys := make([]interfaces.Type, len(m.store)*2)

	i := 0
	for k, v := range m.store {
		keys[i] = k
		i++
		keys[i] = v
		i++
	}
	return keys, nil
}
