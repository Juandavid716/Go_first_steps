package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Animal structure
type Animal interface {
	Eat()
	Move()
	Speak()
}

//Cow type
type Cow struct {
	food, locomotion, noise string
}

//Bird type
type Bird struct {
	food, locomotion, noise string
}

//Snake type
type Snake struct {
	food, locomotion, noise string
}

//Eat - Cow
func (cow Cow) Eat() {
	fmt.Println(cow.food)
}

//Eat - Bird
func (bird Bird) Eat() {
	fmt.Println(bird.food)
}

//Eat - Snake
func (snake Snake) Eat() {
	fmt.Println(snake.food)
}

//Move - Cow
func (cow Cow) Move() {
	fmt.Println(cow.locomotion)
}

//Move - Bird
func (bird Bird) Move() {
	fmt.Println(bird.locomotion)
}

//Move - Snake
func (snake Snake) Move() {
	fmt.Println(snake.locomotion)
}

//Speak - Cow
func (cow Cow) Speak() {
	fmt.Println(cow.noise)
}

//Speak - Bird
func (bird Bird) Speak() {
	fmt.Println(bird.noise)
}

//Speak - Snake
func (snake Snake) Speak() {
	fmt.Println(snake.noise)
}

func main() {

	var animals = make(map[string]string)

	for {
		fmt.Println("> Enter request (newanimal, nameanimal, typeanimal) OR (query, nameanimal, method)")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		split := strings.Split(text, " ")

		//fmt.Printf("request is: %s\n", request)

		if len(split) != 3 {
			fmt.Println("Lengh of input must be 3")
			continue
		}

		fmt.Printf("command: %s\n", text)

		command := split[0]
		name := split[1]
		animalType := split[2]
		fmt.Printf("name: %s\n", name)
		switch command {
		case "newanimal":

			fmt.Printf("typeanimal: %s\n", animalType)

			_, ok := animals[name]
			if ok == false && (animalType == "cow" || animalType == "bird" || animalType == "snake") {
				animals[name] = animalType
				fmt.Println("Created")
			} else {
				fmt.Println("Try with another name")
			}
		case "query":
			method := animalType
			fmt.Printf("method: %s\n", method)

			switch animals[name] {
			case "cow":
				cow := Cow{"grass", "walk", "moo"}
				switch method {
				case "eat":
					cow.Eat()
				case "move":
					cow.Move()
				case "speak":
					cow.Speak()
				default:
					fmt.Println("Try again with one of these methods (eat, move, speak)", method)
					continue
				}
			case "bird":
				bird := Bird{"worms", "fly", "peep"}
				switch method {
				case "eat":
					bird.Eat()
				case "move":
					bird.Move()
				case "speak":
					bird.Speak()
				default:
					fmt.Println("Try again with one of these methods (eat, move, speak)", method)
					continue
				}
			case "snake":
				snake := Snake{"mice", "slither", "hsss"}
				switch method {
				case "eat":
					snake.Eat()
				case "move":
					snake.Move()
				case "speak":
					snake.Speak()
				default:
					fmt.Println("Try again with one of these methods (eat, move, speak)", method)
					continue
				}
			default:
				fmt.Println("Try with one of these animal names (cow, bird, snake)", animals[name])
				continue
			}
		default:
			fmt.Println("Try with one of these options (command, query")
			continue
		}
	}
}
