package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var filename string
	var err error
	names := make([]name, 0)

	fmt.Print("Enter file path: ")
	fmt.Scanln(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nameSlice := strings.Split(scanner.Text(), " ")
		newName := name{fname: nameSlice[0], lname: nameSlice[1]}
		names = append(names, newName)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error: ", err)
	}

	for _, currentName := range names {
		fmt.Printf("%+v\n", currentName)
	}
}
