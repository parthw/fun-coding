package compute

import (
	"math/rand"
	"time"
)

func getcomputerChoice() string {
	choices := []string{"stone", "paper", "scissors"}
	rand.Seed(time.Now().UnixNano())
	return choices[rand.Intn(3)]
}

// GetWinner function To fetch the winner of the game
func GetWinner(userChoice string) (string, string) {
	computerChoice := getcomputerChoice()
	if userChoice == computerChoice {
		return "Tie 🤦", "Computer choice is same as user choice"
	}
	if userChoice == "stone" && computerChoice == "paper" {
		return "User LOST 🙁", "Computer got paper and WINS"
	} else if userChoice == "stone" && computerChoice == "scissors" {
		return "User WON ✌️", "Computer got scissors and LOST"
	} else if userChoice == "paper" && computerChoice == "stone" {
		return "User WON ✌️", "Computer got stone and LOST"
	} else if userChoice == "paper" && computerChoice == "scissors" {
		return "User LOST 🙁", "Computer got scissors and WINS"
	} else if userChoice == "scissors" && computerChoice == "stone" {
		return "User LOST 🙁", "Computer got stone and WINS"
	} else {
		return "User WON ✌️", "Computer got paper and LOST"
	}

}
