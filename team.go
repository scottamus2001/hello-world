package league

type Team struct {
	City         string
	Mascot       string
	Abbreviation string
	Players      []Player
}

func (t *Team) Draft(round, pick, numofplayers uint8, year, overall uint16, p *Player) {
	// todo: draft struct
	d := Draft{
		Year:         year,
		Round:        round,
		Team:         t.Mascot,
		Overall:      overall,
		Pick:         pick,
		NumOfPlayers: numofplayers + 1, //total number of players on team
	}

	// todo: attach draft struct to player
	p.Draft = d

	// todo: add player to team
	t.Players = append(t.Players, *p)
}
