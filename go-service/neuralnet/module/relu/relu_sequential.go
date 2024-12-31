package relu

import (
	"fmt"
	"reflect"
)

type reluSequential struct{}

// in-place op
func (reluSequential reluSequential) Forward(inputPtr any) any {
	inputPtrVal := reflect.ValueOf(inputPtr)
	if inputPtrVal.Kind() != reflect.Pointer || inputPtrVal.IsNil() {
		panic("ReLU: input must be a non-nil pointer")
	}

	input := inputPtrVal.Elem()
	if input.Kind() == reflect.Float32 {
		if input.Float() < 0 {
			input.SetZero()
		}
		return nil
	}

	var stack []reflect.Value
	stack = append(stack, input)

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		switch cur.Kind() {
		case reflect.Array, reflect.Slice:
			for i := range cur.Len() {
				el := cur.Index(i)

				switch el.Kind() {
				case reflect.Array, reflect.Slice:
					stack = append(stack, el)
				case reflect.Float32:
					if el.Float() < 0 {
						el.SetZero()
					}
				default:
					panic(fmt.Sprintf("ReLU: invalid type %v", el.Kind()))
				}
			}
		default:
			panic(fmt.Sprintf("ReLU: expected a slice but got %v", cur.Kind()))
		}
	}

	return nil
}

func NewReluSequential() reluSequential {
	return reluSequential{}
}
