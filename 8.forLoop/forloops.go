package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("the number is : ", i)
	}

	var nums string = "123456789"
	fmt.Println("Lets see the characters of the nums variable")

	for j := 0; j < len(nums); j++ {
		fmt.Print(string(nums[j]), ", ")
	}

}
