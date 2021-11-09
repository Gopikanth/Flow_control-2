package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for { //using for to continue the game again and again
		fmt.Println()

		fmt.Println("Start the Game")
		fmt.Println(" ---------->")
		fmt.Println()
		fmt.Println("Read instructions carefully")
		fmt.Println()
		fmt.Println("Enter 0 for Stone")
		fmt.Println()
		fmt.Println("Enter 1 for Paper")
		fmt.Println()
		fmt.Println("Enter 2 for Scissor")

		rand.Seed(time.Now().UnixNano())
		machine_input := rand.Intn(2)
		var user int
		fmt.Scanln(&user)
		//getting into the functionality
		if user > 2 {
			fmt.Println("Enter the number as per instructions")
			continue

		}
		if user == machine_input {
			fmt.Println("you win")
			continue

		}

		fmt.Println("you loose")
		continue

	}
}
