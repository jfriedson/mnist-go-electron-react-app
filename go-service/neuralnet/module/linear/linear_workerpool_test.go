package linear

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestLinearWorkerpool_Forward(t *testing.T) {
	linearWorkerpool := &linearWorkerpool{
		weights: [][]float32{{1, 2, 3, 4}, {1, 2, 3, 4}},
		bias:    []float32{1, 2},
	}

	input := []float32{1, 2, 3, 4}
	output := linearWorkerpool.Forward(&input)

	outputSlice, ok := output.([]float32)
	if !ok {
		t.Fatal("failed to assert output type")
	}
	expectedOutput := []float32{31, 32}
	if slices.Compare(outputSlice, expectedOutput) != 0 {
		t.Fatal("output result does not match expectations")
	}
}

func BenchmarkLinearWorkerpool(b *testing.B) {
	weightDims := [2]int{256, 1024}
	weights := make([][]float32, weightDims[0])
	for o := range weightDims[0] {
		weights[o] = make([]float32, weightDims[1])
		for i := range weightDims[1] {
			weights[o][i] = rand.Float32()
		}
	}

	bias := make([]float32, weightDims[0])
	for o := range weightDims[0] {
		bias[o] = rand.Float32()
	}

	linearWorkerpool := &linearWorkerpool{
		weights,
		bias,
	}

	input := make([]float32, weightDims[1])
	for i := range weightDims[1] {
		input[i] = rand.Float32()
	}

	b.ResetTimer()
	for range b.N {
		linearWorkerpool.Forward(&input)
	}
}
