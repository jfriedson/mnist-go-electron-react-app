package linear

import (
	"math/rand/v2"
	"runtime"
	"slices"
	"sync"
	"testing"
)

func TestLinear_Forward(t *testing.T) {
	numWorkers := runtime.NumCPU()
	jobs := make(chan linearStaticWpJob, numWorkers)
	var wg sync.WaitGroup

	linearStaticWp := &linearStaticWp{
		weights: [][]float32{{1, 2, 3, 4}, {1, 2, 3, 4}},
		bias:    []float32{1, 2},
		jobs:    jobs,
		wg:      &wg,
	}

	for range numWorkers {
		go linearStaticWp.worker()
	}

	input := []float32{1, 2, 3, 4}
	output := linearStaticWp.Forward(&input)
	close(jobs)

	outputSlice, ok := output.([]float32)
	if !ok {
		t.Fatal("failed to assert output type")
	}
	expectedOutput := []float32{31, 32}
	if slices.Compare(outputSlice, expectedOutput) != 0 {
		t.Fatal("output result does not match expectations")
	}
}

func BenchmarkLinearStaticWp(b *testing.B) {
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

	input := make([]float32, weightDims[1])
	for i := range weightDims[1] {
		input[i] = rand.Float32()
	}

	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup

	jobs := make(chan linearStaticWpJob, weightDims[0])

	linearStaticWp := linearStaticWp{weights, bias, jobs, &wg}
	for range numWorkers {
		go linearStaticWp.worker()
	}

	b.ResetTimer()
	for range b.N {
		linearStaticWp.Forward(&input)
	}

	close(jobs)
}
