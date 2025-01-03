package neuralnet

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/module"
	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/module/conv2d"
	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/module/flatten"
	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/module/linear"
	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/module/logsoftmax"
	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/module/maxpool2d"
	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/module/relu"
)

type Model interface {
	Forward(any) any
}

type model struct {
	modules []module.Module
}

func (model model) Forward(input any) any {
	var output any

	for _, module := range model.modules {
		output = module.Forward(&input)

		// inplace modifiers have nil output
		if output != nil {
			input = output
		}
	}

	if output == nil {
		output = input
	}
	return output
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
		case "Conv2d":
			modules = append(modules, conv2d.NewConv2dGoroutine(moduleInfos, modulesParams))
		case "Flatten":
			modules = append(modules, flatten.NewFlatten(moduleInfos))
		case "Linear":
			modules = append(modules, linear.NewLinearGoroutine(moduleInfos, modulesParams))
		case "LogSoftmax":
			modules = append(modules, logsoftmax.NewLogSoftmax(moduleInfos))
		case "MaxPool2d":
			modules = append(modules, maxpool2d.NewMaxPool2dGoroutine(moduleInfos))
		case "ReLU":
			modules = append(modules, relu.NewReluGoroutine())
		default:
			panic(fmt.Sprintf("unrecognized module type: %s", moduleInfos.GetType()))
		}
	}

	return &model{modules}
}
