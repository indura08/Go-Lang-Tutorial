package main

import "fmt"

func main() {
	fmt.Println("Hello from switch case")

	var day int = 1

	switch day {
	case 1:
		fmt.Println("The day is Monday")

	case 2:
		fmt.Println("The day is Tuesday")

	case 3:
		fmt.Println("The day is wednesday")

	case 4:
		fmt.Println("The day is Thursday")

	case 5:
		fmt.Println("The day is friday")

	case 6:
		fmt.Println("The day is saturday")

	case 7:
		fmt.Println("The day is sunday")

	default:

		fmt.Println("The number is greated than 7 or less that 1")

	}
}
