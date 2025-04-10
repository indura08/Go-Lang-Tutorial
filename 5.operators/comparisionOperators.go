package main

import "fmt"

func main() {
	var num1 = 2
	var num2 = 3

	fmt.Println(num1 == num2) //false
	fmt.Println(num1 != num2) //true
	fmt.Println(num1 > num2)  //false
	fmt.Println(num1 < num2)  //true
	fmt.Println(num1 >= num2) //false
	fmt.Println(num1 <= num2) //true

}
