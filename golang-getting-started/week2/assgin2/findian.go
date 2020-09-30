package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Enter a atring strat with i, end with n and contain a: ")
	fmt.Scanln(&input)
	lowerInput := strings.ToLower(input)
	length := len(input)
	if lowerInput[0] != 'i' || lowerInput[length-1] != 'n' {
		fmt.Println("Not Found")
		return
	}
	for i := 1; i < length-1; i++ {
		if lowerInput[i] == 'a' {
			fmt.Println("Found")
			return
		}
	}
	fmt.Println("Not Found")
}
