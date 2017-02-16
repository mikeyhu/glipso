package interfaces

type Evaluatable interface {
	Evaluate(Scope) Argument
}

type Argument interface {
	IsArg()
}

type Iterable interface {
	Iterate(Scope) Iterable
	ToSlice(Scope) []Argument
}

type Scope interface {
	ResolveRef(argument Argument) Argument
	CreateRef(ref Argument, arg Argument) Argument
	NewChildScope() Scope
}