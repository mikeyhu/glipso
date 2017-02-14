package interfaces

type Evaluatable interface {
	Evaluate() Argument
}

type Argument interface {
	IsArg()
}

type Iterable interface {
	Iterate() Iterable
	ToSlice() []Argument
}
