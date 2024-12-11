package module

import (
	"slices"
	"testing"
)

// func TestForwardScalarInt(t *testing.T) {
// 	flatten := &flatten{
// 		startDim: 0,
// 		endDim:   -1,
// 	}

// 	input := 123
// 	output, err := flatten.Forward(input)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}
// 	if output != input {
// 		t.Error("input must match output")
// 	}
// }

// func TestForward1DimInt(t *testing.T) {
// 	flatten := &flatten{
// 		startDim: 0,
// 		endDim:   -1,
// 	}

// 	input := []int{1, 2, 3}
// 	output, err := flatten.Forward(input)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}
// 	outputSlice, ok := output.([]int)
// 	if !ok {
// 		t.Error("failed to assert output type")
// 	}
// 	if slices.Compare(outputSlice, input) != 0 {
// 		fmt.Println(outputSlice)
// 		fmt.Println(input)
// 		t.Error("input must match output")
// 	}
// }

func TestFlatten_Forward2DimFloat32(t *testing.T) {
	flatten := &flatten{
		startDim: 0,
		endDim:   -1,
	}

	input := [][]float32{{1, 2}, {3, 4}}
	output, err := flatten.Forward(input)
	if err != nil {
		t.Error(err)
	}
	outputSlice, ok := output.([]float32)
	if !ok {
		t.Error("failed to assert output type")
	}
	expectedOutput := []float32{1, 2, 3, 4}
	if slices.Compare(outputSlice, expectedOutput) != 0 {
		t.Error("input must match output")
	}
}
