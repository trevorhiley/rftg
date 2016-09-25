package game

import "testing"

func getDieSides() []DiceSide {
	return []DiceSide{Explore, Develop, Settle, Produce, Ship, Wild}
}

func getDieTypes() []DiceType {
	return []DiceType{Home, Alien, Rare, Consumption, Military, Genes, Novelty}
}

func TestDiceSidesString(t *testing.T) {
	diceString := Explore.String()
	if diceString != "Explore" {
		t.Error("Dice side stringer not working")
	}
}

func TestStringToDiceSide(t *testing.T) {
	diceSideConst := StringToDiceSide("Explore")
	if diceSideConst != Explore {
		t.Error("Dice string to dice side not working")
	}
}

func TestConsumptionDieCreation(t *testing.T) {
	for _, v := range getDieTypes() {
		for i := 0; i <= 1000; i++ {
			die := CreateDie(v, 0)

			if die.Name != v {
				t.Error(v, "die not wild")
			}

			if len(die.Sides) != 6 {
				t.Error(v, "die sides != 6")
			}

			if die.Color == 0 {
				t.Error("no dice color set")
			}

			if die.TradeValue > 6 {
				t.Error("dice trade value too large")
			}

			correctSides := getDieSides()

			die.Roll()
			foundSide := false

			for _, v := range correctSides {
				if die.RollResult == v {
					foundSide = true
				}
			}

			if !foundSide {
				t.Error(v, "dice roll not working, no valid side found")
			}
		}
	}

}
