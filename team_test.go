package league_test

import (
	"testing"

	. "github.com/poy/onpar/expect"
	. "github.com/poy/onpar/matchers"
	"github.com/scottamus2001/hello-world/game/league"
)

func TestDraft(t *testing.T) {
	team := league.Team{}

	tests := []struct {
		description string
		draft       league.Draft
		player      league.Player
	}{
		{
			description: "draft 1",
			draft: league.Draft{
				Year:    2020,
				Round:   2,
				Pick:    20,
				Overall: 50,
			},
			player: league.Player{
				Name:     "Patrick ewing",
				Position: "center",
				Number:   24,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			player_cnt := len(team.Players)
			team.Draft(tc.draft.Round, tc.draft.Pick, tc.draft.Year, tc.draft.Overall, &tc.player)
			Expect(t, len(team.Players)).To(Equal(player_cnt + 1))

			lastPlayer := team.Players[len(team.Players)-1]
			Expect(t, lastPlayer.Name).To(Equal(tc.player.Name))
			Expect(t, tc.player.Draft.Year).To(Equal(tc.draft.Year))
			Expect(t, lastPlayer.Draft.Overall).To(Equal(tc.draft.Overall))

		})
	}
}
