package game

import (
	"testing"
)

func TestAddToCup(t *testing.T) {
	cup := Cup{}
	die1, _ := CreateDie(Home, 0)
	die2, _ := CreateDie(Alien, 0)
	die3, _ := CreateDie(Novelty, 0)
	cup.AddDieToCup(die1)
	cup.AddDieToCup(die2)
	cup.AddDieToCup(die3)

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
