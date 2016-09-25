package game

import (
	"testing"
)

func TestAddToCup(t *testing.T) {
	cup := Cup{}
	cup.AddDieToCup(CreateDie(HOME, ""))
	cup.AddDieToCup(CreateDie(ALIEN, ""))
	cup.AddDieToCup(CreateDie(NOVELTY, ""))

	if len(cup.DiceInCup) != 3 {
		t.Error("proper number of dice not added")
	}

	for _, die := range cup.DiceInCup {
		if die.RollResult != "" {
			t.Error("value set before dice rolled")
		}
	}

	cup.RollCup()

	for _, die := range cup.DiceInCup {
		if die.RollResult == "" {
			t.Error("dice did not get rolled")
		}
	}
}
