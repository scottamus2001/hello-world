package main

import (
	"fmt"

	"github.com/scottamus2001/hello-world/game/league"
)

func main() {
	p := league.Player{
		Name:       "Ben Walker",
		Number:     12,
		Position:   "PG",
		Height:     76,
		Age:        21,
		Experience: 0,
	}
	t := league.Team{
		City:         "Rock Hill",
		Mascot:       "Hornets",
		Abbreviation: "RHH",
	}
	t.Draft(1, 10, 0, 2020, 10, &p)
	fmt.Println("player drafted? ", p.Drafted())
	fmt.Printf("Team : %+v\n", t)
}
