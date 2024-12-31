package relu

import (
	"reflect"
	"slices"
	"testing"
)

func TestReluSequential_ForwardScalar(t *testing.T) {
	reluSequential := &reluSequential{}

	var input float32 = -1
	output := reluSequential.Forward(&input)
	if output != nil {
		t.Fatal("ReLU: output is expected to be nil")
	}

	var expectedOutput float32 = 0
	if input != expectedOutput {
		t.Fatal("output result does not match expectations")
	}
}

func TestReluSequential_Forward1Dim(t *testing.T) {
	reluSequential := &reluSequential{}

	input := []float32{-1, 0, 1}
	output := reluSequential.Forward(&input)
	if output != nil {
		t.Fatal("ReLU: output is expected to be nil")
	}

	expectedOutput := []float32{0, 0, 1}
	if slices.Compare(input, expectedOutput) != 0 {
		t.Fatal("output result does not match expectations")
	}
}

func TestReluSequential_Forward2Dim(t *testing.T) {
	reluSequential := &reluSequential{}

	input := [][]float32{{-4, -3, -2}, {-1, 0, 1}, {2, 3, 4}}
	output := reluSequential.Forward(&input)
	if output != nil {
		t.Fatal("ReLU: output is expected to be nil")
	}

	expectedOutput := [][]float32{{0, 0, 0}, {0, 0, 1}, {2, 3, 4}}
	if !reflect.DeepEqual(input, expectedOutput) {
		t.Fatal("output result does not match expectations")
	}
}

func TestReluSequential_Forward3Dim(t *testing.T) {
	reluSequential := &reluSequential{}

	input := [][][]float32{{{-4, -3}, {-2, -1}}, {{0, 1}, {2, 3}}}
	output := reluSequential.Forward(&input)
	if output != nil {
		t.Fatal("ReLU: output is expected to be nil")
	}

	expectedOutput := [][][]float32{{{0, 0}, {0, 0}}, {{0, 1}, {2, 3}}}
	if !reflect.DeepEqual(input, expectedOutput) {
		t.Fatal("output result does not match expectations")
	}
}

func BenchmarkReluSequential_ForwardDim3(b *testing.B) {
	reluSequential := &reluSequential{}
	input := [][][]float32{{{-4, -3}, {-2, -1}, {0, 1}, {2, 3}}}

	for range b.N {
		reluSequential.Forward(&input)
	}
}
