package game

import (
	"testing"
)

func TestAddToCup(t *testing.T) {
	cup := Cup{}
	cup.AddDieToCup(CreateDie(Home, 0))
	cup.AddDieToCup(CreateDie(Alien, 0))
	cup.AddDieToCup(CreateDie(Novelty, 0))

	if len(cup.DiceInCup) != 3 {
		t.Error("proper number of dice not added")
	}

	for _, die := range cup.DiceInCup {
		if die.RollResult != 0 {
			t.Error("value set before dice rolled")
		}
	}

	cup.RollCup()

	for _, die := range cup.DiceInCup {
		if die.RollResult == 0 {
			t.Error("dice did not get rolled")
		}
	}
}
