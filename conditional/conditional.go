package conditional

//go:generate ifacemaker -f ./conditional.go -s Conditional[A] -i "IConditional[A any]" -p interfaces -o ./interfaces/conditional.go
//go:generate mockgen -source=./interfaces/conditional.go -destination=./mocks/conditional.go -package=mocks

type Conditional[A any] struct {
}

func NewConditional[A any]() *Conditional[A] {
	return &Conditional[A]{}
}

func (c *Conditional[A]) If(condition bool, trueRes A, falseRes A) A {
	if condition {
		return trueRes
	}
	if !condition {
		return falseRes
	}

	panic("Invalid conditional!")
}
