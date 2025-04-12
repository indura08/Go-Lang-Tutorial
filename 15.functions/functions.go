package main

import (
	"fmt"
)

// 1. basic function - no return type no parameters
func sayHello() {
	fmt.Println("Hello world from function 1")
}

// 2.function with parameters
func sayHelloWithName(name string) {
	fmt.Println("hello there,", name)
}

// 3.function with retunr types
func GetCircumferenec(radius int) int { //hodt blnna methan parameters walt passe eturn type eka daanna one typescript walal wage
	var c int = 2 * 3 * radius
	return c
}

// 4.multiple return values inlcuded function
func Division(a int, b int) (int, int) {
	normaldivision := a / b
	modules := a % b

	return normaldivision, modules
}

// 5.named retunr values
func GetName() (firstname string, lastname string) {
	firstname = "Eliana"
	lastname = "Vivienne"

	return
}

// variadic functions- where you dont know how manay arguments it will be taken
func DiffrentArguments(nums ...int) int { //meke thwa arguments danwa num...int walt issrahin daanna oen pitipassen daanna bha
	var total int = 0
	for _, num := range nums {
		total += num
	}

	return total
}

// function inside function : thawa function ekk athule named function ekk ghnna bha anonymous function ekk ghnna one
func OuterFunc(name string) int {
	func() {
		fmt.Println("Helo from outer function")
	}()

	NameCounter := func(name1 string) int {
		//fmt.Println("Nuimber of letters in your name is :", len(name1))
		return len(name1)
	}

	var chars int = NameCounter(name)

	return chars
}

func main() {
	//main method eka athule liynna bha named funtion, liynna puluwan function expressions withri

	sayHello()
	sayHelloWithName("Indura perera")

	c := GetCircumferenec(7)
	fmt.Println(c)

	a, b := Division(2, 2)
	fmt.Println("Normal division value is :", a, " modules value is ", b)

	fmt.Println(GetName())

	//variaou argument numbers
	fmt.Println(DiffrentArguments(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(DiffrentArguments(2, 3, 4))

	//function inside function
	var chars int = OuterFunc("Indura")
	fmt.Println(chars)

	//function expressions
	add := func(x int, y int) int {
		z := x + y
		return z
	}

	fmt.Println(add(23, 32))

	//anonymous functions
	var total int = func(a int, b int) int {
		return a + b
	}(12, 41)

	fmt.Println("anonymous function's result is:", total)

}
