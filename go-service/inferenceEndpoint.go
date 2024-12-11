package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet"
)

type InferenceEndpoint struct {
	model neuralnet.Model
}

func (self InferenceEndpoint) handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		w.Header().Add("Access-Control-Allow-Methods", "OPTIONS, POST")
	case http.MethodPost:
		bodyBytes, err := self.parseRequestBody(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		input, err := self.convertInput(bodyBytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		output, err := self.model.Forward(input)
		if err != nil {
			panic(err)
		}

		outputAssert, ok := output.([]float32)
		if !ok {
			panic("output is an invalid type")
		}

		var maxIdx int = 0
		var maxVal float32 = outputAssert[0]
		for i, v := range outputAssert {
			if v > maxVal {
				maxIdx = i
				maxVal = v
			}
		}

		fmt.Fprint(w, maxIdx)
	default:
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
	}
}

func (self InferenceEndpoint) parseRequestBody(reqBody io.ReadCloser) ([]byte, error) {
	bodyBytes, err := io.ReadAll(reqBody)
	if err != nil {
		return nil, err
	}

	err = reqBody.Close()

	return bodyBytes, err
}

func (self InferenceEndpoint) convertInput(inputBytes []byte) ([][]float32, error) {
	const imgDim = 28

	if len(inputBytes) != imgDim*imgDim {
		err := fmt.Errorf("invalid image size %d", len(inputBytes))
		return nil, err
	}

	input := make([][]float32, imgDim, imgDim)
	for idx, bodyByte := range inputBytes {
		row := idx / imgDim
		col := idx % imgDim

		if col == 0 {
			input[row] = make([]float32, imgDim, imgDim)
		}

		input[row][col] = float32(bodyByte) / 255.
	}

	return input, nil
}
