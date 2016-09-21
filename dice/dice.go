package dice

import (
  "math/rand"
  "time"
  "fmt"
)

type dice struct {
  rolledNumber int
}

func (d *dice) roll()  int  {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  num := r.Intn(6) + 1
  d.rolledNumber = num
  return d.rolledNumber
}
