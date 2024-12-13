package linalg

import (
	"fmt"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/number"
)

type Mat[T number.Number] struct {
	Data [][]T
	Dims []int
}

func (self *Mat[T]) T() *Mat[T] {
	return self
}

func (self *Mat[T]) Add(other *Mat[T]) (*Mat[T], error) {
	if len(self.Dims) != len(other.Dims) {
		return nil, fmt.Errorf("matrices' dimensions do not match {} and {}", self.Dims, other.Dims)
	}

	return nil, nil
}
