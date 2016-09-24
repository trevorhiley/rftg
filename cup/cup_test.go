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

	if len(cup.diceInCup) != 3 {
		t.Error("proper number of dice not added")
	}

	for _, die := range cup.diceInCup {
		if die.RollResult() != "" {
			t.Error("value set before dice rolled")
		}
	}

	cup.RollCup()

	for _, die := range cup.diceInCup {
		if die.RollResult() == "" {
			t.Error("dice did not get rolled")
		}
	}
}
