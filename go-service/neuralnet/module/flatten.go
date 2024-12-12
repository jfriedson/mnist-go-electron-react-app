package module

import (
	"encoding/json"
	"fmt"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
)

type flatten struct {
	startDim int
	endDim   int
}

func (self *flatten) Forward(inputAny any) (any, error) {
	// assert input is 2D slice of float32 for the time being
	input, ok := inputAny.([][]float32)
	if !ok {
		return nil, fmt.Errorf("for now, flatten input must be [][]float32")
	}

	lengthDim0 := len(input)
	lengthDim1 := len(input[0])
	output := make([]float32, lengthDim0*lengthDim1)
	for d0i := range lengthDim0 {
		d0offset := lengthDim0 * d0i
		for d1i := range lengthDim1 {
			output[d0offset+d1i] = input[d0i][d1i]
		}
	}

	return output, nil
}

func NewFlatten(moduleInfo modelarch.ModuleInfo) *flatten {
	var startDim, endDim int

	raw, exists := moduleInfo.GetProp("start_dim")
	if !exists {
		startDim = 0
	} else {
		err := json.Unmarshal(raw, &startDim)
		if err != nil {
			panic("start_dim must be a number")
		}
	}

	raw, exists = moduleInfo.GetProp("end_dim")
	if !exists {
		endDim = -1
	} else {
		err := json.Unmarshal(raw, &endDim)
		if err != nil {
			panic("end_dim must be a number")
		}
	}

	if startDim < 0 {
		panic("start_dim must be 0 or greater")
	}

	if endDim < -1 {
		panic("start_dim must be -1 or greater")
	}

	if endDim > -1 && startDim <= endDim {
		panic("end_dim must be equal to or greater than start_dim if end_dim is not -1")
	}

	return &flatten{startDim, endDim}
}
