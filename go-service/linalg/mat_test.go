package linalg

import (
	"reflect"
	"slices"
	"testing"
)

func TestT(t *testing.T) {
	mat := Mat[int]{
		Data: [][]int{{1, 1, 1}, {2, 2, 2}},
		Dims: []int{2, 3},
	}

	matT := mat.T()

	expectedDims := []int{3, 2}
	if slices.Compare(matT.Dims, expectedDims) != 0 {
		t.Error("matT dims is incorrect")
	}
	expectedData := [][]int{{1, 2}, {1, 2}, {1, 2}}
	if !reflect.DeepEqual(matT.Data, expectedData) {
		t.Error("matT data is incorrect")
	}
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

	matAdd, err := mat1.Add(&mat2)
	if err != nil {
		t.Errorf("addition failed")
	}

	expectedAdd := Mat[int]{
		Data: [][]int{{2, 2, 2}, {4, 4, 4}, {6, 6, 6}},
		Dims: []int{3, 3},
	}
	if !reflect.DeepEqual(matAdd, expectedAdd) {
		t.Error("Mat addition returned invalid result")
	}
}
