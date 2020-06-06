package logic

import "fmt"

func GetTheWinner(numberOfPlayers int, players map[int]string, playerScore map[int]int) (string, int)  {
	maxScore := -15
	var winner string
	for i := 0; i < numberOfPlayers; i++ {
		if playerScore[i] > maxScore {
			maxScore = playerScore[i]
			winner = players[i]
		} else if playerScore[i] == maxScore {
			winner = winner + ", " + players[i]
		}
	}
	return winner, maxScore
}

func RunTheGame(numberOfPlayer int, players map[int]string, playerScore map[int]int)  {
	const numberOfRound = 5
	for i := 0; i < numberOfRound; i++ {
		for j := 0; j < numberOfPlayer; j++ {
			fmt.Printf("%v rolls the dice... (press enter to roll)", players[j])
			fmt.Scanln()
			diceNumber := RollTheDice()
			fmt.Printf("the dice shows %v!\n", diceNumber)
			dicePoint := CalculateTheDiceNumber(diceNumber)
			playerScore[j] += dicePoint
			fmt.Printf("additional %v for %v, so the score is now %v\n", dicePoint, players[j], playerScore[j])
		}
	}
}