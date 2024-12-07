package main

import (
	"fmt"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/linalg/mat"
)

func main() {
	m1 := mat.Mat[int]{
		Data: [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}},
		Rows: 3,
		Cols: 3,
	}

	fmt.Println(m1)
}
