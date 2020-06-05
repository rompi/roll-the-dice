package logic

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MaxDiceNumber = 6
	MinDiceNumber = 1
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RollTheDice() int {
	fmt.Println("the dice is rolled...")
	return rollingDice()
}

func rollingDice()  int{
	return rand.Intn(MaxDiceNumber) + 1
}
