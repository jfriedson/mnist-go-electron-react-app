package maxpool2d

import (
	"encoding/json"
	"reflect"
	"sync"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
)

type maxpool2dGoroutine struct {
	kernel_size int
}

func (maxpool2dGr maxpool2dGoroutine) Forward(inputPtr any) any {
	inputPtrVal := reflect.ValueOf(inputPtr)
	if inputPtrVal.Kind() != reflect.Pointer || inputPtrVal.IsNil() {
		panic("MaxPool2d: input must be non-nil pointer to [][][]float32")
	}

	inputAny := inputPtrVal.Elem().Interface()

	input := inputAny.([][][]float32)

	outHeight := (len(input[0])-maxpool2dGr.kernel_size)/maxpool2dGr.kernel_size + 1
	outWidth := (len(input[0][0])-maxpool2dGr.kernel_size)/maxpool2dGr.kernel_size + 1
	chans := len(input)

	// initialize the output image
	output := make([][][]float32, chans)
	for c := range chans {
		output[c] = make([][]float32, outHeight)
		for oR := range outHeight {
			output[c][oR] = make([]float32, outWidth)
		}
	}

	var wg sync.WaitGroup
	for c := range chans {
		for oR := range outHeight {
			for oC := range outWidth {
				wg.Add(1)
				go maxpool2dGr.goroutine(&wg, input, &output, c, oR, oC)
			}
		}
	}

	return output
}

func (maxpool2dGr maxpool2dGoroutine) goroutine(wg *sync.WaitGroup,
	input [][][]float32, output *[][][]float32,
	c, oR, oC int) {

	defer wg.Done()

	maxVal := input[c][oR*maxpool2dGr.kernel_size][oC*maxpool2dGr.kernel_size]
	for kR := range maxpool2dGr.kernel_size {
		for kC := range maxpool2dGr.kernel_size {
			row := oR*maxpool2dGr.kernel_size + kR
			col := oC*maxpool2dGr.kernel_size + kC
			val := input[c][row][col]
			if val > maxVal {
				maxVal = val
			}
		}
	}
	(*output)[c][oR][oC] = maxVal
}

func NewMaxPool2dGoroutine(moduleInfo modelarch.ModuleInfo) maxpool2dGoroutine {
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

	return maxpool2dGoroutine{kernel_size}
}
