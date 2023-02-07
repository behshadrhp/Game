package main

import "fmt"

func main() {
	fmt.Print("Hello Welcome to Rock-Paper-Scissors Game :)\n\n")
	fmt.Print("If you want to start the game, write the word `start` : ")
	var started string
	fmt.Scan(&started)
	if started == "start" {
		fmt.Print("\n\n-_- well down the game is start -_-\n\n")
	} else {
		fmt.Print("\n\n\n---_--- End Game ---_---\n\n")
		fmt.Print("\nInvalid input Please enter a valid input\n")
	}
}
