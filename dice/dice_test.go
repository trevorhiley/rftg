package dice

import (
	"testing"
)

func TestDiceRoll(t *testing.T) {
	diceRolled := Die{}
	found := false
	sides := 6
	diceRolled.Roll(sides)
	for i := 1; i <= sides; i++ {
		if diceRolled.rolledNumber == i {
			found = true
			break
		}
	}
	if !found {
		t.Error("Dice roll was", diceRolled.rolledNumber)
	}
}
