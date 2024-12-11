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

func (self moduleInfo) GetType() string {
	return self.Type
}

func (self moduleInfo) GetProp(name string) (json.RawMessage, bool) {
	val, exists := self.Props[name]
	return val, exists
}