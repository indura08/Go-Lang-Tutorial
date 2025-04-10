package main

import "fmt"

func main() {

	//in go , there is no keyword for while lop,
	//to ceate a for loop, while loop or even infinate loop we use for keyword

	var num int = 0

	for num < 15 {
		fmt.Println("number is : ", num)
		num++
	}
}
