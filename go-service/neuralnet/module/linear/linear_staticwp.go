package linear

import (
	"encoding/json"
	"reflect"
	"runtime"
	"sync"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
)

type linearStaticWp struct {
	weights [][]float32
	bias    []float32
	jobs    chan linearStaticWpJob
	wg      *sync.WaitGroup
}

func (linearStaticWp linearStaticWp) Forward(inputPtr any) any {
	inputPtrVal := reflect.ValueOf(inputPtr)
	if inputPtrVal.Kind() != reflect.Pointer || inputPtrVal.IsNil() {
		panic("Linear: input must be non-nil pointer to []float32")
	}

	input := inputPtrVal.Elem().Interface().([]float32)

	inFeatures := len(linearStaticWp.weights[0])
	outFeatures := len(linearStaticWp.weights)

	if len(input) != inFeatures {
		panic("Linear: input size is incorrect")
	}

	output := make([]float32, outFeatures)

	for out := range outFeatures {
		linearStaticWp.wg.Add(1)
		linearStaticWp.jobs <- linearStaticWpJob{input, &output, out}
	}
	linearStaticWp.wg.Wait()

	return output
}

type linearStaticWpJob struct {
	input  []float32
	output *[]float32
	out    int
}

func (linearStaticWp linearStaticWp) linearStaticWpWorker() {
	inFeatures := len(linearStaticWp.weights[0])

	for j := range linearStaticWp.jobs {
		var z float32 = 0
		for in := range inFeatures {
			z += linearStaticWp.weights[j.out][in] * j.input[in]
		}
		if linearStaticWp.bias != nil {
			z += linearStaticWp.bias[j.out]
		}
		(*j.output)[j.out] = z

		linearStaticWp.wg.Done()
	}
}

func NewLinearStaticWp(moduleInfo modelarch.ModuleInfo, modulesParams modelarch.ModulesParams) linearStaticWp {
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

	jobs := make(chan linearStaticWpJob, outFeatures)
	var wg sync.WaitGroup

	linearStaticWp := linearStaticWp{weights, bias, jobs, &wg}

	numWorkers := runtime.NumCPU()
	for range numWorkers {
		go linearStaticWp.linearStaticWpWorker()
	}

	return linearStaticWp
}
