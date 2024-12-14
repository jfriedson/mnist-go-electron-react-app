package module

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestReLU_ForwardScalar(t *testing.T) {
	relu := &relu{}

	var input float32 = -1
	output := relu.Forward(&input)
	if output != nil {
		t.Fatal("ReLU: output is expected to be nil")
	}

	var expectedOutput float32 = 0
	if input != expectedOutput {
		fmt.Println(input)
		t.Fatal("output result does not match expectations")
	}
}

func TestReLU_Forward1Dim(t *testing.T) {
	relu := &relu{}

	input := []float32{-1, 0, 1}
	output := relu.Forward(&input)
	if output != nil {
		t.Fatal("ReLU: output is expected to be nil")
	}

	expectedOutput := []float32{0, 0, 1}
	if slices.Compare(input, expectedOutput) != 0 {
		fmt.Println(input)
		t.Fatal("output result does not match expectations")
	}
}

func TestReLU_Forward2Dim(t *testing.T) {
	relu := &relu{}

	input := [][]float32{{-4, -3, -2}, {-1, 0, 1}, {2, 3, 4}}
	output := relu.Forward(&input)
	if output != nil {
		t.Fatal("ReLU: output is expected to be nil")
	}

	expectedOutput := [][]float32{{0, 0, 0}, {0, 0, 1}, {2, 3, 4}}
	if !reflect.DeepEqual(input, expectedOutput) {
		t.Fatal("output result does not match expectations")
	}
}

func TestReLU_Forward3Dim(t *testing.T) {
	relu := &relu{}

	input := [][][]float32{{{-4, -3}, {-2, -1}}, {{0, 1}, {2, 3}}}
	output := relu.Forward(&input)
	if output != nil {
		t.Fatal("ReLU: output is expected to be nil")
	}

	expectedOutput := [][][]float32{{{0, 0}, {0, 0}}, {{0, 1}, {2, 3}}}
	if !reflect.DeepEqual(input, expectedOutput) {
		t.Fatal("output result does not match expectations")
	}
}

func BenchmarkReLU_ForwardDim1(b *testing.B) {
	relu := &relu{}
	var input float32 = 5

	for range b.N {
		relu.Forward(&input)
	}
}

func BenchmarkReLU_ForwardDim2(b *testing.B) {
	relu := &relu{}
	input := [][]int{{-4, -3, -2}, {-1, 0, 1}, {2, 3, 4}}

	for range b.N {
		relu.Forward(&input)
	}
}
