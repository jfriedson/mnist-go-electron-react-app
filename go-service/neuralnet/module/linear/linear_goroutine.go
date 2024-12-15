package linear

import (
	"encoding/json"
	"reflect"
	"sync"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
)

type linearGoroutine struct {
	weights [][]float32
	bias    []float32
}

func (linearGoroutine linearGoroutine) Forward(inputPtr any) any {
	inputPtrVal := reflect.ValueOf(inputPtr)
	if inputPtrVal.Kind() != reflect.Pointer || inputPtrVal.IsNil() {
		panic("Linear: input must be non-nil pointer to []float32")
	}

	input := inputPtrVal.Elem().Interface().([]float32)

	inFeatures := len(linearGoroutine.weights[0])
	if len(input) != inFeatures {
		panic("Linear: input size is incorrect")
	}

	outFeatures := len(linearGoroutine.weights)
	output := make([]float32, outFeatures)
	var wg sync.WaitGroup
	for out := range outFeatures {
		wg.Add(1)
		go linearGoroutine.linearGoroutine(&wg, input, &output, out)
	}
	wg.Wait()

	return output
}

func (linearGoroutine linearGoroutine) linearGoroutine(wg *sync.WaitGroup,
	input []float32, output *[]float32, out int) {

	defer wg.Done()

	inFeatures := len(linearGoroutine.weights[0])

	var z float32 = 0
	for in := range inFeatures {
		z += linearGoroutine.weights[out][in] * input[in]
	}
	if linearGoroutine.bias != nil {
		z += linearGoroutine.bias[out]
	}
	(*output)[out] = z
}

func NewLinearGoroutine(moduleInfo modelarch.ModuleInfo, modulesParams modelarch.ModulesParams) linearGoroutine {
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

	return linearGoroutine{weights, bias}
}
