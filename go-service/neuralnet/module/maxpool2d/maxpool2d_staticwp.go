package maxpool2d

import (
	"encoding/json"
	"reflect"
	"runtime"
	"sync"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
)

type maxpool2dStaticWp struct {
	kernel_size int
	jobs        chan maxpool2dStaticWpJob
	wg          *sync.WaitGroup
}

func (maxpool2dStaticWp maxpool2dStaticWp) Forward(inputPtr any) any {
	inputPtrVal := reflect.ValueOf(inputPtr)
	if inputPtrVal.Kind() != reflect.Pointer || inputPtrVal.IsNil() {
		panic("MaxPool2d: input must be non-nil pointer to [][][]float32")
	}

	inputAny := inputPtrVal.Elem().Interface()

	input := inputAny.([][][]float32)

	outHeight := (len(input[0])-maxpool2dStaticWp.kernel_size)/maxpool2dStaticWp.kernel_size + 1
	outWidth := (len(input[0][0])-maxpool2dStaticWp.kernel_size)/maxpool2dStaticWp.kernel_size + 1
	chans := len(input)

	// initialize the output image
	output := make([][][]float32, chans)
	for c := range chans {
		output[c] = make([][]float32, outHeight)
		for oR := range outHeight {
			output[c][oR] = make([]float32, outWidth)
		}
	}

	for c := range chans {
		for oR := range outHeight {
			for oC := range outWidth {
				maxpool2dStaticWp.wg.Add(1)
				maxpool2dStaticWp.jobs <- maxpool2dStaticWpJob{input, &output, c, oR, oC}
			}
		}
	}
	maxpool2dStaticWp.wg.Wait()

	return output
}

type maxpool2dStaticWpJob struct {
	input     [][][]float32
	output    *[][][]float32
	c, oR, oC int
}

func (maxpool2dStaticWp maxpool2dStaticWp) worker() {
	kernel_size := maxpool2dStaticWp.kernel_size

	for j := range maxpool2dStaticWp.jobs {
		input := j.input
		output := j.output
		c := j.c
		oR := j.oR
		oC := j.oC

		maxVal := input[c][oR*kernel_size][oC*kernel_size]
		for kR := range kernel_size {
			for kC := range kernel_size {
				row := oR*kernel_size + kR
				col := oC*kernel_size + kC
				val := input[c][row][col]
				if val > maxVal {
					maxVal = val
				}
			}
		}
		(*output)[c][oR][oC] = maxVal

		maxpool2dStaticWp.wg.Done()
	}
}

func NewMaxPool2dStaticWp(moduleInfo modelarch.ModuleInfo) maxpool2dStaticWp {
	var kernel_size int

	raw, exists := moduleInfo.GetProp("kernel_size")
	if !exists {
		panic("MaxPool2d: kernel_size must be defined")
	} else {
		err := json.Unmarshal(raw, &kernel_size)
		if err != nil {
			panic("MaxPool2d: kernel_size must be a number")
		}
	}

	if kernel_size < 1 {
		panic("MaxPool2d: kernel_size must be 1 or greater")
	}

	jobs := make(chan maxpool2dStaticWpJob)
	var wg sync.WaitGroup

	maxpool2dStaticWp := maxpool2dStaticWp{kernel_size, jobs, &wg}

	numWorkers := runtime.NumCPU()
	for range numWorkers {
		go maxpool2dStaticWp.worker()
	}

	return maxpool2dStaticWp
}
