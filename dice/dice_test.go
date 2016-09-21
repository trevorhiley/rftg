package dice

import (
  "testing"
)

func TestDiceRoll(t *testing.T)  {
  dice := dice{}
  found:= false
  start := 1
  dice.roll()
  for start <= 6 {
    if dice.rolledNumber == start {
      found = true
      break
    }
    start++
  }
  if !found {
    t.Fail()
  }
  return
}
