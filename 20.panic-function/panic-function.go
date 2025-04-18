package main

import "fmt"

// 1.basic example
func BasicExample() {
	fmt.Println("This is the first fmt prinln satatement")
	panic("Oh no somehting bad happend")
	fmt.Println("This is the second fmt prinln satatement and this will never be printed")
}

// 2.real world exmaple - divide by zero
func DivisionByzero(a int) int {
	if a == 0 {
		panic("You cant divide numbers by zero you idiiot, look what you have done , you crashed the whole program!!")
	}
	z := 10 / a

	return z
}

// 3. panic with recover function (meke note eke thiynwa notes wala , meken wenne recover function ekak use karala panic ekak handle karanna)
func SafeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("This function is about to get panic")
	panic("Oh no somehting went wrong, pancic triggered")
	fmt.Println("this will not be printed")

}

func main() {
	//BasicExample()
	//DivisionByzero(0)
	SafeFunction()
	fmt.Println("i think this will also not be printed, but if recover function is there this will be printed")

}
