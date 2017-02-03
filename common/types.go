package common

type I int

func (i I) IsArg() {}
func (i I) Int() int {
	return int(i)
}
