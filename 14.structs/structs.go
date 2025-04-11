package main

import (
	"fmt"
)

type Club struct {
	name    string
	country string
	owner   Person1
}

type Person1 struct {
	name string
	age  int
}

func main() {

	//1.creating a struct (1st way)
	type Person struct {
		name      string
		age       int
		isManager bool
	}

	var induraPerson Person
	induraPerson.age = 25
	induraPerson.name = "Indura perera"
	induraPerson.isManager = true
	fmt.Println(induraPerson)

	//2. creating a struct (2nd way)
	upasithPerson := Person{name: "Upasith Perera", age: 24, isManager: false}
	fmt.Println(upasithPerson)

	// //3 using range to print values with attributes of a struct
	// pereraPerson := Person{name: "Perera Wijesighe", age: 23, isManager: false}

	// for attribute, value := range pereraPerson {
	// 	fmt.Println("attribute=", attribute, " value=", value)
	// }		struct wala mehm krnna bha error eka blnna one nm uncomment krla

	//4.attribute wenama ganima
	fmt.Println("Age of induraPerson is :", induraPerson.age)

	//struct inside a struct
	var club1 Club
	club1.name = "Bayern Munich"
	club1.country = "Germany"
	club1.owner = Person1{name: "Inter milano", age: 120}

	fmt.Println("club is : ", club1)

	//5,anonymous struct
	mbappe := struct {
		firstname string
		lastName  string
		age       int
	}{
		firstname: "Kylian",
		lastName:  "Mbappe",
		age:       25,
	}

	fmt.Println(mbappe)

}
