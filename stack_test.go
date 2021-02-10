package goforth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := newStack()
	stack.push("test")
	assert.Equal(t, []string{"test"}, stack)
	stack.push("test 1")
	stack.push("test 2")
	assert.Equal(t, []string{"test", "test 1", "test 2"}, stack)
	stack.pop()
	stack.pop()
	assert.Equal(t, []string{"test"}, stack)
	stack.pop()
	assert.Equal(t, []string{}, stack)
}
