package module

// func TestLogSoftmax_Forward(t *testing.T) {
// 	logsoftmax := &logsoftmax{
// 		dim: 1,
// 	}

// 	input := [][][]float32{{1}, 2, 3, 4}
// 	output, err := logsoftmax.Forward(input)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	outputSlice, ok := output.([]float32)
// 	if !ok {
// 		t.Fatal("failed to assert output type")
// 	}
// 	expectedOutput := []float32{-3.4401896, -2.4401896, -1.4401897, -0.4401897}
// 	if slices.Compare(outputSlice, expectedOutput) != 0 {
// 		fmt.Print(outputSlice)
// 		t.Fatal("output result does not match expectations")
// 	}
// }
