package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal interface {
	eat()
	move()
	speak()
	getName() string
}

type cow struct {
	food       string
	locomotion string
	noise      string
	name       string
}
type bird struct {
	food       string
	locomotion string
	noise      string
	name       string
}
type snake struct {
	food       string
	locomotion string
	noise      string
	name       string
}

func (a cow) eat() {
	fmt.Println(a.food)
}
func (b bird) eat() {
	fmt.Println(b.food)
}
func (s snake) eat() {
	fmt.Println(s.food)
}

func (a cow) move() {
	fmt.Println(a.locomotion)
}
func (b bird) move() {
	fmt.Println(b.locomotion)
}
func (s snake) move() {
	fmt.Println(s.locomotion)
}
func (a cow) speak() {
	fmt.Println(a.noise)
}
func (b bird) speak() {
	fmt.Println(b.noise)
}
func (s snake) speak() {
	fmt.Println(s.noise)
}
func (a cow) getName() string {
	return a.name
}
func (b bird) getName() string {
	return b.name
}
func (s snake) getName() string {
	return s.name
}

func main() {
	var newAnimal animal
	reader := bufio.NewReader(os.Stdin)
	animals := []animal{}
	for {
		fmt.Print("Please newanimal, query, or exit by 'X'> ")
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
		switch inputSlice[0] {
		case "newanimal":
			switch inputSlice[1] {
			case "cow":
				newAnimal = cow{name: inputSlice[1], food: "grass", locomotion: "walk", noise: "moo"}
			case "bird":
				newAnimal = bird{name: inputSlice[1], food: "worms", locomotion: "fly", noise: "peep"}
			case "snake":
				newAnimal = snake{name: inputSlice[1], food: "mice", locomotion: "slither", noise: "hsss"}
			default:
				fmt.Printf("Don't know %s is what animal\n", inputSlice[1])
				continue
			}
			animals = append(animals, newAnimal)
			fmt.Println("Created it!")
		case "query":
			found := false
			for _, theAnimal := range animals {
				if theAnimal.getName() == inputSlice[1] {
					found = true
					switch inputSlice[2] {
					case "eat":
						theAnimal.eat()
					case "move":
						theAnimal.move()
					case "speak":
						theAnimal.speak()
					default:
						fmt.Printf("Don't know what is the behavior %s\n", inputSlice[2])
					}
				}
			}
			if !found {
				fmt.Printf("Don't know %s is what animal\n", inputSlice[1])
			}
		default:
			fmt.Printf("Don't know %s is what action\n", inputSlice[0])
		}
	}
}
