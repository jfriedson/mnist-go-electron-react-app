package maxpool2d

import (
	"reflect"
	"runtime"
	"sync"
	"testing"
)

func TestMaxpool2dStaticWp_Forward(t *testing.T) {
	numWorkers := runtime.NumCPU()
	jobs := make(chan maxpool2dStaticWpJob, numWorkers)
	var wg sync.WaitGroup

	maxpool2dStaticWp := &maxpool2dStaticWp{
		kernel_size: 2,
		jobs:        jobs,
		wg:          &wg,
	}

	for range numWorkers {
		go maxpool2dStaticWp.worker()
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

	output := maxpool2dStaticWp.Forward(&input)
	close(jobs)

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

func BenchmarkMaxpool2dStaticWp(b *testing.B) {
	numWorkers := runtime.NumCPU()
	jobs := make(chan maxpool2dStaticWpJob, numWorkers)
	var wg sync.WaitGroup

	maxpool2dStaticWp := &maxpool2dStaticWp{
		kernel_size: 2,
		jobs:        jobs,
		wg:          &wg,
	}

	for range numWorkers {
		go maxpool2dStaticWp.worker()
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

	b.ResetTimer()
	for range b.N {
		maxpool2dStaticWp.Forward(&input)
	}

	close(jobs)
}
