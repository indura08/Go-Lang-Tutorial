package main

import "fmt"

//1.example

// define the interface
type Animal interface {
	speak() string
}

// define a dog struct
type Dog struct {
	name string
	legs int
}

// method for dog class
func (d Dog) speak() string {
	return "This dog barks , baw baw!!"
}

// defining a cat struct
type Cat struct {
	name string
	legs int
}

// defining a method for cat class
func (c Cat) speak() string {
	return "This cat meows , meow meow!!"
}

func main() {

	//now we can use dogs and cats as animals
	var a Animal

	a = Dog{"dog", 4}
	fmt.Println(a.speak())

	a = Cat{"cat", 4}
	fmt.Println(a.speak())

}
