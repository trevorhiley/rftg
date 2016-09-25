package game

//DiceSide constants for dice Sides
type DiceSide int

//constants for dice rolled values
const (
	Explore DiceSide = iota + 1
	Develop
	Settle
	Produce
	Ship
	Wild
)

//DiceSideMap return map of dice side to string
func DiceSideMap() map[DiceSide]string {
	return map[DiceSide]string{
		Explore: "Explore",
		Develop: "Develop",
		Settle:  "Settle",
		Produce: "Produce",
		Ship:    "Ship",
		Wild:    "Wild",
	}
}

func (ds DiceSide) String() string {
	return DiceSideMap()[ds]
}

//StringToDiceSide converts string back to const
func StringToDiceSide(diceSide string) DiceSide {
	diceSideMap := DiceSideMap()
	for k, v := range diceSideMap {
		if v == diceSide {
			return k
		}
	}
	return 0
}
