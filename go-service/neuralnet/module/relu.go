package module

import (
	"fmt"
)

type relu struct {
}

func (self *relu) Forward(input any) (any, error) {
	// assert input is 1D slice of float32 for the time being
	inputAssert, ok := input.([]float32)
	if !ok {
		return nil, fmt.Errorf("for now, relu input must be []float32")
	}

	if len(inputAssert) <= 0 {
		return nil, fmt.Errorf("relu input must have at least 1 element")
	}

	output := make([]float32, len(inputAssert))
	for idx, in := range inputAssert {
		if in >= 0 {
			output[idx] = in
		} else {
			output[idx] = 0
		}
	}

	return output, nil
}

func NewRelu() *relu {
	return &relu{}
}
