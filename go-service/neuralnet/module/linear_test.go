package module

import (
	"slices"
	"testing"
)

func TestLinear_Forward1DimFloat32(t *testing.T) {
	linear := &linear{
		weights: [][]float32{{1, 2, 3, 4}, {1, 2, 3, 4}},
		bias:    []float32{1, 2},
	}

	input := []float32{1, 2, 3, 4}
	output, err := linear.Forward(input)
	if err != nil {
		t.Error(err)
	}
	outputSlice, ok := output.([]float32)
	if !ok {
		t.Error("failed to assert output type")
	}
	expectedOutput := []float32{31, 32}
	if slices.Compare(outputSlice, expectedOutput) != 0 {
		t.Error("output result does not match expectations")
	}
}
