package game

import (
	"testing"
)

func TestDiceRoll(t *testing.T) {
	diceRolled := Die{}
	found := false
	sides := 6
	diceRolled.Roll(sides)
	for i := 0; i <= sides-1; i++ {
		if diceRolled.rolledNumber == i {
			found = true
			break
		}
	}
	if !found {
		t.Error("Dice roll was", diceRolled.rolledNumber)
	}
}
