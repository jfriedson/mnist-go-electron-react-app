package maxpool2d

import (
	"reflect"
	"testing"
)

func TestMaxpool2dSequential_Forward(t *testing.T) {
	maxpool2dSequential := &maxpool2dSequential{
		kernel_size: 2,
	}

	input := [][][]float32{
		{
			{1, 3, 2, 4, 6, 8},
			{5, 6, 7, 8, 1, 3},
			{1, 2, 3, 1, 2, 3},
			{4, 5, 6, 7, 8, 9},
			{10, 11, 12, 13, 14, 15},
			{3, 4, 5, 6, 7, 8},
		},
		{
			{8, 5, 3, 7, 4, 1},
			{9, 6, 4, 8, 2, 5},
			{7, 2, 9, 1, 3, 6},
			{6, 5, 8, 3, 9, 7},
			{5, 1, 4, 6, 8, 2},
			{3, 8, 7, 2, 5, 9},
		},
	}

	output := maxpool2dSequential.Forward(&input)

	outputSlice, ok := output.([][][]float32)
	if !ok {
		t.Fatal("failed to assert output type")
	}

	expectedOutput := [][][]float32{
		{
			{6, 8, 8},
			{5, 7, 9},
			{11, 13, 15},
		},
		{
			{9, 8, 5},
			{7, 9, 9},
			{8, 7, 9},
		},
	}
	if !reflect.DeepEqual(outputSlice, expectedOutput) {
		t.Fatal("output result does not match expectations")
	}
}

func BenchmarkMaxpool2dSequential(b *testing.B) {
	maxpool2dSequential := &maxpool2dSequential{
		kernel_size: 2,
	}

	input := [][][]float32{
		{
			{1, 3, 2, 4, 6, 8},
			{5, 6, 7, 8, 1, 3},
			{1, 2, 3, 1, 2, 3},
			{4, 5, 6, 7, 8, 9},
			{10, 11, 12, 13, 14, 15},
			{3, 4, 5, 6, 7, 8},
		},
		{
			{8, 5, 3, 7, 4, 1},
			{9, 6, 4, 8, 2, 5},
			{7, 2, 9, 1, 3, 6},
			{6, 5, 8, 3, 9, 7},
			{5, 1, 4, 6, 8, 2},
			{3, 8, 7, 2, 5, 9},
		},
	}

	for range b.N {
		maxpool2dSequential.Forward(&input)
	}
}
