package conv2d

import (
	"math/rand/v2"
	"reflect"
	"testing"
)

func TestConv2dSequential_Forward(t *testing.T) {
	conv2dSequential := &conv2dSequential{
		weights: [][][][]float32{
			{
				{
					{1, 0, -1},
					{2, 0, -2},
					{1, 0, -1},
				},
				{
					{0, 1, 2},
					{0, 0, 0},
					{0, -1, -2},
				},
				{
					{1, 2, 1},
					{0, 0, 0},
					{-1, -2, -1},
				},
			},
		},
		bias: []float32{18},
	}

	input := [][][]float32{
		{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		},
		{
			{25, 24, 23, 22, 21},
			{20, 19, 18, 17, 16},
			{15, 14, 13, 12, 11},
			{10, 9, 8, 7, 6},
			{5, 4, 3, 2, 1},
		},
		{
			{1, 3, 5, 7, 9},
			{2, 4, 6, 8, 10},
			{11, 13, 15, 17, 19},
			{12, 14, 16, 18, 20},
			{21, 23, 25, 27, 29},
		},
	}

	output := conv2dSequential.Forward(&input)

	outputSlice, ok := output.([][][]float32)
	if !ok {
		t.Fatal("failed to assert output type")
	}

	expectedOutput := [][][]float32{
		{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	}
	if !reflect.DeepEqual(outputSlice, expectedOutput) {
		t.Fatal("output result does not match expectations")
	}
}

func BenchmarkConv2dSequential(b *testing.B) {
	weightDims := [4]int{32, 24, 3, 3}
	weights := make([][][][]float32, weightDims[0])
	for oCh := range weightDims[0] {
		weights[oCh] = make([][][]float32, weightDims[1])
		for iC := range weightDims[1] {
			weights[oCh][iC] = make([][]float32, weightDims[2])
			for oR := range weightDims[2] {
				weights[oCh][iC][oR] = make([]float32, weightDims[3])
				for oC := range weightDims[3] {
					weights[oCh][iC][oR][oC] = rand.Float32()
				}
			}
		}
	}

	bias := make([]float32, weightDims[0])
	for oCh := range weightDims[0] {
		bias[oCh] = rand.Float32()
	}

	conv2dSequential := &conv2dSequential{weights, bias}

	inputDims := [2]int{24, 24}
	input := make([][][]float32, weightDims[1])
	for iCh := range weightDims[1] {
		input[iCh] = make([][]float32, inputDims[0])
		for iR := range inputDims[0] {
			input[iCh][iR] = make([]float32, inputDims[1])
			for iC := range inputDims[1] {
				input[iCh][iR][iC] = rand.Float32()
			}
		}
	}

	b.ResetTimer()
	for range b.N {
		conv2dSequential.Forward(&input)
	}
}
