package goforth

import (
	"strings"
)

// Run will take a string of values and operators and evaluate
// each in turn.
// It will then return the final value resulting from your expression.
//
// It will panic if your expression is bad, for example, your expression could
// attempt to evaluate an operation that requires two inputs but if only one is left
// on the stack, it will fall over.
func Run(s string) string {
	stack := newStack()
	ss := strings.Split(s, " ")
	for _, s := range ss {
		if fn, ok := ops[s]; ok {
			fn(stack)
			continue
		}
		stack.push(s)
	}
	return stack.pop()
}
