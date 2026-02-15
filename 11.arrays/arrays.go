package main

import "fmt"

func main() {
	var nums [12]int

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = i
	}

	for j := 0; j < len(nums); j++ {
		fmt.Println(nums[j])
	}

	fmt.Println("Array printitng normal: ", nums)

	//string array
	var people = [3]string{"Indura", "Upasith", "Perera"}
	fmt.Println(people)

	fmt.Println("Capacity of the aray is :", cap(people), " lenth of the aray is :", len(people))

}
