package league

type Player struct {
	Name       string
	Number     uint8
	Position   string
	Height     uint8
	Age        uint8
	Draft      Draft // Round selected
	Experience uint8 // experience in years
}

// Drafted returns true if player has been drafted(round is not 0)
func (p *Player) Drafted() bool {
	if p.Draft.Round == 0 {
		return false
	}

	return true
}

type Draft struct {
	Year         uint16
	Round        uint8
	Team         string
	Overall      uint16
	Pick         uint8
	NumOfPlayers uint8
}
