package conv2d

import (
	"encoding/json"
	"reflect"
	"runtime"
	"sync"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
)

type conv2dWorkerpool struct {
	weights [][][][]float32
	bias    []float32
}

// input is operated on in the format (channel, height, width)
// output is formatted (channel, height, width)
func (conv2dWorkerpool conv2dWorkerpool) Forward(inputPtr any) any {
	inputPtrVal := reflect.ValueOf(inputPtr)
	if inputPtrVal.Kind() != reflect.Pointer || inputPtrVal.IsNil() {
		panic("Conv2d: input must be non-nil pointer to [][][]float32")
	}

	inputAny := inputPtrVal.Elem().Interface()

	input := inputAny.([][][]float32)

	inChans := len(input)
	inHeight := len(input[0])
	inWidth := len(input[0][0])
	outChans := len(conv2dWorkerpool.weights)
	filters := len(conv2dWorkerpool.weights[0])
	kernelHeight := len(conv2dWorkerpool.weights[0][0])
	kernelWidth := len(conv2dWorkerpool.weights[0][0][0])

	if inChans != filters {
		panic("Conv2d: input channel does not match kernel count")
	}
	if inHeight < kernelHeight {
		panic("Conv2d: input height is smaller than kernel")
	}
	if inWidth < kernelWidth {
		panic("Conv2d: input width is smaller than kernel")
	}

	// initialize the output image
	output := make([][][]float32, outChans)
	outHeight := inHeight - (kernelHeight - 1)
	outWidth := inWidth - (kernelWidth - 1)
	for oCh := range outChans {
		output[oCh] = make([][]float32, outHeight)
		for oR := range outHeight {
			output[oCh][oR] = make([]float32, outWidth)
		}
	}

	numWorkers := runtime.NumCPU()
	jobs := make(chan conv2dWorkerpoolJob, outChans*outHeight*outWidth)

	var wg sync.WaitGroup
	for range numWorkers {
		wg.Add(1)
		go conv2dWorkerpool.worker(jobs, &wg, input, &output)
	}

	for oCh := range outChans {
		for oR := range outHeight {
			for oC := range outWidth {
				jobs <- conv2dWorkerpoolJob{oCh, oR, oC}
			}
		}
	}
	close(jobs)

	wg.Wait()

	return output
}

type conv2dWorkerpoolJob struct {
	oCh, oR, oC int
}

func (conv2dWorkerpool conv2dWorkerpool) worker(jobs <-chan conv2dWorkerpoolJob, wg *sync.WaitGroup,
	input [][][]float32, output *[][][]float32) {

	defer wg.Done()

	inChans := len(input)
	kernelHeight := len(conv2dWorkerpool.weights[0][0])
	kernelWidth := len(conv2dWorkerpool.weights[0][0][0])

	for j := range jobs {
		oCh := j.oCh
		oR := j.oR
		oC := j.oC

		var z float32 = 0
		for iCh := range inChans {
			for kR := range kernelHeight {
				inR := oR + kR
				for kC := range kernelWidth {
					inC := oC + kC

					z += input[iCh][inR][inC] * conv2dWorkerpool.weights[oCh][iCh][kR][kC]
				}
			}
		}
		(*output)[oCh][oR][oC] = z + conv2dWorkerpool.bias[oCh]
	}
}

func NewConv2dWorkerpool(moduleInfo modelarch.ModuleInfo, modulesParams modelarch.ModulesParams) conv2dWorkerpool {
	var name string
	raw, exists := moduleInfo.GetProp("name")
	if !exists {
		panic("Conv2d: name must be defined")
	} else {
		err := json.Unmarshal(raw, &name)
		if err != nil {
			panic(err)
		}
	}

	weightsRaw, exists := modulesParams[name+".weight"]
	if !exists {
		panic("Conv2d: weights must be defined")
	}
	var weights [][][][]float32
	err := json.Unmarshal(weightsRaw, &weights)
	if err != nil {
		panic("Conv2d: weights must be a four dimensional array")
	}

	outChans := len(weights)
	if outChans < 1 {
		panic("Conv2d: weights out channels has length 0")
	}
	weightInChans := len(weights[0])
	if weightInChans < 1 {
		panic("Conv2d: weights in channels has length 0")
	}
	kernelHeight := len(weights[0][0])
	if kernelHeight < 1 {
		panic("Conv2d: weights kernel height has length 0")
	}
	kernelWidth := len(weights[0][0][0])
	if kernelWidth < 1 {
		panic("Conv2d: weights kernel width has length 0")
	}

	for _, weightDim1 := range weights {
		if len(weightDim1) != weightInChans {
			panic("Conv2d: weight length must be consistent throughout dimension")
		}
		for _, weightDim2 := range weightDim1 {
			if len(weightDim2) != kernelHeight {
				panic("Conv2d: weight length must be consistent throughout dimension")
			}
			for _, weightDim3 := range weightDim2 {
				if len(weightDim3) != kernelWidth {
					panic("Conv2d: weight length must be consistent throughout dimension")
				}
			}
		}
	}

	var bias []float32
	biasRaw, exists := modulesParams[name+".bias"]
	if exists {
		err = json.Unmarshal(biasRaw, &bias)
		if err != nil {
			panic("Conv2d: bias must be a one dimensional array")
		}
	}
	if len(bias) != outChans {
		panic("Conv2d: bias size must match weight out_channels")
	}

	return conv2dWorkerpool{weights, bias}
}
