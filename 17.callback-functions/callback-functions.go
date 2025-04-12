package main

import "fmt"

func division(callback func(int, int) int, a int, b int) (division int, whateverHappenwithCallback int) {
	division = a / b
	whateverHappenwithCallback = callback(a, b)

	return division, whateverHappenwithCallback
}

func Sum(a int, b int) int {
	return a + b
}

func main() {

	divisionValue, callbackvalue := division(Sum, 20, 10)
	fmt.Println("CallbackFunctions result: ", callbackvalue, " and divisions fuctions result : ", divisionValue)

}
