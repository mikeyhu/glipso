package interfaces

type Evaluatable interface {
	Evaluate(Scope) Type
}

type Type interface {
	IsType()
	String() string
}

type Iterable interface {
	Iterate(Scope) Iterable
	ToSlice(Scope) []Type
}

type Scope interface {
	ResolveRef(argument Type) (Type, bool)
	CreateRef(ref Type, arg Type) Type
	NewChildScope() Scope
}

type Equalable interface {
	Equals(Equalable) Type
}

type Comparable interface {
	CompareTo(Comparable) int
}
