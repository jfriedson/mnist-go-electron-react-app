package module

import (
	"fmt"
)

type relu struct {
}

func (self *relu) Forward(inputAny any) (any, error) {
	// assert input is 1D slice of float32 for the time being
	input, ok := inputAny.([]float32)
	if !ok {
		return nil, fmt.Errorf("for now, relu input must be []float32")
	}

	inputLen := len(input)
	if inputLen <= 0 {
		return nil, fmt.Errorf("relu input must have at least 1 element")
	}

	output := make([]float32, inputLen)
	for i, x := range input {
		if x >= 0 {
			output[i] = x
		} else {
			output[i] = 0
		}
	}

	return output, nil
}

func NewRelu() *relu {
	return &relu{}
}
