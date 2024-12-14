package module

import (
	"encoding/json"
	"math"
	"reflect"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
)

type logsoftmax struct {
	dim int
}

// TODO: convert to in-place op
// TODO: function currently expects one dimensional input. calculate softmax across dim
func (logsoftmax logsoftmax) Forward(inputPtr any) any {
	inputPtrVal := reflect.ValueOf(inputPtr)
	if inputPtrVal.Kind() != reflect.Pointer || inputPtrVal.IsNil() {
		panic("LogSoftmax: input must be a non-nil pointer to []float32")
	}

	input := inputPtrVal.Elem().Interface().([]float32)

	length := len(input)
	if length <= 0 {
		panic("LogSoftmax: input must have at least 1 element")
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

	return output
}

func NewLogSoftmax(moduleInfo modelarch.ModuleInfo) *logsoftmax {
	var dim int

	raw, exists := moduleInfo.GetProp("dim")
	if !exists {
		panic("LogSoftmax: dim must be defined")
	} else {
		err := json.Unmarshal(raw, &dim)
		if err != nil {
			panic("LogSoftmax: dim must be a number")
		}
	}

	if dim < 0 {
		panic("LogSoftmax: start_dim must be 0 or greater")
	}

	return &logsoftmax{dim}
}
