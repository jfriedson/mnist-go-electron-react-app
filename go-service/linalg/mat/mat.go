package mat

import (
	"fmt"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/linalg"
)

type Mat[T linalg.Number] struct {
	Data	[][]T
	Dims	int	
	Dim		[Dims]int
}

func (self *Mat[T]) T() *Mat[T] {
	if self.Dim != other.Dim {
		return nil, fmt.Errorf("matrices' dimensions do not match")
	}

	for 
	self.Data + other.Data
}

func (self *Mat[T]) Add(other *Mat[T]) (*Mat[T], error) {
	if self.Dim != other.Dim {
		return nil, fmt.Errorf("matrices' dimensions do not match")
	}

	for 
	self.Data + other.Data
}
