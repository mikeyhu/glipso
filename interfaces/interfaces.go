package interfaces

// Type interfaces are Types within glipso
type Type interface {
	IsType()
	String() string
}

type Value interface {
	IsType()
	String() string
	IsValue()
}

// Iterable interfaces are pairs, lazypairs and anything that can be iterated or converted to a slice
type Iterable interface {
	IsType()
	String() string
	IsValue()
	Head() Value
	HasTail() bool
	Iterate(Scope) (Iterable, error)
	ToSlice(Scope) ([]Type, error)
}

// Equalable interfaces are types that can be checked for sameness
type Equalable interface {
	Equals(Equalable) Value
}

// Evaluatable interfaces are things such as Expressions or References that can be evaluated to return a Value
type Evaluatable interface {
	Evaluate(Scope) (Value, error)
}

// Comparable interfaces are types that can be checked for equality and order
type Comparable interface {
	CompareTo(Comparable) (int, error)
}

// Expandable interfaces are types that will be expanded prior to evaluation
type Expandable interface {
	Expand([]Type) Evaluatable
}

// Appliable interfaces can be applied by expressions to return Values
type Appliable interface {
	IsType()
	String() string
	IsValue()
	Apply([]Type, Scope) (Value, error)
}

// Scope interfaces provice a mechanism for creating variables and looking up references
type Scope interface {
	ResolveRef(argument Type) (Value, bool)
	CreateRef(ref Type, arg Value) Type
	NewChildScope() Scope
	String() string
}
