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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		output, ok := output.(uint)
		if !ok {
			panic("output is of invalid type")
		}

		fmt.Fprintf(w, fmt.Sprint(output))
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
	if len(inputBytes) != 28*28 {
		err := fmt.Errorf("invalid image size %d", len(inputBytes))
		return nil, err
	}

	input := make([][]float32, 28, 28)
	for idx, bodyByte := range inputBytes {
		row := idx / 28
		col := idx % 28

		if col == 0 {
			input[row] = make([]float32, 28, 28)
		}

		input[row][col] = float32(bodyByte) / 255.
	}

	return input, nil
}
