package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

// score's
var user_score = 0
var computer_score = 0

func RulesOfTheGame(user string, computer string) {
	if user == "rock" && computer == "rock" {
		equal := "--- is equals ---"
		user_score += 1
		computer_score += 1
		fmt.Println(equal)
	} else if user == "rock" && computer == "paper" {
		computerWin := "-_- you are loss :)"
		computer_score += 1
		fmt.Println(computerWin)
	} else if user == "rock" && computer == "scissors" {
		userWin := "\n ^--^ you are win ^--^"
		user_score += 1
		fmt.Println(userWin)
	} else if user == "paper" && computer == "paper" {
		equal := "--- is equals ---"
		user_score += 1
		computer_score += 1
		fmt.Println(equal)
	} else if user == "paper" && computer == "rock" {
		userWin := "\n ^--^ you are win ^--^"
		user_score += 1
		fmt.Println(userWin)
	} else if user == "paper" && computer == "scissors" {
		computerWin := "-_- you are loss :)"
		computer_score += 1
		fmt.Println(computerWin)
	} else if user == "scissors" && computer == "scissors" {
		equal := "--- is equals ---"
		user_score += 1
		computer_score += 1
		fmt.Println(equal)
	} else if user == "scissors" && computer == "paper" {
		userWin := "\n ^--^ you are win ^--^"
		user_score += 1
		fmt.Println(userWin)
	} else if user == "scissors" && computer == "rock" {
		computerWin := "-_- you are loss :)"
		computer_score += 1
		fmt.Println(computerWin)
	}
}

func main() {
	fmt.Print("Hello Welcome to Rock-Paper-Scissors Game :)\n\n")
	fmt.Print("please choice `start` or `exit` : ")
	var input string
	fmt.Scan(&input)
	if input == "start" {
		fmt.Print("\n\n-_- well down the game is start -_-\n\n")
		for i := 0; i < 3; i++ {
			RockPaperScissors()
		}
		if user_score > computer_score {
			fmt.Println("Hooray, you won this round of the competition (^--^)")
			fmt.Println(user_score, computer_score)
		} else if computer_score > user_score {
			fmt.Println("Damn, you lost this chance, my dear :(")
			fmt.Println(user_score, computer_score)
		} else {
			fmt.Println("Well, you are equal now, hey -ــ-")
			fmt.Println(user_score, computer_score)
		}
	} else if input == "exit" {
		fmt.Println("--- EXIT ---")
	} else {
		fmt.Print("\n--- EXIT ---\n")
		fmt.Print("\nInvalid input Please enter a valid input\n")
	}
}
