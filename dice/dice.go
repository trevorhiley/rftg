package dice

import (
	"math/rand"
	"time"
)

//Die Struct for rolling a die
type Die struct {
	rolledNumber int
}

//Roll the die
func (d *Die) Roll(sides int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	d.rolledNumber = r.Intn(sides)
	return d.rolledNumber
}
