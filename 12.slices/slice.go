package main

import "fmt"

func main() {

	//1. cerating a slice form an array
	arr := [5]int{10, 20, 30, 40}
	slice := arr[1:4]
	fmt.Println(slice)

	//2.creating a slice directly - create a array without size = a slice
	directlySlice := []string{"this", "is a", "slice"}
	fmt.Println(directlySlice)

	//3.using make() funcion - this create a 3 length array which contains all 0s.
	slice3 := make([]int, 3)
	slice3 = append(slice3, 3) //adding new item to the last piosition of the slice
	fmt.Println(slice3)

	//3.1 with capapcity and length
	slice4 := make([]int, 3, 5) //3 means length 5 means capacity
	slice4 = append(slice4, 10, 87, 94, 63)
	fmt.Println(slice4)
	fmt.Println("Capacity of the aray is :", cap(slice4), " lenth of the aray is :", len(slice4))
	//mekedi me capacity eka resize wenkota eka kalin thibba size eka double wela thami resize wenne , dan 5 ekt 6th element ek awama 10 wenawa capacity eka , 10 thiyeddi 11th element ek awama 20 wenawa capapcty ek. anna ehmai yanne

	//4.range scn eka ena for loop eka
	ppls := []string{"Indura", "Nadeesha", "Anushka", "Tharusha", "Akila", "Dinuka"}
	for index, value := range ppls {
		fmt.Println(index, value)
	}

}
