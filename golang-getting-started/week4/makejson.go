package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var err error
	var name string
	var address string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
	}
	name = strings.TrimSuffix(name, "\n")
	fmt.Print("Enter your address: ")
	address, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
	}
	address = strings.TrimSuffix(address, "\n")
	profile := map[string]string{
		"name":    name,
		"address": address,
	}
	profileJSON, err := json.Marshal(profile)

	if err == nil {
		fmt.Println(string(profileJSON))
	}
}
