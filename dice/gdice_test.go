package dice

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

			if die.name != v {
				t.Error(v, "die not wild")
			}

			if len(die.sides) != 6 {
				t.Error(v, "die sides != 6")
			}

			correctSides := getDieSides()

			die.Roll()
			foundSide := false

			for _, v := range correctSides {
				if die.rollResult == v {
					foundSide = true
				}
			}

			if !foundSide {
				t.Error(v, "dice roll not working, no valid side found")
			}
		}
	}

}
