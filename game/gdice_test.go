package game

import "testing"

func getDieSides() []string {
	return []string{EXPLORE, DEVELOP, SETTLE, PRODUCE, SHIP, WILD}
}

func getDieTypes() []string {
	return []string{HOME, ALIEN, RARE, CONSUMPTION, MILITARY, GENES, NOVELTY}
}

func TestConsumptionDieCreation(t *testing.T) {
	for _, v := range getDieTypes() {
		for i := 0; i <= 1000; i++ {
			die := CreateDie(v, "")

			if die.Name != v {
				t.Error(v, "die not wild")
			}

			if len(die.Sides) != 6 {
				t.Error(v, "die sides != 6")
			}

			if die.Color == "" {
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
