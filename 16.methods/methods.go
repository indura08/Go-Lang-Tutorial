package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	isManager bool
}

type student struct {
	id     int
	person Person
}

func (p Person) GetFullName(message string) string {
	return p.firstName + " " + p.lastName + " and your message is:" + message
}

func main() {
	//method kiyna ewa hdnna mokkma hari eka type ekkt , spescilay use krnne struct walti

	induraPerson := Person{firstName: "Indura", lastName: "Perera", age: 24, isManager: true}
	fmt.Println("first Person's full name is : ", induraPerson.GetFullName("This is my message to you"))

	//me exmaple ek onnm uncomment krla blnna GetFullName method ek use krnna bha
	// student1 := student{id: 1.0, person: induraPerson}
	// fmt.Println(student1.GetFullName("My name is sttudent"))

}
