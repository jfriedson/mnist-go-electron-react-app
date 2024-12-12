package module

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
)

type logsoftmax struct {
	dim int
}

func (self *logsoftmax) Forward(inputAny any) (any, error) {
	// assert input is 1D slice of float32 for the time being
	input, ok := inputAny.([]float32)
	if !ok {
		return nil, fmt.Errorf("for now, logsoftmax input must be []float32")
	}

	length := len(input)
	if length <= 0 {
		return nil, fmt.Errorf("logsoftmax input must have at least 1 element")
	}

	var max float32 = input[0]
	for _, x := range input {
		if x > max {
			max = x
		}
	}

	var sumexp float64 = 0
	for _, x := range input {
		sumexp += math.Exp(float64(x - max))
	}
	logsumexp := math.Log(sumexp)

	output := make([]float32, length)
	for i, x := range input {
		output[i] = x - max - float32(logsumexp)
	}

	return output, nil
}

func NewLogSoftmax(moduleInfo modelarch.ModuleInfo) *logsoftmax {
	var dim int

	raw, exists := moduleInfo.GetProp("dim")
	if !exists {
		panic("dim must be defined")
	} else {
		err := json.Unmarshal(raw, &dim)
		if err != nil {
			panic("dim must be a number")
		}
	}

	if dim < 0 {
		panic("start_dim must be 0 or greater")
	}

	return &logsoftmax{dim}
}
