package cup

import "github.com/trevorhiley/rftg/dice"

//Cup has a cup
type Cup struct {
	diceInCup []dice.GDice
}

//AddDieToCup add dice to cup
func (c *Cup) AddDieToCup(die dice.GDice) {
	c.diceInCup = append(c.diceInCup, die)
}

//RollCup roll all the dice in the cup
func (c *Cup) RollCup() {
	for i, die := range c.diceInCup {
		die.Roll()
		c.diceInCup[i] = die
	}
}
