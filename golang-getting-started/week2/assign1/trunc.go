package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64

	fmt.Print("Enter the number: ")
	fmt.Scanln(&input)
	int, _ := math.Modf(input)
	fmt.Println("int part: ", int)
}
