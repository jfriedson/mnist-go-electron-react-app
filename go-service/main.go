package main

import (
	"fmt"
	"net/http"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet"
)

func main() {
	modelConfig := neuralnet.ModelConfig{
		ArchFile:  "../mnist-model-generator/models/mnist_arch.json",
		ModelFile: "../mnist-model-generator/models/mnist.json",
	}
	model := neuralnet.LoadModel(modelConfig)

	inferenceEndpoint := InferenceEndpoint{model}

	http.HandleFunc("/", cors(inferenceEndpoint.handler))

	fmt.Println("go-service listening on http://localhost:5122/")
	err := http.ListenAndServe(":5122", nil)
	if err != nil {
		panic("error starting http server")
	}
}
