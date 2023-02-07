package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("Hello Welcome to Rock-Paper-Scissors Game :)\n\n")
	fmt.Print("If you want to start the game, write the word `start` : ")
	var started string
	fmt.Scan(&started)
	if started == "start" {
		fmt.Print("\n\n-_- well down the game is start -_-\n\n")
		RockPaperScissors()
	} else {
		fmt.Print("\n\n\n---_--- End Game ---_---\n\n")
		fmt.Print("\nInvalid input Please enter a valid input\n")
	}
}

func RockPaperScissors() {
	random := [3]string{"rock", "paper", "scissors"}
	var user string

	fmt.Print(" please choice `rock` or `paper` or `scissors` -write-> : ")
	fmt.Scan(&user)
	computer := random[rand.Intn(len(random))]
	fmt.Println(user, computer)
}
