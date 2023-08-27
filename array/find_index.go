package array

//go:generate ifacemaker -f ./find_index.go -s FindIndex[T] -i "IFindIndex[T any]" -p interfaces -o ./interfaces/find_index.go
//go:generate mockgen -source=./interfaces/find_index.go -destination=./mocks/find_index.go -package=mocks

type FindIndex[T any] struct {
}

func NewFindIndex[T any]() *FindIndex[T] {
	return &FindIndex[T]{}
}

func (f *FindIndex[T]) FindIndex(values []T, compFunc func(comp T) bool) int {
	for i, val := range values {
		if compFunc(val) {
			return i
		}
	}
	return -1
}
