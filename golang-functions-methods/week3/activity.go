package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal class
type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) eat() string {
	return animal.food
}
func (animal *Animal) move() string {
	return animal.locomotion
}
func (animal *Animal) speak() string {
	return animal.noise
}

func main() {
	cow := &Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := &Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := &Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Please input animal and behavior, seperate by whitespace, or exit by 'X'> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
		} 
		
		inputString := strings.TrimSuffix(input, "\n")
		if inputString == "X" {
			fmt.Println("Goodbye!")
			break
		}

		inputSlice := strings.Split(inputString, " ")

		var animal *Animal

		switch inputSlice[0] {
		case "cow":
			animal = cow
		case "bird":
			animal = bird
		case "snake":
			animal = snake
		default:
			fmt.Printf("Don't know %s is what animal\n", inputSlice[0])
			continue
		}

		switch inputSlice[1] {
		case "eat":
			fmt.Println(animal.eat())
		case "move":
			fmt.Println(animal.move())
		case "speak":
			fmt.Println(animal.speak())
		default:
			fmt.Printf("Don't know what is the behavior %s\n", inputSlice[1])
			continue
		}
	}
}
