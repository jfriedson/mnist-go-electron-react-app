package module

import (
	"slices"
	"testing"
)

func TestLinear_Forward(t *testing.T) {
	linear := &linear{
		weights: [][]float32{{1, 2, 3, 4}, {1, 2, 3, 4}},
		bias:    []float32{1, 2},
	}

	input := []float32{1, 2, 3, 4}
	output := linear.Forward(&input)

	outputSlice, ok := output.([]float32)
	if !ok {
		t.Fatal("failed to assert output type")
	}
	expectedOutput := []float32{31, 32}
	if slices.Compare(outputSlice, expectedOutput) != 0 {
		t.Fatal("output result does not match expectations")
	}
}

func BenchmarkLinear(b *testing.B) {
	linear := &linear{
		weights: [][]float32{{1, 2, 3, 4}, {1, 2, 3, 4}},
		bias:    []float32{1, 2},
	}

	input := []float32{1, 2, 3, 4}

	for range b.N {
		linear.Forward(&input)
	}
}
