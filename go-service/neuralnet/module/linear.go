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

func (self *linear) Forward(inputAny any) (any, error) {
	// assert input is 1D slice of float32 for the time being
	input, ok := inputAny.([]float32)
	if !ok {
		return nil, fmt.Errorf("for now, linear input must be []float32")
	}

	inFeatures := len(self.weights[0])
	if len(input) != inFeatures {
		return nil, fmt.Errorf("linear input size is incorrect")
	}

	// TODO: goroutine this puppy
	outFeatures := len(self.weights)
	output := make([]float32, outFeatures)
	for out := range outFeatures {
		var z float32 = 0
		for in := 0; in < inFeatures; in++ {
			z += self.weights[out][in] * input[in]
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

	outFeatures := len(weights)
	if outFeatures < 1 {
		panic("linear weights out has length 0")
	}
	inFeatures := len(weights[0])
	for _, weightsDim1 := range weights {
		if len(weightsDim1) != inFeatures {
			panic("linear weights in_features must be consistent")
		}
	}

	var bias []float32
	biasRaw, exists := modulesParams[name+".bias"]
	if exists {
		err = json.Unmarshal(biasRaw, &bias)
		if err != nil {
			panic("linear bias must be a one dimensional array")
		}

		if len(bias) != outFeatures {
			panic("linear bias size must match weight out_features")
		}
	}

	return &linear{weights, bias}
}
