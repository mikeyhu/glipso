package interfaces

type Evaluatable interface {
	Evaluate() Argument
}

type Argument interface {
	GetInteger() int
}
