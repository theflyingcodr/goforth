package goforth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		input string
		exp   string
	}{
		"adding two numbers should work": {
			input: "10 43 ADD",
			exp:   "53",
		}, "adding two numbers invalid case should not exec operator": {
			input: "10 43 Add",
			exp:   "Add",
		}, "adding two numbers then adding result should work": {
			input: "10 43 ADD 10 ADD",
			exp:   "63",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.exp, Run(test.input))
		})
	}
}
