package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertInputToSlice(input string) []int {
	withoutBreakLineInput := strings.TrimSuffix(input, "\n")
	stringSlice := strings.Split(withoutBreakLineInput, " ")
	result := make([]int, len(stringSlice))

	for i, str := range stringSlice {
		integer, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("error: ", err)
		}
		result[i] = integer
	}

	return result
}

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter at most 10 integers: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
	}

	intSlice := convertInputToSlice(input)
	bubbleSort(intSlice)

	fmt.Printf("%v\n", intSlice)
}
func bubbleSort(slice []int) {
	for i := len(slice); i > 0; i-- {
		for j := 1; j < i; j++ {
			if slice[j-1] > slice[j] {
				swap(slice, j-1)
			}
		}
	}
}
func swap(values []int, index int) {
	values[index+1], values[index] = values[index], values[index+1]
}
