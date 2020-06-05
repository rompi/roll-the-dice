package logic_test

import (
	"fmt"
	"testing"
	"github.com/rompi/roll-the-dice/logic"
)

func TestRollTheDice(t *testing.T) {
	t.Run("run once", func(t *testing.T) {
		diceNumber := logic.RollTheDice()
		minDiceNumber := logic.MinDiceNumber
		maxDiceNumber := logic.MaxDiceNumber
		fmt.Println(diceNumber)
		if diceNumber < minDiceNumber || diceNumber > maxDiceNumber {
			t.Errorf("expected dice number between %v and %v, got %v", minDiceNumber, maxDiceNumber, diceNumber)
		}
	})
}
