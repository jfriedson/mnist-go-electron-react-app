package module

import (
	"reflect"
)

type flatten struct {
	start_dim int
	end_dim   int
}

func NewFlatten(start_dim int) *flatten {
	if reflect.TypeOf(start_dim).Kind() != reflect.Uint {
		panic("flatten input must not be 0 dimensional")
	}

	return &flatten{start_dim: 0, end_dim: 0}
}

func (self *flatten) Forward(input any) any {
	if reflect.TypeOf(input).Kind() != reflect.Slice {
		panic("flatten input must not be 0 dimensional")
	}

	return uint(21)
}
