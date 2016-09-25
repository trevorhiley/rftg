package game

//Cup has a cup
type Cup struct {
	DiceInCup []GDice
}

//AddDieToCup add dice to cup
func (c *Cup) AddDieToCup(die GDice) {
	c.DiceInCup = append(c.DiceInCup, die)
}

//RollCup roll all the dice in the cup
func (c *Cup) RollCup() {
	for i, die := range c.DiceInCup {
		die.Roll()
		c.DiceInCup[i] = die
	}
}
