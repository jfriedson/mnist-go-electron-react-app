package main

import (
	"encoding/json"
	"os"
)

type ModuleInfo struct {
	ModuleType string         `json:"type"`
	Props      map[string]any `json:"props"`
}

type ModelArch []ModuleInfo

func loadMnistModelArch() ModelArch {
	modelBytes, err := os.ReadFile("../mnist-model-generator/models/mnist_test_arch.json")
	if err != nil {
		panic(err)
	}

	var modelArch ModelArch
	err = json.Unmarshal(modelBytes, &modelArch)
	if err != nil {
		panic(err)
	}

	return modelArch
}

type Model map[string][]any

func loadMnistModel() Model {
	modelBytes, err := os.ReadFile("../mnist-model-generator/models/mnist_test.json")
	if err != nil {
		panic(err)
	}

	var model Model
	err = json.Unmarshal(modelBytes, &model)
	if err != nil {
		panic(err)
	}

	return model
}
