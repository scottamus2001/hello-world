package dice_test

import (
	"testing"

	. "github.com/poy/onpar/expect"
	. "github.com/poy/onpar/matchers"
	"github.com/scottamus2001/hello-world/game/dice"
)

func TestRoll(t *testing.T) {
	tests := []struct {
		description string
		input       int
		max         int
	}{
		{
			description: "0 sides -default to 6 sides",
			input:       0,
			max:         6,
		},
		{
			description: "6 sides",
			input:       6,
			max:         6,
		},
		{
			description: "21 sides",
			input:       21,
			max:         21,
		},
		{
			description: "99 sides",
			input:       99,
			max:         99,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				output := dice.Roll(tc.input)
				Expect(t, output).To(BeAbove(0))
				Expect(t, output).To(Not(BeAbove(float64(tc.max))))
			}
		})
	}
}
