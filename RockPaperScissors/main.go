package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print("Hello Welcome to Rock-Paper-Scissors Game :)\n\n")
	fmt.Print("please choice `start` or `exit` : ")
	var input string
	fmt.Scan(&input)
	if input == "start" {
		fmt.Print("\n\n-_- well down the game is start -_-\n\n")
		RockPaperScissors()
	} else if input == "exit" {
		fmt.Println("--- EXIT ---")
	} else {
		fmt.Print("\n--- EXIT ---\n")
		fmt.Print("\nInvalid input Please enter a valid input\n")
	}
}

func RockPaperScissors() {
	selection := [3]string{"rock", "paper", "scissors"}
	var user string

	fmt.Print("please choice `rock` or `paper` or `scissors` -write-> : ")
	fmt.Scan(&user)
	rand.Seed(time.Now().UnixNano())
	randomNum := random(0, 3)
	computer := selection[randomNum]
	RulesOfTheGame(user, computer)
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func RulesOfTheGame(user string, computer string) {
	if user == "rock" && computer == "rock" {
		equal := "--- is equals ---"
		fmt.Println(equal)
	} else if user == "rock" && computer == "paper" {
		userWin := "\n ^--^ you are win ^--^"
		fmt.Println(userWin)
	} else if user == "rock" && computer == "scissors" {
		computerWin := "-_- you are loss :)"
		fmt.Println(computerWin)
	} else if user == "paper" && computer == "paper" {
		equal := "--- is equals ---"
		fmt.Println(equal)
	} else if user == "paper" && computer == "rock" {
		userWin := "\n ^--^ you are win ^--^"
		fmt.Println(userWin)
	} else if user == "paper" && computer == "scissors" {
		computerWin := "-_- you are loss :)"
		fmt.Println(computerWin)
	}
}
