package module

import (
	"encoding/json"
	"fmt"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
)

type linear struct {
	weights [][]float32
	bias    []float32
}

func (self *linear) Forward(input any) (any, error) {
	// assert input is 1D slice of float32 for the time being
	inputAssert, ok := input.([]float32)
	if !ok {
		return nil, fmt.Errorf("for now, linear input must be []float32")
	}

	if len(inputAssert) != len(self.weights[0]) {
		return nil, fmt.Errorf("linear input size is incorrect")
	}

	// TODO: goroutine this puppy
	inputLen := len(inputAssert)
	outputLen := len(self.weights)
	output := make([]float32, outputLen)
	for out := 0; out < outputLen; out++ {
		var z float32 = 0
		for in := 0; in < inputLen; in++ {
			z += self.weights[out][in] * inputAssert[in]
		}
		if self.bias != nil {
			z += self.bias[out]
		}
		output[out] = z
	}

	return output, nil
}

func NewLinear(moduleInfo modelarch.ModuleInfo, modulesParams modelarch.ModulesParams) *linear {
	var name string
	raw, exists := moduleInfo.GetProp("name")
	if !exists {
		panic("linear name must be defined")
	} else {
		err := json.Unmarshal(raw, &name)
		if err != nil {
			panic(err)
		}
	}

	weightsRaw, exists := modulesParams[name+".weight"]
	if !exists {
		panic("linear weights must be defined")
	}
	var weights [][]float32
	err := json.Unmarshal(weightsRaw, &weights)
	if err != nil {
		panic("linear weights must be a two dimensional array")
	}
	if len(weights) < 1 {
		panic("linear weights must have at least 1 layer")
	}
	weightLayer0Dim := len(weights[0])
	for _, weightLayer := range weights {
		if len(weightLayer) != weightLayer0Dim {
			panic("linear weight layers must be the same size")
		}
	}

	var bias []float32
	biasRaw, exists := modulesParams[name+".bias"]
	if exists {
		err = json.Unmarshal(biasRaw, &bias)
		if err != nil {
			panic("linear bias must be a one dimensional array")
		}
	}
	if len(bias) != len(weights) {
		panic("linear bias size must match weight layers")
	}

	return &linear{weights, bias}
}
