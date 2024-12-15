package linear

import (
	"encoding/json"
	"reflect"
	"runtime"
	"sync"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
)

type linearWorkerpool struct {
	weights [][]float32
	bias    []float32
}

func (linearWorkerpool linearWorkerpool) Forward(inputPtr any) any {
	inputPtrVal := reflect.ValueOf(inputPtr)
	if inputPtrVal.Kind() != reflect.Pointer || inputPtrVal.IsNil() {
		panic("Linear: input must be non-nil pointer to []float32")
	}

	input := inputPtrVal.Elem().Interface().([]float32)

	inFeatures := len(linearWorkerpool.weights[0])
	outFeatures := len(linearWorkerpool.weights)

	if len(input) != inFeatures {
		panic("Linear: input size is incorrect")
	}

	output := make([]float32, outFeatures)

	jobs := make(chan linearWorkerpoolJob, outFeatures)
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()
	for range numWorkers {
		wg.Add(1)
		go linearWorkerpool.linearWorkerpoolWorker(jobs, &wg, input, &output)
	}

	for out := range outFeatures {
		jobs <- linearWorkerpoolJob{out}
	}
	close(jobs)

	wg.Wait()

	return output
}

type linearWorkerpoolJob struct {
	out int
}

func (linearWorkerpool linearWorkerpool) linearWorkerpoolWorker(jobs <-chan linearWorkerpoolJob, wg *sync.WaitGroup,
	input []float32, output *[]float32) {

	defer wg.Done()

	inFeatures := len(linearWorkerpool.weights[0])

	for j := range jobs {
		var z float32 = 0
		for in := range inFeatures {
			z += linearWorkerpool.weights[j.out][in] * input[in]
		}
		if linearWorkerpool.bias != nil {
			z += linearWorkerpool.bias[j.out]
		}
		(*output)[j.out] = z
	}
}

func NewLinearWorkerpool(moduleInfo modelarch.ModuleInfo, modulesParams modelarch.ModulesParams) linearWorkerpool {
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

	return linearWorkerpool{weights, bias}
}
