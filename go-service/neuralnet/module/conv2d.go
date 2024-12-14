package module

import (
	"encoding/json"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
)

type conv2d struct {
	weights [][][][]float32
	bias    []float32
}

// input is operated on in the format (channel, height, width)
// output is formatted (channel, height, width)
func (conv2d conv2d) Forward(inputPtr any) any {
	inputPtrTyped, ok := inputPtr.(*[][][]float32)
	if !ok {
		input2D, ok := inputPtr.(*[][]float32)
		if !ok {
			panic("Conv2d: input must be a non-nil pointer to a 2D or 3D float32")
		}
		inputExtended := make([][][]float32, 1)
		inputExtended[0] = *input2D
		inputPtrTyped = &inputExtended
	}

	input := *inputPtrTyped

	inChans := len(input)
	inHeight := len(input[0])
	inWidth := len(input[0][0])
	outChans := len(conv2d.weights)
	filters := len(conv2d.weights[0])
	kernelHeight := len(conv2d.weights[0][0])
	kernelWidth := len(conv2d.weights[0][0][0])

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

	// TODO: goroutine this puppy
	for oCh := range outChans {
		for iCh := range inChans {
			for oR := range outHeight {
				for oC := range outWidth {
					// apply kernel
					for kR := range kernelHeight {
						inR := oR + kR
						for kC := range kernelWidth {
							inC := oC + kC

							output[oCh][oR][oC] += input[iCh][inR][inC] * conv2d.weights[oCh][iCh][kR][kC]
						}
					}
				}
			}
		}
		for y := range outHeight {
			for x := range outWidth {
				output[oCh][y][x] += conv2d.bias[oCh]
			}
		}
	}

	return output
}

func NewConv2d(moduleInfo modelarch.ModuleInfo, modulesParams modelarch.ModulesParams) conv2d {
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

	return conv2d{weights, bias}
}
