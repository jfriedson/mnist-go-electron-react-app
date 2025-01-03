package flatten

import (
	"reflect"
	"slices"
	"testing"
)

func TestFlatten_ForwardScalar(t *testing.T) {
	flatten := &flatten{
		startDim: 0,
		endDim:   -1,
	}

	var input float32 = 123
	output := flatten.Forward(&input)

	if output != input {
		t.Fatal("input must match output")
	}
}

func TestFlatten_ForwardDim1(t *testing.T) {
	flatten := &flatten{
		startDim: 0,
		endDim:   -1,
	}

	input := []float32{1, 2, 3}
	output := flatten.Forward(&input)

	outputSlice, ok := output.([]float32)
	if !ok {
		t.Fatal("failed to assert output type")
	}
	if slices.Compare(outputSlice, input) != 0 {
		t.Fatal("input must match output")
	}
}

func TestFlatten_ForwardDim2(t *testing.T) {
	flatten := &flatten{
		startDim: 1,
		endDim:   1,
	}

	input := [][]float32{{1, 2}, {3, 4}}
	output := flatten.Forward(&input)

	outputSlice, ok := output.([]float32)
	if !ok {
		t.Fatal("failed to assert output type")
	}
	expectedOutput := []float32{1, 2, 3, 4}
	if slices.Compare(outputSlice, expectedOutput) != 0 {
		t.Fatal("output result does not match expectations")
	}
}

func TestFlatten_ForwardDim3(t *testing.T) {
	flatten := &flatten{
		startDim: 0,
		endDim:   -1,
	}

	input := [][][]float32{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	output := flatten.Forward(&input)

	outputSlice, ok := output.([]float32)
	if !ok {
		t.Fatal("failed to assert output type")
	}
	expectedOutput := []float32{1, 2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(outputSlice, expectedOutput) {
		t.Fatal("output result does not match expectations")
	}
}

func BenchmarkFlatten_ForwardDim3(b *testing.B) {
	flatten := &flatten{}
	input := [][][]float32{{{-4, -3}, {-2, -1}, {0, 1}, {2, 3}}}

	b.ResetTimer()
	for range b.N {
		flatten.Forward(&input)
	}
}
