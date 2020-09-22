package dice

import (
	"math/rand"
	"time"
)

// Roll accepts variable sides and retuns random number
func Roll(sides int) int {
	if sides == 0 {
		sides = 6
	}

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(sides)

	// so zero will not be returned
	return x + 1
}
