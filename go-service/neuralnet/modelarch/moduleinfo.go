package modelarch

import (
	"encoding/json"
)

type ModuleInfo interface {
	GetType() string
	GetProp(string) (json.RawMessage, bool)
}

type moduleInfo struct {
	Type  string                     `json:"type"`
	Props map[string]json.RawMessage `json:"props"`
}

func (moduleInfo moduleInfo) GetType() string {
	return moduleInfo.Type
}

func (moduleInfo moduleInfo) GetProp(name string) (json.RawMessage, bool) {
	val, exists := moduleInfo.Props[name]
	return val, exists
}
