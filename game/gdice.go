package game

//GDice structure to hold dice defs
type GDice struct {
	Name       DiceType
	Color      DiceColor
	Sides      []DiceSide
	TradeValue int
	RollResult DiceSide
}

//DiceColor for setting dice color constants
type DiceColor int

//DiceColor constants for dice colors
const (
	Purple DiceColor = iota + 1
	Yellow
	Green
	Red
	White
	Blue
	Brown
)

//CreateDie creates a purple die
func CreateDie(diceName DiceType, rolledValue DiceSide) GDice {
	die := GDice{}
	switch diceName {
	case Consumption:
		die = GDice{
			Name:       diceName,
			Color:      Purple,
			Sides:      []DiceSide{Explore, Develop, Ship, Ship, Ship, Wild},
			TradeValue: 0,
		}
	case Home:
		die = GDice{
			Name:       diceName,
			Color:      White,
			Sides:      []DiceSide{Explore, Explore, Develop, Settle, Ship, Produce},
			TradeValue: 0,
		}
	case Military:
		die = GDice{
			Name:       diceName,
			Color:      Red,
			Sides:      []DiceSide{Explore, Develop, Develop, Settle, Settle, Wild},
			TradeValue: 0,
		}
	case Novelty:
		die = GDice{
			Name:       diceName,
			Color:      Blue,
			Sides:      []DiceSide{Explore, Produce, Produce, Ship, Ship, Wild},
			TradeValue: 3,
		}
	case Rare:
		die = GDice{
			Name:       diceName,
			Color:      Brown,
			Sides:      []DiceSide{Explore, Develop, Develop, Produce, Ship, Wild},
			TradeValue: 4,
		}
	case Alien:
		die = GDice{
			Name:       diceName,
			Color:      Yellow,
			Sides:      []DiceSide{Develop, Settle, Produce, Wild, Wild, Wild},
			TradeValue: 6,
		}
	case Genes:
		die = GDice{
			Name:       diceName,
			Color:      Green,
			Sides:      []DiceSide{Explore, Settle, Settle, Produce, Wild, Wild},
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
