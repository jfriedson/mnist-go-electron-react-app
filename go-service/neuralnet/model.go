package neuralnet

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/module"
)

type Model interface {
	Forward(*any) (any, error)
}

type model struct {
	modules []module.Module
}

func (self *model) Forward(input *any) (any, error) {
	// var next_input any = input
	// var output any

	return "21", nil
}

type ModelConfig struct {
	ArchFile  string
	ModelFile string
}

// TODO: build graph of modules at load time
func LoadModel(config ModelConfig) *model {
	arch := loadModelArch(config.ArchFile)

	modelBytes, err := os.ReadFile(config.ModelFile)
	if err != nil {
		panic(err)
	}

	var moduleParams map[string]any
	err = json.Unmarshal(modelBytes, &moduleParams)
	if err != nil {
		panic(err)
	}

	fmt.Println(moduleParams)

	model := buildModel(arch, moduleParams)

	return model
}

func buildModel(arch ModelArch, moduleParam map[string]any) *model {
	model := &model{
		modules: []module.Module{module.NewFlatten()},
	}

	return model
}
