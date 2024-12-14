package module

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/jfriedson/mnist-go-electron-react-app/go-service/neuralnet/modelarch"
)

type flatten struct {
	startDim int
	endDim   int
}

// TODO: function currently flattens all layers. flatten only specified layers
func (flatten flatten) Forward(inputPtr any) any {
	inputPtrVal := reflect.ValueOf(inputPtr)
	if inputPtrVal.Kind() != reflect.Pointer || inputPtrVal.IsNil() {
		panic("Flatten: input must be a non-nil pointer")
	}

	input := inputPtrVal.Elem().Interface()
	inputVal := reflect.ValueOf(input)
	if inputVal.Kind() == reflect.Float32 {
		return inputVal.Interface()
	}

	// output := []any{}
	output := []float32{}

	type qEl struct {
		val reflect.Value
		// depth int
	}
	q := []qEl{{inputVal /*0*/}}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		switch cur.val.Kind() {
		case reflect.Array, reflect.Slice:
			for i := range cur.val.Len() {
				el := cur.val.Index(i)

				switch el.Kind() {
				case reflect.Array, reflect.Slice:
					q = append(q, qEl{el /*, el.depth + 1*/})
				case reflect.Float32:
					output = append(output, el.Interface().(float32))
				default:
					panic(fmt.Sprintf("Flatten: invalid type %v", el.Kind()))
				}
			}
		default:
			panic(fmt.Sprintf("Flatten: expected a slice but got %v", cur.val.Kind()))
		}
	}

	return output
}

func NewFlatten(moduleInfo modelarch.ModuleInfo) flatten {
	var startDim, endDim int

	raw, exists := moduleInfo.GetProp("start_dim")
	if !exists {
		startDim = 0
	} else {
		err := json.Unmarshal(raw, &startDim)
		if err != nil {
			panic("Flatten: start_dim must be a number")
		}
	}

	raw, exists = moduleInfo.GetProp("end_dim")
	if !exists {
		endDim = -1
	} else {
		err := json.Unmarshal(raw, &endDim)
		if err != nil {
			panic("Flatten: end_dim must be a number")
		}
	}

	if startDim < 0 {
		panic("Flatten: start_dim must be 0 or greater")
	}

	if endDim < -1 {
		panic("Flatten: start_dim must be -1 or greater")
	}

	if endDim > -1 && startDim <= endDim {
		panic("Flatten: end_dim must be greater than start_dim if end_dim is not -1")
	}

	return flatten{startDim, endDim}
}
