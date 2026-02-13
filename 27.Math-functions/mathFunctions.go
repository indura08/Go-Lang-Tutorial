package main

import (
	"fmt"
	"math"
)

func main() {

	var x float64 = 4.7

	fmt.Println("Floor:", math.Floor(x))
	fmt.Println("Ceil:", math.Ceil(x))
	fmt.Println("Round:", math.Round(x))
	fmt.Println("Square root:", math.Sqrt(x))
}
