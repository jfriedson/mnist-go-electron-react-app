package neuralnet

import (
	"encoding/json"
	"fmt"
	"iter"
	"os"
)

type ModuleInfo interface {
	GetType() string
	GetProp(string) (any, bool)
}

type moduleInfo struct {
	Type  string         `json:"type"`
	Props map[string]any `json:"props"`
}

func (self moduleInfo) GetType() string {
	return self.Type
}

func (self moduleInfo) GetProp(name string) (any, bool) {
	val, exists := self.Props[name]
	return val, exists
}

type ModelArch interface {
	GetModuleInfos() iter.Seq[ModuleInfo]
}

type modelArch []moduleInfo

func (self modelArch) GetModuleInfos() iter.Seq[ModuleInfo] {
	return func(yield func(ModuleInfo) bool) {
		for _, v := range self {
			if !yield(v) {
				return
			}
		}
	}
}

func loadModelArch(archFile string) *modelArch {
	modelBytes, err := os.ReadFile(archFile)
	if err != nil {
		panic(err)
	}

	modelArch := &modelArch{}
	err = json.Unmarshal(modelBytes, modelArch)
	if err != nil {
		panic(err)
	}

	fmt.Println(modelArch)

	return modelArch
}
