package main

import (
	"reflect"
	"testing"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet"
)

func TestInferenceEndpoint_ConvertInput(t *testing.T) {
	modelConfig := neuralnet.ModelConfig{
		ArchFile:  "../mnist-model-generator/models/mnist_arch.json",
		ModelFile: "../mnist-model-generator/models/mnist.json",
	}
	model := neuralnet.LoadModel(modelConfig)
	inferenceEndpoint := InferenceEndpoint{model}

	const imgDim int = 3
	input := []byte{
		0, 1, 2,
		0, 51, 255,
		253, 254, 255,
	}

	output, err := inferenceEndpoint.convertInput(input, imgDim)
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := [][][]float32{
		{
			{0, 0.003921569, 0.007843138},
			{0, 0.2, 1},
			{0.99215686, 0.99607843, 1},
		},
	}

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatal("output result does not match expectations")
	}
}
