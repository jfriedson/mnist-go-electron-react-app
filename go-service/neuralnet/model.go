package neuralnet

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/module"
)

type Model interface {
	Forward(any) (any, error)
}

type model struct {
	modules []module.Module
}

func (self *model) Forward(input any) (any, error) {
	var output any

	for _, module := range self.modules {
		var err error
		output, err = module.Forward(input)
		if err != nil {
			return nil, err
		}
		input = output
	}

	return output, nil
}

type ModelConfig struct {
	ArchFile  string
	ModelFile string
}

func LoadModel(config ModelConfig) *model {
	arch := modelarch.LoadModelArch(config.ArchFile)

	modelBytes, err := os.ReadFile(config.ModelFile)
	if err != nil {
		panic(err)
	}

	var modulesParams modelarch.ModulesParams
	err = json.Unmarshal(modelBytes, &modulesParams)
	if err != nil {
		panic(err)
	}

	model := buildModel(arch, modulesParams)

	return model
}

func buildModel(arch modelarch.ModelArch, modulesParams modelarch.ModulesParams) *model {
	modules := []module.Module{}

	for moduleInfos := range arch.GetModuleInfos() {
		switch moduleInfos.GetType() {
		case "flatten":
		case "Flatten":
			modules = append(modules, module.NewFlatten(moduleInfos))
		case "Linear":
			modules = append(modules, module.NewLinear(moduleInfos, modulesParams))
		case "LogSoftmax":
			// not required for inference, but implemented anyways :)
			// modules = append(modules, module.NewLogSoftmax(moduleInfos))
		case "MaxPool2d":
			modules = append(modules, module.NewMaxPool2d(moduleInfos))
		case "ReLU":
			modules = append(modules, module.NewReLU())
		default:
			panic(fmt.Sprintf("unrecognized module type: %s", moduleInfos.GetType()))
		}
	}

	return &model{modules}
}
