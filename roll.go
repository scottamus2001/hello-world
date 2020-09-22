package main

import (
	"fmt"

	"github.com/scottamus2001/hello-world/game/dice"
)

func main() {
	fmt.Println("first dice roll is", dice.Roll(6))
	fmt.Println("second dice roll is", dice.Roll(6))
}
