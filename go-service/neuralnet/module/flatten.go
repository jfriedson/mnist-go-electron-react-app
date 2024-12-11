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

func (self *flatten) Forward(input any) (any, error) {
	// assert input is 2D slice of float32 for the time being
	inputAssert, ok := input.([][]float32)
	if !ok {
		return nil, fmt.Errorf("for now, flatten input must be [][]float32")
	}

	output := make([]float32, len(inputAssert)*len(inputAssert[0]))
	for rowIdx, row := range inputAssert {
		rowMult := len(inputAssert) * rowIdx
		for colIdx, col := range row {
			output[rowMult+colIdx] = col
		}
	}

	return output, nil

	// // 0 dim / scalar
	// if reflect.TypeOf(input).Kind() != reflect.Slice {
	// 	if self.startDim > 0 || self.endDim > 0 {
	// 		return nil, fmt.Errorf("flatten input not of correct dimensionality")
	// 	}
	// 	return input, nil
	// }

	// output := []any{}

	// layer := input
	// for dim := 0; dim <= self.endDim || self.endDim == -1; dim++ {
	// 	var ok bool
	// 	layer, ok = layer.([]any)
	// 	if !ok {
	// 		if self.endDim == -1 {
	// 			break
	// 		} else {
	// 			return nil, fmt.Errorf("flatten input not of correct dimensionality")
	// 		}
	// 	}

	// }

	// return output, nil
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
