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

		const imgDim int = 28
		input, err := self.convertInput(bodyBytes, imgDim)
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

func (self InferenceEndpoint) convertInput(input []byte, imgDim int) ([][]float32, error) {
	if len(input) != imgDim*imgDim {
		err := fmt.Errorf("invalid image size %d", len(input))
		return nil, err
	}

	output := make([][]float32, imgDim)

	for y := 0; y < imgDim; y++ {
		colAdj := y * imgDim
		newRow := make([]float32, imgDim)
		for x := 0; x < imgDim; x++ {
			newRow[x] = float32(input[colAdj+x]) / 255.
		}
		output[y] = newRow
	}

	return output, nil
}
