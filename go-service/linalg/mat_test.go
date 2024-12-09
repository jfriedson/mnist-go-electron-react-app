package linalg

import (
	"testing"
)

func TestT(t *testing.T) {
	mat := Mat[int]{
		Data: [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}},
		Dims: []int{3, 3},
	}

	matT := mat.T()
}

func TestAdd(t *testing.T) {
	mat1 := Mat[int]{
		Data: [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}},
		Dims: []int{3, 3},
	}

	mat2 := Mat[int]{
		Data: [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}},
		Dims: []int{3, 3},
	}

	matRes, err := mat1.Add(&mat2)
	if err != nil {
		t.Errorf("addition failed")
	}
}
