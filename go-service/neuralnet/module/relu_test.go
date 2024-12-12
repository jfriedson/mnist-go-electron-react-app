package module

import (
	"fmt"
	"slices"
	"testing"
)

func TestRelu_Forward1DimFloat32(t *testing.T) {
	relu := &relu{}

	input := []float32{-1, 0, 1}
	output, err := relu.Forward(input)
	if err != nil {
		t.Error(err)
	}
	outputSlice, ok := output.([]float32)
	if !ok {
		t.Error("failed to assert output type")
	}
	expectedOutput := []float32{0, 0, 1}
	if slices.Compare(outputSlice, expectedOutput) != 0 {
		fmt.Print(outputSlice)
		t.Error("output result does not match expectations")
	}
}
