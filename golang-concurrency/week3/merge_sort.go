package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func mergeSort(list []int) []int {
	length := len(list)
	if length <= 1 {
		return list
	}

	mid := length / 2
	left := mergeSort(list[:mid])
	right := mergeSort(list[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, len(left)+len(right))
	i := 0
	for ; len(left) > 0 && len(right) > 0; i++ {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}

	}

	for j := 0; j < len(left); j++ {
		result[i+j] = left[j]
	}

	for j := 0; j < len(right); j++ {
		result[i+j] = right[j]
	}
	return result
}

func worker(wg *sync.WaitGroup, list *[]int) {
	defer wg.Done()

	fmt.Printf("subroute list befor sorting: %v\n", *list)
	*list = mergeSort(*list)
	fmt.Printf("subroute list after sorting: %v\n", *list)
}

func main() {
	var wg sync.WaitGroup

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please input integers, sperated by whitespace: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
	}

	intList := strings.TrimSuffix(input, "\n")
	inputSlice := strings.Split(intList, " ")
	list := make([]int, len(inputSlice))
	for i, str := range inputSlice {
		convertedInt, err := strconv.Atoi(str)
		if err == nil {
			list[i] = convertedInt
		}
	}
	mid := len(list) / 4
	start := 0
	end := mid
	slices := make([][]int, 4)
	fmt.Printf("list befor sorting: %v\n", list)
	for i := 0; i < 4; i++ {
		fmt.Println(start, end)
		fmt.Println("Main: Waiting for workers to finish")
		if i < 3 {
			slices[i] = list[start:end]
		} else {
			slices[i] = list[start:]
		}

		wg.Add(1)
		go worker(&wg, &slices[i])
		start = end
		end = start + mid
	}
	wg.Wait()

	result := merge(merge(slices[0], slices[1]), merge(slices[2], slices[3]))
	fmt.Printf("sorted list after sorting: %v\n", result)
}
