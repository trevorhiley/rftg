package dice

//GDice structure to hold dice defs
type GDice struct {
	name       string
	color      string
	sides      []string
	tradeValue int
	rollResult string
}

//constants for dice rolled values
const (
	EXPLORE = "Explore"
	DEVELOP = "Develop"
	SETTLE  = "Settle"
	PRODUCE = "Produce"
	SHIP    = "Ship"
	WILD    = "Wild"
)

//dice names
const (
	CONSUMPTION = "Consumption"
	HOME        = "Home"
	MILITARY    = "Military"
	NOVELTY     = "Novelty"
	RARE        = "Rare Elements"
	ALIEN       = "Alien Technology"
	GENES       = "Genes"
)

//CreateDie creates a purple die
func CreateDie(diceName string, rolledValue string) GDice {
	die := GDice{}
	switch diceName {
	case CONSUMPTION:
		die = GDice{
			name:       CONSUMPTION,
			color:      "purple",
			sides:      []string{EXPLORE, DEVELOP, SHIP, SHIP, SHIP, WILD},
			tradeValue: 0,
		}
	case HOME:
		die = GDice{
			name:       HOME,
			color:      "white",
			sides:      []string{EXPLORE, EXPLORE, DEVELOP, SETTLE, SHIP, PRODUCE},
			tradeValue: 0,
		}
	case MILITARY:
		die = GDice{
			name:       MILITARY,
			color:      "red",
			sides:      []string{EXPLORE, DEVELOP, DEVELOP, SETTLE, SETTLE, WILD},
			tradeValue: 0,
		}
	case NOVELTY:
		die = GDice{
			name:       NOVELTY,
			color:      "blue",
			sides:      []string{EXPLORE, PRODUCE, PRODUCE, SHIP, SHIP, WILD},
			tradeValue: 3,
		}
	case RARE:
		die = GDice{
			name:       RARE,
			color:      "brown",
			sides:      []string{EXPLORE, DEVELOP, DEVELOP, PRODUCE, SHIP, WILD},
			tradeValue: 4,
		}
	case ALIEN:
		die = GDice{
			name:       ALIEN,
			color:      "yellow",
			sides:      []string{DEVELOP, SETTLE, PRODUCE, WILD, WILD, WILD},
			tradeValue: 6,
		}
	case GENES:
		die = GDice{
			name:       GENES,
			color:      "green",
			sides:      []string{EXPLORE, SETTLE, SETTLE, PRODUCE, WILD, WILD},
			tradeValue: 5,
		}
	}
	return die
}

//Roll rolls the gdice
func (g *GDice) Roll() {
	die := Die{}
	die.Roll(6)
	g.rollResult = g.sides[die.rolledNumber]
}

//RollResult getter
func (g *GDice) RollResult() string {
	return g.rollResult
}
