package main

import (
	"fmt"
	"reflect"
)

// 1.generic function which takes any type of argument
func TakeNaythingFunction[T any](val T) {
	fmt.Println("Value is :", val)
	fmt.Println("Type is :", reflect.TypeOf(val))
}

// 2.generic struct
type Box[T any] struct {
	value T
}

func (b Box[T]) Get() T {
	return b.value
}

func main() {
	// 1. generic function
	TakeNaythingFunction(23.223)

	//2.generic struct example
	firstbox := Box[int]{value: 43}
	fmt.Println("Value of the First box : ", firstbox.Get())

	secondBox := Box[string]{value: "Hello"}
	fmt.Println("Value of the Second box : ", secondBox.Get())

}
