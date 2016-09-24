package cup

import (
	"testing"

	"github.com/trevorhiley/rftg/dice"
)

func TestAddToCup(t *testing.T) {
	cup := Cup{}
	cup.AddDieToCup(dice.CreateDie(dice.HOME, ""))
	cup.AddDieToCup(dice.CreateDie(dice.ALIEN, ""))
	cup.AddDieToCup(dice.CreateDie(dice.NOVELTY, ""))

	cup.RollCup()

	for _, die := range cup.diceInCup {
		if die.RollResult() == "" {
			t.Error("dice did not get rolled")
		}
	}
}
