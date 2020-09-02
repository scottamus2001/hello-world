package locale_test

import (
	"testing"

	. "github.com/poy/onpar/expect"
	. "github.com/poy/onpar/matchers"
	"github.com/scottamus2001/hello-world/locale"
)

func TestHello(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    string
	}{
		{
			description: "hello in english",
			input:       "",
			expected:    "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			output := locale.Hello(tc.input)
			Expect(t, output).To(Equal(tc.expected))
		})
	}
}
