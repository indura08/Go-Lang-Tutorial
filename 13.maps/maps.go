package main

import "fmt"

func main() {
	fmt.Println("hello from maps")
	fmt.Println("----------------")

	//1.create map and add data at once
	var map1 = map[int]string{
		1: "Indura perera",
		2: "Selenan gomez",
		3: "Sydney sweeney",
	}
	fmt.Println(map1)

	//2.create map and add values later
	//make is a built-in Go function that is used to create and initialize special types of data structures like: Slices, Maps, Channels
	map2 := make(map[string]string)

	map2["name"] = "Indura perera"
	map2["age"] = "24"
	map2["city"] = "Kiribathgoda"

	fmt.Println(map2)
	fmt.Println(map2["age"])

	//3.to check if a key exists
	value, exists := map2["age"] //go waldi mehm map2["age"] kiyl dunnama eke value eki eka thiynwad nadda kiyl boolean value ekakui denwa , eka go walal wisheshathwayak
	if exists {
		fmt.Println("Value is : ", value)
	} else {
		fmt.Println("Value is not there at the moment")
	}

	//4-delete a key from map
	delete(map2, "age")
	fmt.Println("the maps after deleteing age key is : ", map2)

	//5-using the range keyword to print the data with value + key
	for x, y := range map2 {
		fmt.Println("Key=", x, "value:", y)
	}

}
