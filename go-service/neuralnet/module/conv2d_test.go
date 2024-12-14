package module

import (
	"reflect"
	"testing"
)

func TestConv2d_Forward2Dim(t *testing.T) {
	conv2d := &conv2d{
		weights: [][][][]float32{
			{
				{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
		},
		bias: []float32{1},
	}

	input := [][]float32{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	output := conv2d.Forward(&input)

	outputSlice, ok := output.([][][]float32)
	if !ok {
		t.Fatal("failed to assert output type")
	}

	expectedOutput := [][][]float32{
		{
			{8, 9, 10},
			{13, 14, 15},
			{18, 19, 20},
		},
	}
	if !reflect.DeepEqual(outputSlice, expectedOutput) {
		t.Fatal("output result does not match expectations")
	}
}

func TestConv2d_Forward3Dim(t *testing.T) {
	conv2d := &conv2d{
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

	output := conv2d.Forward(&input)

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

func BenchmarkConv2d2D(b *testing.B) {
	conv2d := &conv2d{
		weights: [][][][]float32{
			{
				{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
		},
		bias: []float32{1},
	}

	input := [][]float32{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	for range b.N {
		conv2d.Forward(&input)
	}
}

func BenchmarkConv2d3D(b *testing.B) {
	conv2d := &conv2d{
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
		bias: []float32{5},
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

	for range b.N {
		conv2d.Forward(&input)
	}
}
