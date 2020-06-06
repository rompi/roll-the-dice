package logic

import "fmt"

func RegisterPlayer(numberOfPlayer int, players map[int]string)  {
	var playerName string
	for i := 0; i < numberOfPlayer; i++ {
		fmt.Printf("Please input name for player %v: ", i+1)
		_, err := fmt.Scanln(&playerName)
		panicIfError(err)

		players[i] = playerName
	}
}

func panicIfError(err error)  {
	if err != nil {
		panic(err)
	}
}
