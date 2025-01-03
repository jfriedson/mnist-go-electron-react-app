package relu

import (
	"reflect"
	"slices"
	"testing"
)

func TestReluGoroutine_ForwardScalar(t *testing.T) {
	reluGoroutine := &reluGoroutine{}

	var input float32 = -1
	output := reluGoroutine.Forward(&input)
	if output != nil {
		t.Fatal("ReLU: output is expected to be nil")
	}

	var expectedOutput float32 = 0
	if input != expectedOutput {
		t.Fatal("output result does not match expectations")
	}
}

func TestReluGoroutine_Forward1Dim(t *testing.T) {
	reluGoroutine := &reluGoroutine{}

	input := []float32{-1, 0, 1}
	output := reluGoroutine.Forward(&input)
	if output != nil {
		t.Fatal("ReLU: output is expected to be nil")
	}

	expectedOutput := []float32{0, 0, 1}
	if slices.Compare(input, expectedOutput) != 0 {
		t.Fatal("output result does not match expectations")
	}
}

func TestReluGoroutine_Forward2Dim(t *testing.T) {
	reluGoroutine := &reluGoroutine{}

	input := [][]float32{{-4, -3, -2}, {-1, 0, 1}, {2, 3, 4}}
	output := reluGoroutine.Forward(&input)
	if output != nil {
		t.Fatal("ReLU: output is expected to be nil")
	}

	expectedOutput := [][]float32{{0, 0, 0}, {0, 0, 1}, {2, 3, 4}}
	if !reflect.DeepEqual(input, expectedOutput) {
		t.Fatal("output result does not match expectations")
	}
}

func TestReluGoroutine_Forward3Dim(t *testing.T) {
	reluGoroutine := &reluGoroutine{}

	input := [][][]float32{{{-4, -3}, {-2, -1}}, {{0, 1}, {2, 3}}}
	output := reluGoroutine.Forward(&input)
	if output != nil {
		t.Fatal("ReLU: output is expected to be nil")
	}

	expectedOutput := [][][]float32{{{0, 0}, {0, 0}}, {{0, 1}, {2, 3}}}
	if !reflect.DeepEqual(input, expectedOutput) {
		t.Fatal("output result does not match expectations")
	}
}

func BenchmarkReluGoroutine_ForwardDim3(b *testing.B) {
	reluGoroutine := &reluGoroutine{}
	input := [][][]float32{{{-4, -3}, {-2, -1}, {0, 1}, {2, 3}}}

	b.ResetTimer()
	for range b.N {
		reluGoroutine.Forward(&input)
	}
}
