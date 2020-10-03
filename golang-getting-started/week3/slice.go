package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sortedInputs := make([]int, 0)
	var input string

	for {
		fmt.Print("Enter an integer or quit by input 'X': ")
		fmt.Scanln(&input)
		inputedInt, err := strconv.Atoi(input)

		if input == "X" {
			fmt.Println("Goodbye!")
			break
		} else if err != nil || strconv.Itoa(inputedInt) != input {
			fmt.Println("Cannot infer what it is, please try again.")
			continue
		}

		index := sort.Search(len(sortedInputs), func(i int) bool { return sortedInputs[i] >= inputedInt })
		sortedInputs = append(sortedInputs, 0)
		copy(sortedInputs[index+1:], sortedInputs[index:])
		sortedInputs[index] = inputedInt
		fmt.Println("You input the array so far:")
		fmt.Printf("%v\n", sortedInputs)
	}
}
