package main

import (
	"fmt"
	"github.com/rompi/roll-the-dice/logic"
)

func main() {
	var numberOfPlayer int
	players := make(map[int]string)
	playerScore := make(map[int]int)
	fmt.Print("Input number of players: ")
	_, err := fmt.Scanln(&numberOfPlayer)
	panicIfError(err)
	logic.RegisterPlayer(numberOfPlayer, players)
	logic.RunTheGame(numberOfPlayer, players, playerScore)
	winner, maxScore := logic.GetTheWinner(numberOfPlayer, players, playerScore)

	fmt.Printf("The winner is %v with the score %v", winner, maxScore)
}

func panicIfError(err error)  {
	if err != nil {
		panic(err)
	}
}
