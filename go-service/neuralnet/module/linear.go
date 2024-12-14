package module

import (
	"encoding/json"
	"reflect"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
)

type linear struct {
	weights [][]float32
	bias    []float32
}

func (linear linear) Forward(inputPtr any) any {
	inputPtrVal := reflect.ValueOf(inputPtr)
	if inputPtrVal.Kind() != reflect.Pointer || inputPtrVal.IsNil() {
		panic("Linear: input must be non-nil pointer to []float32")
	}

	input := inputPtrVal.Elem().Interface().([]float32)

	inFeatures := len(linear.weights[0])
	if len(input) != inFeatures {
		panic("Linear: input size is incorrect")
	}

	// TODO: goroutine this puppy
	outFeatures := len(linear.weights)
	output := make([]float32, outFeatures)
	for out := range outFeatures {
		var z float32 = 0
		for in := range inFeatures {
			z += linear.weights[out][in] * input[in]
		}
		if linear.bias != nil {
			z += linear.bias[out]
		}
		output[out] = z
	}

	return output
}

func NewLinear(moduleInfo modelarch.ModuleInfo, modulesParams modelarch.ModulesParams) linear {
	var name string
	raw, exists := moduleInfo.GetProp("name")
	if !exists {
		panic("Linear: name must be defined")
	} else {
		err := json.Unmarshal(raw, &name)
		if err != nil {
			panic(err)
		}
	}

	weightsRaw, exists := modulesParams[name+".weight"]
	if !exists {
		panic("Linear: weights must be defined")
	}
	var weights [][]float32
	err := json.Unmarshal(weightsRaw, &weights)
	if err != nil {
		panic("Linear: weights must be a two dimensional array")
	}

	outFeatures := len(weights)
	if outFeatures < 1 {
		panic("Linear: weights out has length 0")
	}
	inFeatures := len(weights[0])
	for _, weightsDim1 := range weights {
		if len(weightsDim1) != inFeatures {
			panic("Linear: weights in_features must be consistent")
		}
	}

	var bias []float32
	biasRaw, exists := modulesParams[name+".bias"]
	if exists {
		err = json.Unmarshal(biasRaw, &bias)
		if err != nil {
			panic("Linear: bias must be a one dimensional array")
		}

		if len(bias) != outFeatures {
			panic("Linear: bias size must match weight out_features")
		}
	}

	return linear{weights, bias}
}
