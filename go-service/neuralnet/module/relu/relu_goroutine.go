package relu

import (
	"fmt"
	"reflect"
	"sync"
)

type reluGoroutine struct{}

// in-place op
func (reluGoroutine reluGoroutine) Forward(inputPtr any) any {
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

	var wg sync.WaitGroup

	wg.Add(1)
	go reluGoroutine.reluGoroutine(&wg, input)

	wg.Wait()

	return nil
}

func (reluGoroutine reluGoroutine) reluGoroutine(wg *sync.WaitGroup,
	val reflect.Value) {

	defer wg.Done()

	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := range val.Len() {
			el := val.Index(i)

			switch el.Kind() {
			case reflect.Array, reflect.Slice:
				wg.Add(1)
				go reluGoroutine.reluGoroutine(wg, el)
			case reflect.Float32:
				if el.Float() < 0 {
					el.SetZero()
				}
			default:
				panic(fmt.Sprintf("ReLU: invalid type %v", el.Kind()))
			}
		}
	default:
		panic(fmt.Sprintf("ReLU: expected a slice but got %v", val.Kind()))
	}
}

func NewReluGoroutine() reluGoroutine {
	return reluGoroutine{}
}
