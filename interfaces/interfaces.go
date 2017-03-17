package interfaces

// Evaluatable interfaces are things such as Expressions or References that can be evaluated to return a Type
type Evaluatable interface {
	Evaluate(Scope) Type
}

// Type interfaces are Types within glipso
type Type interface {
	IsType()
	String() string
}

// Iterable interfaces are pairs, lazypairs and anything that can be iterated or converted to a slice
type Iterable interface {
	IsType()
	String() string
	Head() Type
	HasTail() bool
	Iterate(Scope) Iterable
	ToSlice(Scope) []Type
}

// Scope interfaces provice a mechanism for creating variables and looking up references
type Scope interface {
	ResolveRef(argument Type) (Type, bool)
	CreateRef(ref Type, arg Type) Type
	NewChildScope() Scope
	String() string
}

// Equalable interfaces are types that can be checked for sameness
type Equalable interface {
	Equals(Equalable) Type
}

// Comparable interfaces are types that can be checked for equality and order
type Comparable interface {
	CompareTo(Comparable) int
}

// Expandable interfaces are types that will be expanded prior to evaluation
type Expandable interface {
	Expand([]Type) Type
}

// Function interfaces can be applied to expressions
type Function interface {
	IsType()
	String() string
	Apply([]Type, Scope) Type
}
