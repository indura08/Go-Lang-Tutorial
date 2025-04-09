package main

import (
	"fmt"
	"strings"
)

func main() {

	var firstname string = "Inuri"
	var lastName string = "Eliana"

	//concatination
	var fullName string = firstname + " " + lastName
	fmt.Println("Concatinations (fullName): ", fullName)

	//string length
	var length int = len(fullName)
	fmt.Println("String length of full name : ", length)

	//accessing character in a string
	var thirdchar string = string(fullName[2])
	fmt.Println("third character of fullname variable: ", thirdchar)

	//string comparision
	fmt.Println("Is firstname and lastname equal?: ", firstname == lastName)
	fmt.Println("string comparision with strings.Compare() method: ", strings.Compare(firstname, lastName)) //mehemaht puluwan , mekn false nm -1 da true nm 1 da enwa

	//string to uppercase lowercase
	fmt.Println("FullName in upperCase: ", strings.ToUpper(fullName))
	fmt.Println("FullName in lowerCase: ", strings.ToLower(fullName))

	//palceholder ekk widiyt character ekk daana widiya:
	fmt.Println(string(fullName[4])) //mehma dammoth i enwa string kalla nathuwa dammoth pahala wage 105 enwa
	fmt.Println(fullName[4])
	fmt.Printf("%c", fullName[4])
	fmt.Println()

	//contains method ek
	fmt.Println(strings.Contains(fullName, "ri"))

	//charcter ekk define krnwa nm data type eka awailla rune
	var firstchracter rune = 'A'
	fmt.Println(firstchracter)         //mehm kalama 65 enne
	fmt.Println(string(firstchracter)) //mehm kalama A enwas

	//multiline string ekk danwa nm
	var multiLine string = `This 
							is a 
							multiLine 
								string` //back tik ekn patn gnna one multiline string ekk patn gnnwa nm
	fmt.Println(multiLine)

}
