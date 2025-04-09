package main

import "fmt"

func main() {

	var x int = 10
	//var y int = x++ - go waldi mehm eka paara sign krnnna bha wenama x++ daala thami y ekt x++ assign krnna puluwan

	x++
	var y = x
	//me widiyt kalama x ekath 11 y ekath 11 i x eka 10 thiyn y eka 11 krnna nm paatha widiyt trick ekk krnna one s

	var a int = 20
	var b = a
	a++
	b = a

	fmt.Println("Incerementing confusion: ")
	fmt.Println("Adding: ", a+b)
	fmt.Println("Substraction: ", a-b)
	fmt.Println("Multiplication: ", a*b)
	fmt.Println("division: ", a/b)

	fmt.Println()
	fmt.Println("adding: ", x+y)
	fmt.Println("Substract ", y-x)
	fmt.Println("Multiplication: ", x*y)
	fmt.Println("Division: ", y/x)

}
