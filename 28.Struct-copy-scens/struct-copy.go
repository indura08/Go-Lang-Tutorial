package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func changeName(u User) {
	u.Name = "Jhon"
	fmt.Println("The user name after chanigng {note that this print is inside the function }", u.Name)
}

func main() {

	user := User{Name: "Indura", Age: 17}

	fmt.Println("Users name is : ", user.Name)

	changeName(user)
	fmt.Println("Users name after changing the object is {outside the changeName function} : ", user.Name)

	//in this case note that when we pass the copy of a struct it is changing the passing copy only not the original object
	//so this is why we pass pointer variables in cotroller methods , middlewares when it comes to a gin application
	//note that other programming languages like c# java passes the object , but under the hood it passes the erefrence , not a copy of a obejct

}
