package main

import "fmt"

func main() {

	var ismanager bool = true
	var isfather bool = true

	if ismanager && isfather {
		fmt.Println("User is a father and also a manager")
	} else if ismanager && !isfather {
		fmt.Println("this is a manager user only")
	} else { //me else eka liyaddi hariytm '}' n passe ma liynna one newline ekakat yanna bha
		fmt.Println("This user is a father but not a manager")
	}

}
