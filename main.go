package main

import (
	"fmt"
	"github.com/rompi/roll-the-dice/logic"
)

func main() {
	const numberOfRound = 5
	var numberOfPlayer int
	players := make(map[int]string)
	playerScore := make(map[int]int)
	var playerName string
	fmt.Print("Input number of players: ")
	_, err := fmt.Scanln(&numberOfPlayer)
	if err != nil {
		panic(err)
	}
	for i := 0; i < numberOfPlayer; i++ {
		fmt.Printf("Please input name for player %v: ", i+1)
		_, err = fmt.Scanln(&playerName)
		if err != nil {
			panic(err)
		}

		players[i] = playerName
	}

	for i := 0; i < numberOfRound; i++ {
		for j := 0; j < numberOfPlayer; j++ {
			fmt.Printf("%v rolls the dice... (press enter to roll)", players[j])
			fmt.Scanln()
			diceNumber := logic.RollTheDice()
			fmt.Printf("the dice shows %v!\n", diceNumber)
			if diceNumber % 2 == 0 {
				playerScore[j] -= 3
				fmt.Printf("minus 3 for %v, so the score is now %v\n", players[j], playerScore[j])
			} else {
				playerScore[j] += 5
				fmt.Printf("plus 5 for %v, so the score is now %v\n", players[j], playerScore[j])
			}
		}
	}

	fmt.Println(playerScore)
	var maxScore int
	var winner string
	for i := 0; i < numberOfPlayer; i++ {
		if playerScore[i] > maxScore {
			maxScore = playerScore[i]
			winner = players[i]
		} else if playerScore[i] == maxScore{
			winner = winner + ", " + players[i]
		}
	}

	fmt.Printf("The winner is %v with the score %v", winner, maxScore)
}
