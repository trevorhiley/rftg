package game

//DiceType for setting dice type constants
type DiceType int

//dice Names
const (
	Consumption DiceType = iota + 1
	Home
	Military
	Novelty
	Rare
	Alien
	Genes
)

//DiceTypeMap return map of dice side to string
func DiceTypeMap() map[DiceType]string {
	return map[DiceType]string{
		Home:        "Home",
		Alien:       "Alien Technology",
		Novelty:     "Novelty Goods",
		Genes:       "Genes",
		Military:    "Military",
		Rare:        "Rare Goods",
		Consumption: "Consumption",
	}
}

func (dt DiceType) String() string {
	return DiceTypeMap()[dt]
}

//StringToDiceType converts string back to const
func StringToDiceType(diceType string) DiceType {
	diceTypeMap := DiceTypeMap()
	for k, v := range diceTypeMap {
		if v == diceType {
			return k
		}
	}
	return 0
}
