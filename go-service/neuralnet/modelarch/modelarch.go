package modelarch

import (
	"encoding/json"
	"iter"
	"os"
)

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

func LoadModelArch(archFile string) *modelArch {
	modelBytes, err := os.ReadFile(archFile)
	if err != nil {
		panic(err)
	}

	modelArch := &modelArch{}
	err = json.Unmarshal(modelBytes, modelArch)
	if err != nil {
		panic(err)
	}

	return modelArch
}
