package logsoftmax

import (
	"slices"
	"testing"
)

func TestLogSoftmax_Forward1Dim(t *testing.T) {
	logsoftmax := &logsoftmax{
		dim: 1,
	}

	input := []float32{1, 2, 3, 4}
	output := logsoftmax.Forward(&input)

	outputSlice, ok := output.([]float32)
	if !ok {
		t.Fatal("failed to assert output type")
	}
	expectedOutput := []float32{-3.4401896, -2.4401896, -1.4401897, -0.4401897}
	if slices.Compare(outputSlice, expectedOutput) != 0 {
		t.Fatal("output result does not match expectations")
	}
}

func BenchmarkLogSoftmax(b *testing.B) {
	logsoftmax := &logsoftmax{
		dim: 1,
	}

	input := []float32{1, 2, 3, 4}

	b.ResetTimer()
	for range b.N {
		logsoftmax.Forward(&input)
	}
}
