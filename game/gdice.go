package game

//GDice structure to hold dice defs
type GDice struct {
	Name       string
	Color      string
	Sides      []string
	TradeValue int
	RollResult string
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

//dice Names
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
			Name:       CONSUMPTION,
			Color:      "purple",
			Sides:      []string{EXPLORE, DEVELOP, SHIP, SHIP, SHIP, WILD},
			TradeValue: 0,
		}
	case HOME:
		die = GDice{
			Name:       HOME,
			Color:      "white",
			Sides:      []string{EXPLORE, EXPLORE, DEVELOP, SETTLE, SHIP, PRODUCE},
			TradeValue: 0,
		}
	case MILITARY:
		die = GDice{
			Name:       MILITARY,
			Color:      "red",
			Sides:      []string{EXPLORE, DEVELOP, DEVELOP, SETTLE, SETTLE, WILD},
			TradeValue: 0,
		}
	case NOVELTY:
		die = GDice{
			Name:       NOVELTY,
			Color:      "blue",
			Sides:      []string{EXPLORE, PRODUCE, PRODUCE, SHIP, SHIP, WILD},
			TradeValue: 3,
		}
	case RARE:
		die = GDice{
			Name:       RARE,
			Color:      "brown",
			Sides:      []string{EXPLORE, DEVELOP, DEVELOP, PRODUCE, SHIP, WILD},
			TradeValue: 4,
		}
	case ALIEN:
		die = GDice{
			Name:       ALIEN,
			Color:      "yellow",
			Sides:      []string{DEVELOP, SETTLE, PRODUCE, WILD, WILD, WILD},
			TradeValue: 6,
		}
	case GENES:
		die = GDice{
			Name:       GENES,
			Color:      "green",
			Sides:      []string{EXPLORE, SETTLE, SETTLE, PRODUCE, WILD, WILD},
			TradeValue: 5,
		}
	}
	return die
}

//Roll rolls the gdice
func (g *GDice) Roll() {
	die := Die{}
	die.Roll(6)
	g.RollResult = g.Sides[die.rolledNumber]
}
