package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Animal method
type Animal struct {
	food, locomotion, noise string
}

//Eat method
func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

//Move method
func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

//Speak method
func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {

	var an Animal

	for {
		fmt.Println("> Enter name of animal and method:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		split := strings.Split(text, " ")
		animal := split[0]
		method := split[1]

		switch animal {
		case "cow":
			an = Animal{"grass", "walk", "moo"}

		case "bird":
			an = Animal{"worms", "fly", "peep"}
		case "snake":
			an = Animal{"mice", "slither", "hsss"}
		default:
			fmt.Printf("Animal not found")
			return
		}

		switch method {
		case "eat":
			an.Eat()
		case "move":
			an.Move()
		case "speak":
			an.Speak()
		default:
			fmt.Printf("Method not found")
			return
		}

	}

}
