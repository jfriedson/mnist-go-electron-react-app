package maxpool2d

import (
	"encoding/json"
	"reflect"
	"runtime"
	"sync"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
)

type maxpool2dWorkerpool struct {
	kernel_size int
}

func (maxpool2dWp maxpool2dWorkerpool) Forward(inputPtr any) any {
	inputPtrVal := reflect.ValueOf(inputPtr)
	if inputPtrVal.Kind() != reflect.Pointer || inputPtrVal.IsNil() {
		panic("MaxPool2d: input must be non-nil pointer to [][][]float32")
	}

	inputAny := inputPtrVal.Elem().Interface()

	input := inputAny.([][][]float32)

	outHeight := (len(input[0])-maxpool2dWp.kernel_size)/maxpool2dWp.kernel_size + 1
	outWidth := (len(input[0][0])-maxpool2dWp.kernel_size)/maxpool2dWp.kernel_size + 1
	chans := len(input)

	// initialize the output image
	output := make([][][]float32, chans)
	for c := range chans {
		output[c] = make([][]float32, outHeight)
		for oR := range outHeight {
			output[c][oR] = make([]float32, outWidth)
		}
	}

	numWorkers := runtime.NumCPU()
	jobs := make(chan maxpool2DWorkerpoolJob, chans*outHeight*outWidth)

	var wg sync.WaitGroup
	for range numWorkers {
		wg.Add(1)
		go maxpool2dWp.worker(jobs, &wg, input, &output)
	}

	for c := range chans {
		for oR := range outHeight {
			for oC := range outWidth {
				jobs <- maxpool2DWorkerpoolJob{c, oR, oC}
			}
		}
	}
	close(jobs)

	wg.Wait()

	return output
}

type maxpool2DWorkerpoolJob struct {
	c, oR, oC int
}

func (maxpool2dWp maxpool2dWorkerpool) worker(jobs <-chan maxpool2DWorkerpoolJob, wg *sync.WaitGroup,
	input [][][]float32, output *[][][]float32) {

	defer wg.Done()

	for j := range jobs {
		c := j.c
		oR := j.oR
		oC := j.oC

		maxVal := input[c][oR*maxpool2dWp.kernel_size][oC*maxpool2dWp.kernel_size]
		for kR := range maxpool2dWp.kernel_size {
			for kC := range maxpool2dWp.kernel_size {
				row := oR*maxpool2dWp.kernel_size + kR
				col := oC*maxpool2dWp.kernel_size + kC
				val := input[c][row][col]
				if val > maxVal {
					maxVal = val
				}
			}
		}
		(*output)[c][oR][oC] = maxVal
	}
}

func NewMaxPool2dWorkerpool(moduleInfo modelarch.ModuleInfo) maxpool2dWorkerpool {
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

	return maxpool2dWorkerpool{kernel_size}
}
