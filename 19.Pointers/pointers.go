package main

import "fmt"

func main() {

	//1.pointer mini example
	num := 10
	ptr := &num

	fmt.Println("Value of the num variable is :", num)
	fmt.Println("Addresss of the num variable is :", &num)
	fmt.Println("Value of the pointer variable is, and see how it is equal to address of the num vraible :", ptr)

	//2.dereferenecing
	number := 42
	ptr1 := &number

	fmt.Println("real value of the pointer variable is : ", *ptr1, " memeory address of the pinter varaible is ", ptr1)

	//3.changing the value of a avraible using pointer
	number2 := 5
	ptr2 := &number2

	fmt.Println("Value of the number2 variable is before changing using pointer :", number2)

	*ptr2 = 100
	fmt.Println("Value of the number2 variable is after changing using pointer :", number2)

	//4.pointers in functions
	changeValue := func(val *int) {
		*val = 100
	}

	number3 := 3
	fmt.Println("Value of the number3 variable is before changing using pointer :", number3)
	changeValue(&number3)
	fmt.Println("Value of the number3 variable is after changing using pointer :", number3)

}
