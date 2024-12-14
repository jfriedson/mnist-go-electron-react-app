package module

import (
	"encoding/json"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
)

type maxpool2d struct {
	dim int
}

func (maxpool2d maxpool2d) Forward(inputPtr any) any {
	return nil
	// inputPtrVal := reflect.ValueOf(inputPtr)
	// if inputPtrVal.Kind() != reflect.Pointer || inputPtrVal.IsNil() {
	// 	return nil, fmt.Errorf("LogSoftmax: input must be a non-nil pointer")
	// }

	// var max float32 = input[0]
	// for _, x := range input {
	// 	if x > max {
	// 		max = x
	// 	}
	// }

	// var sumexp float64 = 0
	// for _, x := range input {
	// 	sumexp += math.Exp(float64(x - max))
	// }
	// logsumexp := math.Log(sumexp)

	// output := make([]float32, length)
	// for i, x := range input {
	// 	output[i] = x - max - float32(logsumexp)
	// }

	// return output, nil
}

func NewMaxPool2d(moduleInfo modelarch.ModuleInfo) maxpool2d {
	var dim int

	raw, exists := moduleInfo.GetProp("dim")
	if !exists {
		panic("MaxPool2d: dim must be defined")
	} else {
		err := json.Unmarshal(raw, &dim)
		if err != nil {
			panic("MaxPool2d: dim must be a number")
		}
	}

	if dim < 0 {
		panic("MaxPool2d: start_dim must be 0 or greater")
	}

	return maxpool2d{dim}
}
