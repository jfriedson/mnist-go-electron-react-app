package main

import (
	"fmt"
	"net/http"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet"
)

func main() {
	modelConfig := neuralnet.ModelConfig{
		ArchFile:  "../mnist-model-generator/models/mnist_test_arch.json",
		ModelFile: "../mnist-model-generator/models/mnist_test.json",
	}
	model := neuralnet.LoadModel(modelConfig)

	inferenceEndpoint := InferenceEndpoint{model}

	http.HandleFunc("/", cors(inferenceEndpoint.handler))

	fmt.Println("go-service listening on http://localhost:5122/")
	http.ListenAndServe(":5122", nil)
}
